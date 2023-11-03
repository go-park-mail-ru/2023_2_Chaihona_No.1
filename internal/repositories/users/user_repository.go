package users

import (
	model "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"
)

type UserRepository interface {
	RegisterNewUser(user *model.User) (int, error)
	DeleteUser(id int) error
	CheckUser(login string) (*model.User, error)
	GetUser(id int) (model.User, error)
	GetUserWithSubscribers(id int) (model.User, error)
	ChangeUser(user model.User) error
	GetUsers() ([]model.User, error)
}
