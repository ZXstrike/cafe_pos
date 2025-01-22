package models

import (
	"time"

	"gorm.io/gorm"
)

type MenuHistory struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	MenuID    uint           `json:"menu_id"`
	Quantity  int            `json:"quantity"`
	Total     int            `json:"total"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
