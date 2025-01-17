package migration

import (
	"github.com/bigxxby/digital-travel-test/internal/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		return err
	}

	return nil
}
