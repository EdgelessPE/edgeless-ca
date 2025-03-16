package models

import (
	"time"

	"gorm.io/gorm"
)

type Verify struct {
	gorm.Model
	Email       string `gorm:"unique;index;not null"`
	VerifyCode  string
	ExpireAt    time.Time
	AllowResend time.Time
}
