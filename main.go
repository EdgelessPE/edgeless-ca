package main

import (
	"log"
	"nep-keychain-backend/config"
	"nep-keychain-backend/models"
	"nep-keychain-backend/routers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Routers
	routers.RegisterUserRoutes(r)
	routers.RegisterAuthRoutes(r.Group("/auth"))

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
