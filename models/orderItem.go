package models

type OrderItem struct {
	ID           uint    `gorm:"primaryKey" json:"id"`
	OrderID      uint    `gorm:"not null"`
	ProductID    uint    `gorm:"not null"`
	Quantity     int     `json:"quantity"`
	Size         string  `json:"size"`
	Color        string  `json:"color"`
	PriceAtOrder float64 `json:"price_at_order"`
}
