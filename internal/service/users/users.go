package user_service

import (
	"context"
	"fmt"
	"log/slog"
	"userservice/internal/domain/models"
	"userservice/internal/storage"
	"userservice/pkg/logger/sl"
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

func (u UserService) Get(ctx context.Context) ([]models.User, error) {
	const op = "service.users.get"
	users, err := u.UserStorage.Get(ctx)
	if err != nil {
		u.log.Error("error to retrieve users", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return users, nil

}

func (u UserService) GetById(ctx context.Context, id int) (models.User, error) {
	const op = "service.users.getById"
	user, err := u.UserStorage.GetById(ctx, id)
	if err != nil {
		u.log.Error("error to retrieve user by id", sl.Err(err))
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}

func (u UserService) Insert(ctx context.Context, user models.User) error {
	const op = "service.users.insert"
	err := u.UserStorage.Insert(ctx, user)
	if err != nil {
		u.log.Error("error inserting user", sl.Err(err))
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (u UserService) Update(ctx context.Context, id int, user models.User) error {
	const op = "service.users.update"
	err := u.UserStorage.Update(ctx, id, user)
	if err != nil {
		u.log.Error("error updating user", sl.Err(err))
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (u UserService) Delete(ctx context.Context, id int) error {
	const op = "service.users.delete"
	err := u.UserStorage.Delete(ctx, id)
	if err != nil {
		u.log.Error("error deleting user", sl.Err(err))
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
