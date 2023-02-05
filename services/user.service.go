package services

import "example.mongo-api/models"

type UserService interface {
	CreateUser(*models.User) error
	GetUser(*string) (*models.User, error)
	GetAllUser() ([]*models.User, error)
	UpdateUser(*models.User) error
	DeleteUser(*string) error
}
