package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID        *uuid.UUID     `json:"id" gorm:"type:uuid;primaryKey"`
	Name      string         `json:"name" gorm:"not null"`
	Price     float64        `json:"price" gorm:"not null"`
	Quantity  int            `json:"quantity" gorm:"not null"`
	CreatedAt *time.Time     `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt *time.Time     `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	new := uuid.New()
	p.ID = &new
	return
}
