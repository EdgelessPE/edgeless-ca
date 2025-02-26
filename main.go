package main

import (
	"log"
	"nep-keychain-backend/config"
	"nep-keychain-backend/models"
)

func main() {
	config.InitDB()
	config.DB.AutoMigrate(&models.User{})

	log.Println("Database initialized and tables migrated!")
}
