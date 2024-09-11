package repositories

import "equocenterback/pkg/models"

type ProfessionalRepository interface {
	CreateProfessional(*models.Professional) error
}