package service

import "userservice/internal/domain/models"

type IUserService interface {
	Get() ([]models.User, error)
	GetById(int) (models.User, error)
	Insert(models.User) error
	Update(int, models.User) error
	Delte(int) error
}
