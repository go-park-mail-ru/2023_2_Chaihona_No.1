package users

import (
	"errors"
	"fmt"
)

var (
	ErrNoSuchUser             = errors.New("no such user")
	ErrUserLoginAlreadyExists = errors.New("user with this login already exists")
)

type ErrorUserRegistration struct {
	Err        error `json:"error"`
	StatusCode int
}

func (e ErrorUserRegistration) Error() string {
	return e.Err.Error() + fmt.Sprintf("with status code (%d)", e.StatusCode)
}
