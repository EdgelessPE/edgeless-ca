package models

import (
	"time"

	"gorm.io/gorm"
)

type Limit struct {
	gorm.Model
	Ip          string `gorm:"unique;index;not null;type:inet"`
	ActionEmail string `gorm:"not null"`
	ExpireAt    time.Time
}
