package models

type OrderItem struct {
	ID           uint `json:"id"`
	OrderID      uint
	ProductID    uint
	Quantity     int
	Size         string `json:"size"`
	Color        string `json:"color"`
	PriceAtOrder float64
}
