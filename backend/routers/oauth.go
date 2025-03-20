package routers

import (
	"context"
	"encoding/json"
	"io"
	"nep-keychain-backend/config"
	"nep-keychain-backend/models"
	"nep-keychain-backend/utils"
	"net/http"
	"os"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"

	"nep-keychain-backend/vo"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	oauthConf = &oauth2.Config{
		ClientID:     "",
		ClientSecret: "",
		RedirectURL:  "",
		Scopes:       []string{"user:email"},
		Endpoint:     github.Endpoint,
	}
	oauthStateString = "github-login"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}
	oauthConf.ClientID = os.Getenv("GITHUB_CLIENT_ID")
	oauthConf.ClientSecret = os.Getenv("GITHUB_CLIENT_SECRET")
	oauthConf.RedirectURL = os.Getenv("GITHUB_REDIRECT_URL")
}

func OAuthLogin(c *gin.Context) {
	url := oauthConf.AuthCodeURL(oauthStateString, oauth2.AccessTypeOffline)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func OAuthCallback(c *gin.Context) {
	state := c.Query("state")
	if state != oauthStateString {
		c.JSON(http.StatusBadRequest, vo.BaseResponse[any]{Code: http.StatusBadRequest, Msg: config.Translate("invalid_oauth_state", c), Data: nil})
		return
	}

	code := c.Query("code")
	token, err := oauthConf.Exchange(context.Background(), code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, vo.BaseResponse[any]{Code: http.StatusInternalServerError, Msg: config.Translate("token_exchange_error", c), Data: nil})
		return
	}

	client := oauthConf.Client(context.Background(), token)
	resp, err := client.Get("https://api.github.com/user")
	if err != nil {
		c.JSON(http.StatusInternalServerError, vo.BaseResponse[any]{Code: http.StatusInternalServerError, Msg: config.Translate("get_user_info_error", c), Data: nil})
		return
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)
	var gh_user map[string]interface{}
	if err := json.Unmarshal(data, &gh_user); err != nil {
		c.JSON(http.StatusInternalServerError, vo.BaseResponse[any]{Code: http.StatusInternalServerError, Msg: config.Translate("parse_user_info_error", c), Data: nil})
		return
	}

	// 获取用户邮箱信息
	emailResp, err := client.Get("https://api.github.com/user/emails")
	if err != nil {
		c.JSON(http.StatusInternalServerError, vo.BaseResponse[any]{Code: http.StatusInternalServerError, Msg: config.Translate("get_user_email_error", c), Data: nil})
		return
	}
	defer emailResp.Body.Close()

	emailData, _ := io.ReadAll(emailResp.Body)
	var emails []map[string]interface{}
	if err := json.Unmarshal(emailData, &emails); err != nil {
		c.JSON(http.StatusInternalServerError, vo.BaseResponse[any]{Code: http.StatusInternalServerError, Msg: config.Translate("parse_user_email_error", c), Data: nil})
		return
	}

	// 提取主邮箱
	var primaryEmail string
	for _, email := range emails {
		if email["primary"].(bool) {
			primaryEmail = email["email"].(string)
			break
		}
	}

	// 如果用户邮箱不存在则在数据库中创建用户
	var user models.User
	var tmpOpt string
	config.DB.Where("email = ?", primaryEmail).First(&user)
	if user.Email == "" {
		tmpPwd := utils.RandomString(16)
		// 生成密钥对
		publicKey, privateKey := GenerateKeyPair()
		user = models.User{
			Name:         gh_user["login"].(string),
			Email:        primaryEmail,
			PwdHash:      utils.HashStringToHexBlake3(tmpPwd),
			PublicToken:  publicKey,
			PrivateToken: privateKey,
		}
		config.DB.Create(&user)
		// 生成临时 opt
		tmpOpt = utils.GenerateRandomCode()
		// 写到数据库中
		expireAt := time.Now().Add(time.Minute * 10)
		allowResend := time.Now().Add(time.Second * 60)
		config.DB.Create(&models.Verify{Email: primaryEmail, VerifyCode: tmpOpt, ExpireAt: expireAt, AllowResend: allowResend})
	}

	// 生成token
	authToken, _ := config.GenerateToken(user.ID)

	// 返回用户信息和主邮箱
	c.JSON(http.StatusOK, vo.BaseResponse[map[string]string]{Code: http.StatusOK, Msg: config.Translate("success", c), Data: map[string]string{"name": user.Name, "email": primaryEmail, "token": authToken, "tmpOpt": tmpOpt}})
}

func RegisterOAuthRoutes(r *gin.RouterGroup) {
	r.GET("/login", OAuthLogin)
	r.GET("/callback", OAuthCallback)
}
