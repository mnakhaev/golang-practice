package models

import "time"

type Product struct {
	ID            string     `json:"id,omitempty"`
	Name          string     `json:"name,omitempty"`
	Description   string     `json:"description,omitempty"`
	ImageID       *string    `json:"image_id,omitempty"`
	Price         int        `json:"price,omitempty"`
	CurrencyID    int        `json:"currency_id,omitempty"`
	Rating        int        `json:"rating,omitempty"`
	CategoryID    int        `json:"category_id,omitempty"`
	Specification *string    `json:"specification,omitempty"`
	CreatedAt     time.Time  `json:"created_at,omitempty"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty"`
}
