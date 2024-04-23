package ports

type FoodRepository interface {
	Create(food *Food) (*Food, error)
	GetAll()([] *Food, error)
	GetByID(uint32) (*Food, error)
	Update(food *Food)(*Food, error)
	Delete(id uint32) error
}