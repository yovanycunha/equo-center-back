package repositories

import "equocenterback/pkg/models"

type UserRepository interface {
	CreateUser(*models.User) error
	GetUser(*string) (*models.User, error)
}