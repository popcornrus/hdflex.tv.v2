package services

import (
	"go-rust-drop/internal/api/repositories"
	"gorm.io/gorm"
)

type UserService struct {
	userRepo repositories.UserRepository
	db       *gorm.DB
}
