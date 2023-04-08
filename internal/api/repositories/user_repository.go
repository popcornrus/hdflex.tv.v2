package repositories

import (
	"go-rust-drop/internal/api/models"
	"gorm.io/gorm"
)

type UserRepository struct {
}

func (ur UserRepository) FindUserByID(userID uint64) (models.User, error) {
	var err error
	var user models.User

	if err = MysqlDB.Where("id = ?", userID).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return models.User{}, err
		}
		return models.User{}, err
	}

	return user, nil
}
