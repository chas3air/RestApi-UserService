package service

import (
	"context"
	"userservice/internal/domain/models"
)

type IUserService interface {
	Get(context.Context) ([]models.User, error)
	GetById(context.Context, int) (models.User, error)
	Insert(context.Context, models.User) error
	Update(context.Context, int, models.User) error
	Delete(context.Context, int) error
}
