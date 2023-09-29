package registration

import (
	model "project/model"
)

type UserRepository interface {
	RegisterNewUser(user model.User) error
	DeleteUser(userId uint) error
	CheckUser(userId uint) (model.User, bool)
	GetUsers() ([]model.User, error)
}
