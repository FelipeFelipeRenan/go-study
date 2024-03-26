package repository

import (
	"context"
	"evento-service/internal/models"
	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) *EventRepository {
	return &EventRepository{db: db}
}

func (r *EventRepository) CreateEvent(ctx context.Context, event *models.Event) error {
	result := r.db.Create(event)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *EventRepository) GetEventByID(ctx context.Context, id int) (*models.Event, error) {
	var event models.Event
	result := r.db.First(&event, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &event, nil
}
