package routers

import (
	"nep-keychain-backend/config"
	"nep-keychain-backend/models"
	"nep-keychain-backend/vo"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.RouterGroup) {
	r.POST("/register", func(c *gin.Context) {
		var payload vo.RegisterPayload
		if err := c.BindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 用户不能为空且长度不少于 7 个字符
		if payload.Name == "" || len(payload.Name) < 7 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "user name is required and must be at least 7 characters long"})
			return
		}

		// 密码不能为空
		if payload.PwdHash == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "password is required"})
			return
		}

		// 检查用户名和邮箱重复
		var user models.User
		config.DB.Where("name = ?", payload.Name).First(&user)
		if user.Name != "" {
			c.JSON(http.StatusConflict, gin.H{"error": "user name already exists"})
			return
		}
		config.DB.Where("email = ?", payload.Email).First(&user)
		if user.Name != "" {
			c.JSON(http.StatusConflict, gin.H{"error": "user email already exists"})
			return
		}

		user = models.User{
			Name:    payload.Name,
			Email:   payload.Email,
			PwdHash: payload.PwdHash,
		}
		config.DB.Create(&user)
		c.JSON(http.StatusOK, gin.H{"user": user})
	})
	r.POST("/login", config.JWTMiddleware(), func(c *gin.Context) {
		var payload vo.LoginPayload
		if err := c.ShouldBind(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		var user models.User
		// 查找用户，名称或邮箱匹配
		config.DB.Where("name =? OR email =?", payload.Name, payload.Name).First(&user)
		if user.Name == "" {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		}
		if payload.PwdHash != user.PwdHash {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "wrong password"})
		}
		token, _ := config.GenerateToken(user.ID)
		c.JSON(http.StatusOK, gin.H{"name": user.Name, "email": user.Email, "token": token})
	})
}
