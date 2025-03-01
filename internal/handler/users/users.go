package users_handler

import (
	"log/slog"
	"net/http"
	"userservice/internal/service"
)

type UserHandler struct {
	log         *slog.Logger
	userService service.IUserService
}

func New(log *slog.Logger, userService service.IUserService) *UserHandler {
	return &UserHandler{
		log:         log,
		userService: userService,
	}
}

func (uc *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	panic("unimplement")
}

func (uc *UserHandler) GetById(w http.ResponseWriter, r *http.Request) {
	panic("unimplement")
}

func (uc *UserHandler) Insert(w http.ResponseWriter, r *http.Request) {
	panic("unimplement")
}

func (uc *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	panic("unimplement")
}

func (uc *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	panic("unimplement")
}
