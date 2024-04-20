package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

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

func NewAdapter(dataSourceUrl string)(*Adapter, error)  {
	db, openErr := gorm.Open(mysql.Open(dataSourceUrl), &gorm.Config{})
	if openErr != nil{
		return nil, fmt.Errorf("db connection error: %v", openErr)
	}
	
	err := db.AutoMigrate(&Order{}, OrderItem{})
	if err != nil {
		return nil, fmt.Errorf("db migration error: %v", err)
	}
	return &Adapter{db: db}, nil
}