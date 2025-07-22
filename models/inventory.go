package models

import "time"

type Inventory struct {
	Id        int `json:"id"`
	ProductId uint
	Quantity  int    `json:"quantity"`
	Size      string `json:"size"`
	Color     string `json:"color"`
	UpdateAt  time.Time
}
