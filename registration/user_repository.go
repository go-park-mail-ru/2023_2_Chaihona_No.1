package registration

import (
	model "project/model"
)

type UserRepository interface {
	RegisterNewUser(user *model.User) error
	DeleteUser(login string) error
	CheckUser(login string) (*model.User, bool)
	GetUsers() ([]model.User, error)
}
