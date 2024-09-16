package repositories

import "equocenterback/pkg/models"

type ActivityRepository interface {
	CreateActivity(*models.Activity) error
	GetActivity(*string) (*models.Activity, error)
	GetAllActivities() ([]*models.Activity, error)
	UpdateActivity(*models.Activity) error
	DeleteActivity(*string) error
}