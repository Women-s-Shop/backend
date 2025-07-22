package models

import "time"

type Order struct {
	ID     uint `json:"id"`
	UserID uint

	OrderStatus     string `json:"order_status"`
	TotalAmount     float64
	ShippingAddress string `json:"shipping_address"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
