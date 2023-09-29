package registration

import (
	model "project/model"
	"sync"
)

type UserStorage struct {
	Users map[uint]model.User
	Mu    sync.Mutex
	Size  uint32
}

func CreateUserStorage() *UserStorage {
	storage := &UserStorage{
		Users: make(map[uint]model.User),
	}

	return storage
}

func (storage *UserStorage) RegisterNewUser(user model.User) error {
	storage.Mu.Lock()
	defer storage.Mu.Unlock()

	storage.Size++
	storage.Users[user.ID] = user

	return nil
}

func (storage *UserStorage) DeleteUser(userId uint) error {
	storage.Mu.Lock()
	defer storage.Mu.Unlock()

	if _, ok := storage.Users[userId]; !ok {
		return ErrNoSuchUser
	}

	delete(storage.Users, userId)
	storage.Size--

	return nil
}

func (storage *UserStorage) CheckUser(userId uint) (*model.User, bool) {
	storage.Mu.Lock()
	defer storage.Mu.Unlock()

	val, ok := storage.Users[userId]

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
