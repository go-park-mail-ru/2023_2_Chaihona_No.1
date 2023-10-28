package users

import (
	"database/sql"
	"net/http"
	"sync"

	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/dbscan"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/configs"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"
)

func SelectUserSQL(login string) squirrel.SelectBuilder {
	return squirrel.Select("*").
		From(configs.UserTable).
		Where("email LIKE (?)", login).
		PlaceholderFormat(squirrel.Dollar)
}

type UserStorage struct {
	Users map[string]model.User
	Mu    sync.RWMutex
	db    *sql.DB
	Size  uint32
}

func CreateUserStorage(db *sql.DB) *UserStorage {
	return &UserStorage{
		db: db,
	}
}

func (storage *UserStorage) RegisterNewUser(user *model.User) error {
	storage.Mu.Lock()
	defer storage.Mu.Unlock()

	is_author := true
	if user.UserType == model.SimpleUserStatus {
		is_author = false
	}
	query := squirrel.Insert("public.user").
		Columns("nickname", "email", "password", "is_author", "status", "avatar_path", "background_path", "description").
		Values(user.Nickname, user.Login, user.Password, is_author, user.Status, user.Avatar, user.Background, user.Description).
		Suffix("RETURING \"id\"").
		RunWith(storage.db)
	var userId int
	err := query.QueryRow().Scan(userId)
	if err != nil {
		return ErrorUserRegistration{
			ErrUserLoginAlreadyExists,
			http.StatusBadRequest,
		}
	}
	//return userId
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
	rows, err := SelectUserSQL(login).RunWith(storage.db).Query()
	if err != nil {
		return nil, false
	}
	var users []*model.User
	err = dbscan.ScanAll(&users, rows)
	if err != nil {
		return nil, false
	}
	return users[0], true
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
