package app

import (
	"log/slog"
	"userservice/internal/domain/models"
	"userservice/internal/service"
	user_service "userservice/internal/service/users"
	"userservice/internal/storage"
)

type HTTPApplication struct {
	log         *slog.Logger
	address     string
	userService service.IUserService
}

func New(log *slog.Logger, address string, userStorage storage.Repository[models.User]) *HTTPApplication {
	userService := user_service.New(log, userStorage)

	return &HTTPApplication{
		log:         log,
		address:     address,
		userService: userService,
	}
}
