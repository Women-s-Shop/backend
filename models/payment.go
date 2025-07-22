package models

import "time"

type Payment struct {
	ID            uint `json:"id"`
	OrderID       uint
	PaymentMethod string `json:"payment_method"`
	PaymentStatus string `json:"payment_status"`
	TransactionID string
	PaidAt        time.Time
}
