package models

import "time"

type Inventory struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	ProductID uint   `gorm:"not null"`
	Quantity  int    `json:"quantity"`
	Size      string `json:"size"`
	Color     string `json:"color"`
	UpdatedAt time.Time
}
