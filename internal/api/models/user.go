package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID              *uint64 `gorm:"primaryKey"`
	UUID            string  `gorm:"unique"`
	Name            *string
	AvatarURL       *string
	Email           *string
	EmailVerifiedAt *time.Time
	Password        *string
	Active          bool
	RememberToken   *string
	gorm.Model
}
