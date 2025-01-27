package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	TableID   uint           `json:"table_id"`
	Name      string         `json:"name"`
	Note      string         `json:"note"`
	Total     float64        `json:"total"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
