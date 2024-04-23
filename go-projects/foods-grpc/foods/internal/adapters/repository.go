package food

import "gorm.io/gorm"

type FoodRepository struct {
	db *gorm.DB
}

func NewFoodRepository(db *gorm.DB) *FoodRepository  {
	return &FoodRepository{db: db}
}

func (r *FoodRepository)Create(food *Food) (*Food, error){
	if err := r.db.Create(food).Error; err != nil{
		return nil, err
	}
	return food, nil
}

