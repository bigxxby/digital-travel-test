package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	ID         *uuid.UUID     `json:"id" gorm:"type:uuid;primaryKey"`
	CustomerID *uuid.UUID     `json:"customer_id" gorm:"type:uuid;not null"` // Foreign key to User
	Customer   User           `json:"customer" gorm:"foreignKey:CustomerID"` // Relation to User
	Status     string         `json:"status" gorm:"not null"`
	TotalPrice float64        `json:"total_price" gorm:"not null"`
	Products   []Product      `json:"products" gorm:"many2many:order_products;"`
	CreatedAt  *time.Time     `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt  *time.Time     `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

func (o *Order) BeforeCreate(tx *gorm.DB) (err error) {
	new := uuid.New()
	o.ID = &new
	return
}
