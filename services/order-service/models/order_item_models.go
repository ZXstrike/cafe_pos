package models

import (
	"time"

	"gorm.io/gorm"
)

type OrderItem struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	OrderID   uint           `json:"order_id"`
	ProductID uint           `json:"product_id"`
	Quantity  uint           `json:"quantity"`
	Price     float64        `json:"price"`
	Total     float64        `json:"total"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
