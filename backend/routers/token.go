package routers

import (
	"crypto/ed25519"
	"encoding/base64"
	"nep-keychain-backend/config"
	"nep-keychain-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GenerateKeyPair 生成ED25519密钥对
func GenerateKeyPair() (publicKey, privateKey string) {
	pub, priv, _ := ed25519.GenerateKey(nil)
	return base64.StdEncoding.EncodeToString(pub), base64.StdEncoding.EncodeToString(priv)
}

// NewTokenHandler 生成新的密钥对
func NewTokenHandler(c *gin.Context) {
	var user models.User

	// 查找当前用户
	result := config.DB.First(&user, c.GetUint("userID"))
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户未找到"})
		return
	}

	// 检查用户是否已经有密钥对
	if user.PublicToken != "" {
		c.JSON(http.StatusConflict, gin.H{"error": "用户已经有密钥对"})
		return
	}

	// 生成密钥对
	publicKey, privateKey := GenerateKeyPair()

	// 更新用户信息
	user.PublicToken = publicKey
	user.PrivateToken = privateKey
	config.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{
		"message": "密钥对生成成功",
	})
}

// GetPublicKeyHandler 获取用户公钥
func GetPublicKeyHandler(c *gin.Context) {
	var user models.User
	name := c.Query("name")

	// 查找用户
	result := config.DB.Where("name = ? OR email = ?", name, name).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户未找到"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"public_key": user.PublicToken,
	})
}

// GetKeyPairHandler 获取当前用户的密钥对
func GetKeyPairHandler(c *gin.Context) {
	var user models.User

	// 查找当前用户
	result := config.DB.First(&user, c.GetUint("userID"))
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "用户未找到"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"public_key":  user.PublicToken,
		"private_key": user.PrivateToken,
	})
}

func RegisterTokenRoutes(r *gin.RouterGroup) {
	r.POST("/new", config.JWTMiddleware(), NewTokenHandler)
	r.GET("/public", GetPublicKeyHandler)
	r.GET("/keypair", config.JWTMiddleware(), GetKeyPairHandler)
}
