package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string `gorm:"unique;index;not null"`
	Email        string `gorm:"unique;index;not null"`
	PwdHash      string `gorm:"not null"`
	PublicToken  string
	PrivateToken string
}
