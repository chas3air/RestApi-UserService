package app

import (
	"fmt"
	"log/slog"
	"net/http"
	"userservice/internal/domain/models"
	users_handler "userservice/internal/handler/users"
	"userservice/internal/service"
	user_service "userservice/internal/service/users"
	"userservice/internal/storage"

	"github.com/gorilla/mux"
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

func (s *HTTPApplication) Start() error {
	const op = "app.Start"

	users_handler := users_handler.New(s.log, s.userService)

	r := mux.NewRouter()
	r.HandleFunc("/users", users_handler.Get).Methods(http.MethodGet)
	r.HandleFunc("/users/{id}", users_handler.GetById).Methods(http.MethodGet)
	r.HandleFunc("/users", users_handler.Insert).Methods(http.MethodPost)
	r.HandleFunc("/users/{id}", users_handler.Update).Methods(http.MethodPut)
	r.HandleFunc("/users/{id}", users_handler.Delete).Methods(http.MethodDelete)

	if err := http.ListenAndServe(s.address, r); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
