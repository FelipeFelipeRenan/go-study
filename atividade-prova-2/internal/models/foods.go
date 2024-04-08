package models

import (
	"time"

	"gorm.io/gorm"
)

type Food struct {
	gorm.Model
	Name         string    `json:"name"`
	Category     string    `json:"category"`
	Quantity     int64     `json:"quantity"`
	Price        float64   `json:"price"`
	ExpirationAt time.Time `json:"expiration_at"`
}
