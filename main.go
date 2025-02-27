package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"nep-keychain-backend/config"
	"nep-keychain-backend/models"
	"nep-keychain-backend/routers"
	"net/http"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// User
	routers.RegisterUserRoutes(r)

	return r
}

func main() {
	config.InitDB()
	config.DB.AutoMigrate(&models.User{})
	log.Println("Database initialized and tables migrated!")

	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
