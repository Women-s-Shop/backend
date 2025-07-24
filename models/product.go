package models

import "time"

type Product struct {
	ID          uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	ImageUrl    string `json:"image_url"`
	CategoryID  uint   `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
