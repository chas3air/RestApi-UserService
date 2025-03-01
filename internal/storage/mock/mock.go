package memory

import (
	"errors"
	"sync"
	"userservice/internal/domain/models"
)

type UserRepo struct {
	mu     sync.Mutex
	users  map[int]models.User
	nextID int
}

func New() *UserRepo {
	return &UserRepo{
		users:  make(map[int]models.User),
		nextID: 1,
	}
}

func (ur *UserRepo) Get() ([]models.User, error) {
	ur.mu.Lock()
	defer ur.mu.Unlock()

	usersList := make([]models.User, 0, len(ur.users))
	for _, user := range ur.users {
		usersList = append(usersList, user)
	}
	return usersList, nil
}

func (ur *UserRepo) GetById(id int) (models.User, error) {
	ur.mu.Lock()
	defer ur.mu.Unlock()

	user, exists := ur.users[id]
	if !exists {
		return models.User{}, errors.New("user not found")
	}
	return user, nil
}

func (ur *UserRepo) Insert(user models.User) error {
	ur.mu.Lock()
	defer ur.mu.Unlock()

	user.Id = ur.nextID
	ur.users[ur.nextID] = user
	ur.nextID++
	return nil
}

func (ur *UserRepo) Update(id int, user models.User) error {
	ur.mu.Lock()
	defer ur.mu.Unlock()

	if _, exists := ur.users[id]; !exists {
		return errors.New("user not found")
	}
	user.Id = id
	ur.users[id] = user
	return nil
}

func (ur *UserRepo) Delete(id int) error {
	ur.mu.Lock()
	defer ur.mu.Unlock()

	if _, exists := ur.users[id]; !exists {
		return errors.New("user not found")
	}
	delete(ur.users, id)
	return nil
}
