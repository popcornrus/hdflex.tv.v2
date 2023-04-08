package services

import (
	"go-boilerplate/internal/api/repositories"
	"gorm.io/gorm"
)

type UserService struct {
	userRepo repositories.UserRepository
	db       *gorm.DB
}
