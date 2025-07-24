package models

import "time"

type Promocode struct {
	ID              uint      `gorm:"primaryKey" json:"id"`
	Code            string    `gorm:"unique;not null" json:"code"`
	DiscountPercent *int      `json:"discount_percent"`
	DiscountAmount  *float64  `json:"discount_amount"`
	ValidFrom       time.Time `json:"valid_from"`
	ValidUntil      time.Time `json:"valid_until"`
	UsageLimit      int       `json:"usage_limit"`
	TimesUsed       int       `json:"times_used"`
}
