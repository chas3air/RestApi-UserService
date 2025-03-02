package users_handler

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"userservice/internal/domain/models"
	"userservice/internal/service"
	"userservice/pkg/logger/sl"

	"github.com/gorilla/mux"
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
	const op = "handler.users.get"
	uc.log.Info("/users (Get) is running...")

	users, err := uc.userService.Get(r.Context())
	if err != nil {
		uc.log.Warn("Failed to retrieve users", sl.Err(err))
		http.Error(w, fmt.Errorf("%s: %w", op, err).Error(), http.StatusInternalServerError)
		return
	}

	bs, err := json.Marshal(users)
	if err != nil {
		uc.log.Warn("Failed to marshal users to json", sl.Err(err))
		http.Error(w, fmt.Errorf("%s: %w", op, err).Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bs)
	uc.log.Info("/users (Get) done")
}

func (uc *UserHandler) GetById(w http.ResponseWriter, r *http.Request) {
	const op = "handler.users.getById"
	uc.log.Info("/users/{id} (Get) is running...")

	id_s, ok := mux.Vars(r)["id"]
	if !ok {
		uc.log.Error("error getting id")
		http.Error(w, fmt.Errorf("%s: %s", op, "error getting id").Error(), http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(id_s)
	if err != nil {
		uc.log.Error("id must be int")
		http.Error(w, fmt.Errorf("%s: %s", op, "id must be int").Error(), http.StatusBadRequest)
		return
	}

	user, err := uc.userService.GetById(r.Context(), id)
	if err != nil {
		uc.log.Warn("Failed to retrieve user by id", sl.Err(err))
		http.Error(w, fmt.Errorf("%s: %w", op, err).Error(), http.StatusInternalServerError)
		return
	}

	bs, err := json.Marshal(user)
	if err != nil {
		uc.log.Warn("Failed to marshal user to json", sl.Err(err))
		http.Error(w, fmt.Errorf("%s: %w", op, err).Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bs)
	uc.log.Info("/users/{id} (Get) done")
}

func (uc *UserHandler) Insert(w http.ResponseWriter, r *http.Request) {
	const op = "handler.users.insert"
	uc.log.Info("/users (Post) is running...")

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		uc.log.Warn("Failed to decode user from json", sl.Err(err))
		http.Error(w, fmt.Errorf("%s: %w", op, err).Error(), http.StatusBadRequest)
		return
	}

	err = uc.userService.Insert(r.Context(), user)
	if err != nil {
		uc.log.Warn("Failed to insert user", sl.Err(err))
		http.Error(w, fmt.Errorf("%s: %w", op, err).Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	uc.log.Info("/users (Post) done")
}

func (uc *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	const op = "handler.users.update"
	uc.log.Info("/users/{id} (Put) is running...")

	id_s, ok := mux.Vars(r)["id"]
	if !ok {
		uc.log.Error("error getting id")
		http.Error(w, fmt.Errorf("%s: %s", op, "error getting id").Error(), http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(id_s)
	if err != nil {
		uc.log.Error("id must be int")
		http.Error(w, fmt.Errorf("%s: %s", op, "id must be int").Error(), http.StatusBadRequest)
		return
	}

	var user models.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		uc.log.Warn("Failed to decode user from json", sl.Err(err))
		http.Error(w, fmt.Errorf("%s: %w", op, err).Error(), http.StatusBadRequest)
		return
	}

	err = uc.userService.Update(r.Context(), id, user)
	if err != nil {
		uc.log.Warn("Failed to update user", sl.Err(err))
		http.Error(w, fmt.Errorf("%s: %w", op, err).Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	uc.log.Info("/users/{id} (Put) done")
}

func (uc *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	const op = "handler.users.delete"
	uc.log.Info("/users/{id} (Delete) is running...")

	id_s, ok := mux.Vars(r)["id"]
	if !ok {
		uc.log.Error("error getting id")
		http.Error(w, fmt.Errorf("%s: %s", op, "error getting id").Error(), http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(id_s)
	if err != nil {
		uc.log.Error("id must be int")
		http.Error(w, fmt.Errorf("%s: %s", op, "id must be int").Error(), http.StatusBadRequest)
		return
	}

	err = uc.userService.Delete(r.Context(), id)
	if err != nil {
		uc.log.Warn("Failed to delete user", sl.Err(err))
		http.Error(w, fmt.Errorf("%s: %w", op, err).Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	uc.log.Info("/users/{id} (Delete) done")
}
