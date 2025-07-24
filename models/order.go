package models

import "time"

type Order struct {
	ID              uint    `gorm:"primaryKey" json:"id"`
	UserID          uint    `gorm:"not null"`
	OrderStatus     string  `json:"order_status"`
	TotalAmount     float64 `json:"total_amount"`
	ShippingAddress string  `json:"shipping_address"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
