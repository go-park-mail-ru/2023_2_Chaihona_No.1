package registration

import (
	model "project/model"
	"sync"
)

type UserStorage struct {
	Users map[string]model.User
	Mu    sync.Mutex
	Size  uint32
}

func CreateUserStorage() *UserStorage {
	storage := &UserStorage{
		Users: make(map[string]model.User),
	}

	return storage
}

func (storage *UserStorage) RegisterNewUser(user *model.User) error {
	storage.Mu.Lock()
	defer storage.Mu.Unlock()

	_, ok := storage.Users[user.Login]

	if ok {
		return ErrUserLoginAlreadyExists
	}

	user.ID = uint(storage.Size)
	storage.Size++
	storage.Users[user.Login] = *user

	return nil
}

func (storage *UserStorage) DeleteUser(login string) error {
	storage.Mu.Lock()
	defer storage.Mu.Unlock()

	if _, ok := storage.Users[login]; !ok {
		return ErrNoSuchUser
	}

	delete(storage.Users, login)
	storage.Size--

	return nil
}

func (storage *UserStorage) CheckUser(login string) (*model.User, bool) {
	storage.Mu.Lock()
	defer storage.Mu.Unlock()

	val, ok := storage.Users[login]

	if ok {
		copy := val

		return &copy, true
	}

	return nil, false
}

func (storage *UserStorage) GetUsers() ([]model.User, error) {
	storage.Mu.Lock()
	defer storage.Mu.Unlock()

	users := make([]model.User, 0)

	for _, user := range storage.Users {
		users = append(users, user)
	}

	return users, nil
}
