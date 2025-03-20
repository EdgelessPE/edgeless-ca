package routers

import (
	"crypto/ed25519"
	"encoding/base64"
	"nep-keychain-backend/config"
	"nep-keychain-backend/models"
	"nep-keychain-backend/vo"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GenerateKeyPair 生成ED25519密钥对
func GenerateKeyPair() (publicKey, privateKey string) {
	pub, priv, _ := ed25519.GenerateKey(nil)
	return base64.StdEncoding.EncodeToString(pub), base64.StdEncoding.EncodeToString(priv)
}

// GetPublicKeyHandler 获取用户公钥
func GetPublicKeyHandler(c *gin.Context) {
	var user models.User
	name := c.Query("name")

	// 查找用户
	result := config.DB.Where("name = ? OR email = ?", name, name).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, vo.BaseResponse[any]{Code: http.StatusNotFound, Msg: config.Translate("user_not_found", c), Data: nil})
		return
	}

	c.JSON(http.StatusOK, vo.BaseResponse[string]{Code: http.StatusOK, Msg: config.Translate("success", c), Data: user.PublicToken})
}

// GetKeyPairHandler 获取当前用户的密钥对
func GetKeyPairHandler(c *gin.Context) {
	var user models.User

	// 查找当前用户
	result := config.DB.First(&user, c.GetUint("userID"))
	if result.Error != nil {
		c.JSON(http.StatusNotFound, vo.BaseResponse[any]{Code: http.StatusNotFound, Msg: config.Translate("user_not_found", c), Data: nil})
		return
	}

	c.JSON(http.StatusOK, vo.BaseResponse[map[string]string]{Code: http.StatusOK, Msg: config.Translate("success", c), Data: map[string]string{"public_key": user.PublicToken, "private_key": user.PrivateToken}})
}

func RegisterTokenRoutes(r *gin.RouterGroup) {
	r.GET("/public", GetPublicKeyHandler)
	r.GET("/keypair", config.JWTMiddleware(), GetKeyPairHandler)
}
