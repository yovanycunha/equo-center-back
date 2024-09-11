package repositories

import "equocenterback/pkg/models"

type ProfessionalRepository interface {
	CreateProfessional(*models.Professional) error
	GetProfessional(*string) (*models.Professional, error)
	GetAllProfessionals() ([]*models.Professional, error)
	UpdateProfessional(*models.ProfessionalUpdate) error
	DeleteProfessional(*string) error
}