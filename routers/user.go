package routers

import (
	"github.com/gin-gonic/gin"
	"nep-keychain-backend/config"
	"nep-keychain-backend/models"
	"net/http"
)

func RegisterUserRoutes(r *gin.Engine) {
	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		var userModel models.User
		config.DB.Where("name = ?", user).First(&userModel)

		c.JSON(http.StatusOK, gin.H{"user": user, "value": userModel.Name})
	})
	// Create user
	r.POST("/user", func(c *gin.Context) {
		var user models.User
		c.BindJSON(&user)
		config.DB.Create(&user)
		c.JSON(http.StatusOK, gin.H{"user": user})
	})
	// Update user
	r.PUT("/user", func(c *gin.Context) {

		var nextUser models.User
		c.BindJSON(&nextUser)

		var curUser models.User
		config.DB.Where("name =?", nextUser.Name).First(&curUser)

		config.DB.Model(&curUser).Updates(nextUser)

		c.JSON(http.StatusOK, gin.H{"user": nextUser})
	})
	// Delete user
	r.DELETE("/user/:name", func(c *gin.Context) {
		var user models.User
		config.DB.Where("name =?", c.Params.ByName("name")).Delete(&user)
		c.JSON(http.StatusOK, gin.H{"user": user})
	})
}
