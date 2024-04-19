package repository

import (
	"context"
	"order-service/internals/models"

	"gorm.io/gorm"
)



type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository{
	return &OrderRepository{db: db}
}


func (r *OrderRepository) CreateOrder(ctx context.Context, order *models.Order) error{
	result := r.db.Create(order)

	if result.Error != nil{
		return result.Error
	}
	return nil
}