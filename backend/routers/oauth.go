package routers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"

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
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid oauth state"})
		return
	}

	code := c.Query("code")
	token, err := oauthConf.Exchange(context.Background(), code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to exchange token"})
		return
	}

	client := oauthConf.Client(context.Background(), token)
	resp, err := client.Get("https://api.github.com/user")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get user info"})
		return
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)
	var user map[string]interface{}
	if err := json.Unmarshal(data, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse user info"})
		return
	}

	// 获取用户邮箱信息
	emailResp, err := client.Get("https://api.github.com/user/emails")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get user emails"})
		return
	}
	defer emailResp.Body.Close()

	emailData, _ := io.ReadAll(emailResp.Body)
	var emails []map[string]interface{}
	if err := json.Unmarshal(emailData, &emails); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to parse user emails"})
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

	// 返回用户信息和主邮箱
	c.JSON(http.StatusOK, gin.H{
		"name":  user["name"],
		"email": primaryEmail,
	})
}

func RegisterOAuthRoutes(r *gin.RouterGroup) {
	r.GET("/login", OAuthLogin)
	r.GET("/callback", OAuthCallback)
}
