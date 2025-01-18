package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	ID        *uuid.UUID     `json:"id" gorm:"type:uuid;primaryKey"`
	Name      string         `json:"name" gorm:"not null"`
	CreatedAt *time.Time     `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt *time.Time     `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (r *Role) BeforeCreate(tx *gorm.DB) (err error) {
	new := uuid.New()
	r.ID = &new
	return
}
