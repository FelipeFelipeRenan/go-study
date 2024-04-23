package food

type FoodService struct {
	repo FoodRepository
}

func NewFoodService(repo FoodRepository) *FoodService {
	return &FoodService{repo: repo}
}


