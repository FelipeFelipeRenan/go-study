package repository

import (
	"context"
	"foods/internal/models"
	"log"

	"gorm.io/gorm"
)

type FoodRepository interface {
	GetAllFoods(ctx context.Context) ([]*models.Food, error)
}

type foodRepository struct {
	db *gorm.DB
}

func NewFoodRepository(db *gorm.DB) FoodRepository {
	return &foodRepository{
		db: db,
	}
}

func (r *foodRepository) GetAllFoods(ctx context.Context) ([]*models.Food, error) {
	var foods []*models.Food
	if err := r.db.Find(&foods).Error; err != nil {
		log.Println("Failed to get all foods", err)
	}
	return foods, nil
}
