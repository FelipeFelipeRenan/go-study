package food

type FoodService struct {
	repo FoodRepository
}

func NewFoodService(repo FoodRepository) *FoodService {
	return &FoodService{repo: repo}
}

func (s *FoodService) CreateFood(name, category string) (*Food, error){
	food := &Food{Name: name, Category: category}
	if err := s.repo.db.Create(food).Error; err != nil{
		return nil, err
	}
	return food, nil
}


func (s *FoodService) GetFoodByID(id uint32, )
