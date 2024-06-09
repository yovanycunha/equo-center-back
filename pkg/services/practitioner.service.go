package services

import "equocenterback/pkg/models"

type PractitionerService interface {
	CreatePractitioner(*models.Practitioner) error
}