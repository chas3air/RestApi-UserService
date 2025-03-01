package user_service

import (
	"log/slog"
	"userservice/internal/domain/models"
	"userservice/internal/storage"
)

type UserService struct {
	log         *slog.Logger
	UserStorage storage.Repository[models.User]
}

func New(log *slog.Logger, userStorage storage.Repository[models.User]) *UserService {
	return &UserService{
		log:         log,
		UserStorage: userStorage,
	}
}

func (u UserService) Get() ([]models.User, error) {
	panic("unimplemented")
}

func (u UserService) GetById(int) (models.User, error) {
	panic("unimplemented")
}

func (u UserService) Insert(models.User) error {
	panic("unimplemented")
}

func (u UserService) Update(int, models.User) error {
	panic("unimplemented")
}

func (u UserService) Delete(int) error {
	panic("unimplemented")
}
