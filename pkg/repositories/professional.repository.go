package repositories

import "equocenterback/pkg/models"

type ProfessionalRepository interface {
	CreateProfessional(*models.Professional) error
	GetProfessional(*string) (*models.Professional, error)
}