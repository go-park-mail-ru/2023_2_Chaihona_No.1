package registration

import (
	"net/http"
	"sync"

	model "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/model"
)

type UserStorage struct {
	Users map[string]model.User
	Mu    sync.RWMutex
	Size  uint32
}

func CreateUserStorage() *UserStorage {
	storage := &UserStorage{
		Users: make(map[string]model.User),
		Mu:    sync.RWMutex{},
		Size:  0,
	}

	return storage
}

func (storage *UserStorage) RegisterNewUser(user *model.User) *ErrorRegistration {
	storage.Mu.Lock()
	defer storage.Mu.Unlock()

	_, ok := storage.Users[user.Login]

	if ok {
		return &ErrorRegistration{
			ErrUserLoginAlreadyExists,
			http.StatusBadRequest,
		}
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
	storage.Mu.RLock()
	defer storage.Mu.RUnlock()

	val, ok := storage.Users[login]

	if ok {
		copy := val

		return &copy, true
	}

	return nil, false
}

func (storage *UserStorage) GetUsers() ([]model.User, error) {
	storage.Mu.RLock()
	defer storage.Mu.RUnlock()

	users := make([]model.User, storage.Size)

	i := 0
	for _, user := range storage.Users {
		users[i] = user
		i++
	}

	return users, nil
}
