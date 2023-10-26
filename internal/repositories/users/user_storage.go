package users

import (
	"database/sql"
	"fmt"
	"net/http"
	"sync"

	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/dbscan"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/configs"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"
)

func InsertUserSQL(user model.User) squirrel.InsertBuilder {
	is_author := true
	if user.UserType == model.SimpleUserStatus {
		is_author = false
	}
	return squirrel.Insert(configs.UserTable).
		Columns("nickname", "email", "password", "is_author", "status", "avatar_path", "background_path", "description").
		Values(user.Nickname, user.Login, user.Password, is_author, user.Status, user.Avatar, user.Background, user.Description).
		Suffix("RETURNING \"id\"").
		PlaceholderFormat(squirrel.Dollar)
}

func DeleteUserSQL(id int) squirrel.DeleteBuilder {
	return squirrel.Delete(configs.UserTable).
		Where(squirrel.Eq{"id": id})
}

func SelectUserSQL(login string) squirrel.SelectBuilder {
	return squirrel.Select("*").
		From(configs.UserTable).
		Where("email LIKE (?)", login).
		PlaceholderFormat(squirrel.Dollar)
}

func UpdateUserSQL(user model.User) squirrel.UpdateBuilder {
	return squirrel.Update(configs.UserTable).
		SetMap(map[string]interface{}{
			"email":           user.Login,
			"nickname":        user.Nickname,
			"password":        user.Password,
			"is_author":       user.Is_author,
			"status":          user.Status,
			"avatar_path":     user.Avatar,
			"background_path": user.Background,
			"description":     user.Description,
			//update_at: time.Now(),
		}).
		Where(squirrel.Eq{"id": user.ID})
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

func (storage *UserStorage) RegisterNewUser(user *model.User) (int, error) {
	var userId int
	err := InsertUserSQL(*user).RunWith(storage.db).QueryRow().Scan(&userId)
	if err != nil {
		return 0, ErrorUserRegistration{
			ErrUserLoginAlreadyExists,
			http.StatusBadRequest,
		}
	}
	return userId, nil
}

func (storage *UserStorage) DeleteUser(id int) error {
	_, err := DeleteUserSQL(id).RunWith(storage.db).Query()
	if err != nil {
		return ErrorUserRegistration{
			ErrNoSuchUser,
			http.StatusBadRequest,
		}
	}
	return nil
}

func (storage *UserStorage) CheckUser(login string) (*model.User, error) {
	rows, err := SelectUserSQL(login).RunWith(storage.db).Query()
	if err != nil {
		return nil, err
	}
	var users []*model.User
	err = dbscan.ScanAll(&users, rows)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return users[0], nil
}

func (storage *UserStorage) ChangeUser(user model.User) error {
	_, err := UpdateUserSQL(user).RunWith(storage.db).Query()
	if err != nil {
		return err
	}
	return nil
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
