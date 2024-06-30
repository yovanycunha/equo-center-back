package repositories

import "equocenterback/pkg/models"

// RepositoryService
type PractitionerRepository interface {
	CreatePractitioner(*models.Practitioner) error
	GetPractitioner(*string) (*models.Practitioner, error)
	GetAllPractitioners() ([]*models.Practitioner, error)
	UpdatePractitioner(*models.Practitioner) error
	DeletePractitioner(*string) error
}