package models

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Location    string `json:"location"`
	StartAt     string `json:"start_at"`
	EndAt       string `json:"end_at"`
	// Adicionar outros campos conforme necessário
}
