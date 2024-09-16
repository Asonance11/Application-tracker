package models

import (
	"github.com/Asonance11/Application-tracker/internal/database"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
}

func CreateUser(user *User) error {
	if err := database.GetDB().Create(user).Error; err != nil {
		return err
	}

	return nil
}
