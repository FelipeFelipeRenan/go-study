package models

import "gorm.io/gorm"

type Participante struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}
