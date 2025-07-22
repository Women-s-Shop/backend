package models

import "time"

type Cart struct {
	Id        uint `json:"id"`
	UserId    uint
	ProductId uint
	Quantity  string `json:"quantity"`
	Size      string `json:"size"`
	Color     string `json:"color"`
	CreatedAt time.Time
	UpdateAt  time.Time
}
