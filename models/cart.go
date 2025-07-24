package models

import "time"

type Cart struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	UserID    uint   `gorm:"not null"`
	ProductID uint   `gorm:"not null"`
	Quantity  string `json:"quantity"`
	Size      string `json:"size"`
	Color     string `json:"color"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
