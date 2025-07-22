package models

import "time"

type Product struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	ImageUrl    string
	CategoryId  uint
	CreatedAt   time.Time
	UpdateAt    time.Time
}
