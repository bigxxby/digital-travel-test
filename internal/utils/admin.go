package utils

import (
	userRepo "github.com/bigxxby/digital-travel-test/internal/api/repo/user"
	"github.com/bigxxby/digital-travel-test/internal/models"
	"gorm.io/gorm"
)

func CreateAdmin(db *gorm.DB) error {
	admin := models.User{
		Username: "admin",
		Password: "Admin123!",
		Role:     "admin",
	}
	err := admin.ValidatePassword()
	if err != nil {
		return err
	}
	admin.HashPassword()
	_, err = userRepo.NewUserRepo(db).CreateUser(admin)
	if err != nil {
		return err
	}
	return nil
}
