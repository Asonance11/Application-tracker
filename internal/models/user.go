package models

import (
	"errors"

	"github.com/Asonance11/Application-tracker/internal/database"
	"github.com/Asonance11/Application-tracker/internal/types"
	"gorm.io/gorm"
)

func CreateUser(user *types.User) error {
	if err := database.GetDB().Create(user).Error; err != nil {
		return err
	}

	return nil
}

func GetUserByUsername(username string) (*types.User, error) {
	var user types.User

	result := database.GetDB().Where("username = ?", username).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, result.Error
	}

	return &user, nil
}

func GetUserByID(id uint) (*types.User, error) {
	var user types.User

	err := database.GetDB().First(&user, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}

		return nil, err
	}

	return &user, nil
}
