package users

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/dbscan"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/configs"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"
)

func InsertUserSQL(user model.User) squirrel.InsertBuilder {
	return squirrel.Insert(configs.UserTable).
		Columns("nickname", "email", "password", "is_author", "status", "description").
		Values(user.Nickname, user.Login, user.Password, user.Is_author, user.Status, user.Description).
		Suffix("RETURNING \"id\"").
		PlaceholderFormat(squirrel.Dollar)
}

func DeleteUserSQL(id int) squirrel.DeleteBuilder {
	return squirrel.Delete(configs.UserTable).
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar)
}

func SelectUserSQL(login string) squirrel.SelectBuilder {
	return squirrel.Select("*").
		From(configs.UserTable).
		Where("email LIKE (?)", login).
		PlaceholderFormat(squirrel.Dollar)
}

func SelectUserByIdSQL(id int) squirrel.SelectBuilder {
	return squirrel.Select("*").
		From(configs.UserTable).
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar)
}

func SelectUserByIdSQLWithSubscribers(id int) squirrel.SelectBuilder {
	return squirrel.Select(
		fmt.Sprintf("%s.id, %s.nickname, %s.email, %s.is_author, %s.status, %s.avatar_path, %s.background_path, %s.description, COUNT(*) as subscribers",
			configs.UserTable, configs.UserTable, configs.UserTable,
			configs.UserTable, configs.UserTable, configs.UserTable,
			configs.UserTable, configs.UserTable)).
		From(configs.UserTable).
		LeftJoin(fmt.Sprintf("%s s ON %s.id = s.creator_id", configs.SubscriptionTable, configs.UserTable)).
		Suffix(fmt.Sprintf("WHERE %s.id = %d", configs.UserTable, id)).
		Suffix("GROUP BY " + configs.UserTable + ".id")
}

func UpdateUserSQL(user model.User) squirrel.UpdateBuilder {
	return squirrel.Update(configs.UserTable).
		SetMap(map[string]interface{}{
			"email":           user.Login,
			"nickname":        user.Nickname,
			"password":        user.Password,
			// "is_author":       user.Is_author,
			"status":          user.Status,
			"avatar_path":     user.Avatar,
			"background_path": user.Background,
			"description":     user.Description,
			//update_at: time.Now(),
		}).
		Where(squirrel.Eq{"id": user.ID}).
		PlaceholderFormat(squirrel.Dollar)
}

type UserStorage struct {
	db *sql.DB
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
		return nil, ErrNoSuchUser
	}
	return users[0], nil
}

func (storage *UserStorage) GetUser(id int) (model.User, error) {
	rows, err := SelectUserByIdSQL(id).RunWith(storage.db).Query()
	if err != nil {
		return model.User{}, err
	}
	var users []*model.User
	err = dbscan.ScanAll(&users, rows)
	if err != nil {
		return model.User{}, err
	}
	return *users[0], nil
}

func (storage *UserStorage) GetUserWithSubscribers(id int) (model.User, error) {
	rows, err := SelectUserByIdSQLWithSubscribers(id).RunWith(storage.db).Query()
	if err != nil {
		return model.User{}, err
	}
	var users []model.User
	err = dbscan.ScanAll(&users, rows)
	if err != nil {
		return model.User{}, err
	}
	if len(users) > 0 {
		return users[0], nil
	}
	return model.User{}, nil
}

func (storage *UserStorage) ChangeUser(user model.User) error {
	_, err := UpdateUserSQL(user).RunWith(storage.db).Query()
	if err != nil {
		return err
	}
	return nil
}

func (storage *UserStorage) GetUsers() ([]model.User, error) {
	return make([]model.User, 0), nil
}