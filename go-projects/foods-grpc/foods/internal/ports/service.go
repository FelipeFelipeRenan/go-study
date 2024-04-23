package ports

type FoodService interface {
	Create(name, category string) (*Food , error)
	GetAllFood()([]*Food, error)
	GetFoodByID(id uint32) (*Food, error)
	GetAllFoodByCategory(category string)([]*Food, error)
	UpdateFood(id uint32, name, category string)(*Food, error)
	DeleteFood(id uint32) error 
}