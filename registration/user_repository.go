package registration

import (
	model "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/model"
)

type UserRepository interface {
	RegisterNewUser(user *model.User) *ErrorRegistration
	DeleteUser(login string) error
	CheckUser(login string) (*model.User, bool)
	GetUsers() ([]model.User, error)
}
