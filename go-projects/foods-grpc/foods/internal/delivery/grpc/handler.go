package grpc

import (
	"context"
	food "foods/internal/adapters"
	pb "foods/gen"
)


type FoodHandler struct {
	service food.FoodService
}

func NewFoodHandler(service food.FoodService) *FoodHandler  {
	return &FoodHandler{service: service}
}

func (h *FoodHandler)CreateFood(ctx context.Context, req *pb.CreateFoodRequest) (pb.FoodResponse, error)  {
	food, err := h.service.CreateFood(req.Name, req.Category)
}