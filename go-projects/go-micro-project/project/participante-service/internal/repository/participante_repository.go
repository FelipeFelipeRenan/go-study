package repository

import (
	"context"
	"participante-service/internal/models"
	"gorm.io/gorm"
)

type ParticipanteRepository struct {
	db *gorm.DB
}

func NewParticipanteRepository(db *gorm.DB) *ParticipanteRepository {
	return &ParticipanteRepository{db: db}
}

func (r *ParticipanteRepository) CreateParticipante(ctx context.Context, Participante *models.Participante) error {
	result := r.db.Create(Participante)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *ParticipanteRepository) GetParticipanteByID(ctx context.Context, id int) (*models.Participante, error) {
	var Participante models.Participante
	result := r.db.First(&Participante, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &Participante, nil
}

func (r *ParticipanteRepository) UpdateParticipante(ctx context.Context, Participante *models.Participante) error {
	result := r.db.Save(Participante)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *ParticipanteRepository) DeleteParticipante(ctx context.Context, id int) error {
	result := r.db.Delete(&models.Participante{}, id)
	if result.Error != nil{
		return result.Error
	}
	return nil
}