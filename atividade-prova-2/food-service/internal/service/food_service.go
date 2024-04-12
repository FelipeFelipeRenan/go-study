package service

import (
	"context"
	"foods/internal/models"
	"foods/internal/repository"
)

type Service interface {
	GetAllFoods(ctx context.Context) ([]models.Food, error)
}

type foodService struct {
	repo repository.FoodRepository
}

func NewFoodService(repo repository.FoodRepository) Service {
	return &foodService{
		repo: repo,
	}
}

func (s *foodService) GetAllFoods(ctx context.Context) ([]models.Food, error) {
	// Lógica para obter todos os alimentos do repositório.
	foods, err := s.repo.GetAllFoods(ctx)
	if err != nil {
		return nil, err
	}
	// Convertendo []*models.Food para []models.Food
	var result []models.Food
	for _, f := range foods {
		result = append(result, *f)
	}
	return result, nil
}
