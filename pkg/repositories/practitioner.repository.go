package repositories

import "equocenterback/pkg/models"

// RepositoryService
type PractitionerRepository interface {
	CreatePractitioner(*models.Practitioner) error
}