package users

import (
	model "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"
)

type UserRepository interface {
	RegisterNewUser(user *model.User) error
	DeleteUser(login string) error
	CheckUser(login string) (*model.User, bool)
	GetUsers() ([]model.User, error)
}
