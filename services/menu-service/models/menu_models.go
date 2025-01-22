package models

import (
	"time"

	"gorm.io/gorm"
)

type Menu struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Price       int            `json:"price"`
	Quantity    int            `json:"quantity"`
	CategoryID  uint           `json:"category_id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}
