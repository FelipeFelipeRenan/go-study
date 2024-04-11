package repository

import (
	"context"
	"errors"
	"foods/internal/models"
	"log"

	"gorm.io/gorm"
)

type FoodRepository struct {
	db *gorm.DB
}

func NewFoodRepository(db *gorm.DB) *FoodRepository {
	return &FoodRepository{db: db}
}

func (r *FoodRepository) CreateFood(ctx context.Context, Food *models.Food) error {
	result := r.db.Create(Food)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *FoodRepository) GetAllFoods(ctx context.Context) ([]*models.Food, error) {
	var foods []*models.Food
	result := r.db.Find(&foods)
	if result.Error != nil {
		return nil, result.Error
	}
	return foods, nil
}

func (r *FoodRepository) GetAllFoodsByCategory(ctx context.Context, category string) ([]*models.Food, error) {
	var foods []*models.Food
	result := r.db.Where("category = ?", category).Find(&foods)
	if result.Error != nil {
		return nil, result.Error
	}
	return foods, nil
}

func (r *FoodRepository) GetFoodsByID(ctx context.Context, id int) (*models.Food, error) {
	var Food models.Food
	result := r.db.First(&Food, id)
	if result.Error != nil {
		return nil, result.Error

	}
	return &Food, nil
}

func (r *FoodRepository) UpdateFood(ctx context.Context, Food *models.Food) error {
	result := r.db.Save(Food)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *FoodRepository) DeleteFood(ctx context.Context, id int) error {
	result := r.db.Delete(&models.Food{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *FoodRepository) UpdateFoodQuantity(ctx context.Context, id int, quantity int) error {
	var food models.Food

	result := r.db.First(&food, id)
	log.Println(food.Name)
	if result.Error != nil {
		return result.Error
	}
	if food.Quantity >= int(quantity) {
		food.Quantity -= int(quantity)
	}else{
		return errors.New("Quantidade de comida insuficiente")
	}
	if err := r.db.Save(&food).Error; err != nil {
		return err
	}
	return nil
}
