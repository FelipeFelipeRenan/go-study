package ports

import "context"

type FoodRepository interface {
	Create(ctx context.Context, food *Food) (*Food, error)
	GetAll(ctx context.Context) ([]*Food, error)
	GetByID(ctx context.Context, id uint32) (*Food, error)
	GetAllFoodByCategory(ctx context.Context, category string) ([]*Food, error)
	Update(ctx context.Context, food *Food) (*Food, error)
	Delete(ctx context.Context, id uint32) error
}
