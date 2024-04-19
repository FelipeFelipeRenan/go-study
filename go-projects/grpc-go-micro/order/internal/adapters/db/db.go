package db

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	CustomerID int64
	Status string
	OrderItem []OrderItem
}

type OrderItem struct {
	gorm.Model
	ProductCode string
	UnitPrice  float32
	Quantity int32
	OrderID uint
}

type Adapter struct {
	db *gorm.DB	
}