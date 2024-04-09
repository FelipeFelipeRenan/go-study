package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	CustomerID uint    `json:"customer_id"`
	FoodID     uint    `json:"food_id"`
	Quantity   uint    `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
	Status     string  `json:"status"`
}
