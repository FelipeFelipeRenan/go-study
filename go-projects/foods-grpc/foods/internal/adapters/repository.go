package food

import (
	"context"
	"strings"

	"gorm.io/gorm"
)

type FoodRepository struct {
	db *gorm.DB
}

func NewFoodRepository(db *gorm.DB) *FoodRepository {
	return &FoodRepository{db: db}
}

func (r *FoodRepository) Create(ctx context.Context, food *Food) (*Food, error) {
	if err := r.db.Create(food).Error; err != nil {
		return nil, err
	}
	return food, nil
}

func (r *FoodRepository) GetAll(ctx context.Context) ([]*Food, error) {
	var foods []*Food
	if err := r.db.Find(&foods).Error; err != nil {
		return nil, err
	}

	return foods, nil

}

func (r *FoodRepository) GetByID(ctx context.Context, id uint32) (*Food, error) {
	var food Food
	if err := r.db.First(&food, id).Error; err != nil {
		return nil, err
	}
	return &food, nil
}

func (r *FoodRepository) GetAllFoodByCategory(ctx context.Context, category string) ([]*Food, error) {
	var foods []*Food
	categoryNormalized := strings.Title(strings.ToLower(category))
	if err := r.db.Where("category = ?", categoryNormalized).Find(&foods).Error; err != nil {
		return nil, err
	}
	return foods, nil
}

func (r *FoodRepository) Update(ctx context.Context, food *Food) (*Food, error) {
	if err := r.db.Save(food).Error; err != nil {
		return nil, err
	}

	return food, nil
}

func (r *FoodRepository) Delete(ctx context.Context, id uint32) error {
	if err := r.db.Delete(&Food{}, id).Error; err != nil {
		return err
	}
	return nil
}
