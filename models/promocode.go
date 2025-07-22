package models

import "time"

type Promocode struct {
	ID              uint   `json:"id"`
	Code            string `json:"code"`
	DiscountPercent *int
	DiscountAmount  *float64
	ValidFrom       time.Time
	ValidUntil      time.Time
	UsageLimit      int
	TimesUsed       int
}
