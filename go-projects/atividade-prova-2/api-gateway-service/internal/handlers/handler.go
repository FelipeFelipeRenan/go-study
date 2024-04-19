package handlers

import (
	"api-gateway/internal/service"
	"context"

	"github.com/go-kit/kit/endpoint"
)

func MakeGetAllFoodsEndpoit(s service.FoodService) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (interface{},error){
		foods, err := s.GetAllFoods(ctx)
		return foods, err
	}
}  