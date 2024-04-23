package food

import "gorm.io/gorm"

type Food struct {
	gorm.Model
	Name string `json:"name"`
	Category string `json:"category"`
	Quantity uint32 `json:"quantity"`
}

