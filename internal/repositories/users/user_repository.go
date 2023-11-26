package users

import (
	model "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"
)

type UserRepository interface {
	RegisterNewUser(user *model.User) (int, error)
	DeleteUser(id int) error
	CheckUser(login string) (*model.User, error)
	GetUser(id int) (model.User, error)
	GetUserWithSubscribers(id, visiterId int) (model.User, error)
	ChangeUser(user model.User) error
	ChangeUserDescription(description string, id int) error
	ChangeUserStatus(status string, id int) error
	GetUsers() ([]model.User, error)
	GetTopUsers(limit int) ([]model.User, error)
	Search(nickname string) ([]model.User, error)
}