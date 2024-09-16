package models

import (
	"github.com/Asonance11/Application-tracker/internal/database"
	"github.com/Asonance11/Application-tracker/internal/types"
)

func CreateUser(user *types.User) error {
	if err := database.GetDB().Create(user).Error; err != nil {
		return err
	}

	return nil
}
