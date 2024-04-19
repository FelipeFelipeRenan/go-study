package service

import "context"

type FoodService interface {
	GetAllFoods(ctx context.Context) ([]Food, error)
}

type OrderService interface {
	GetOrders(ctc context.Context) ([]Order, error)
}

