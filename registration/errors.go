package registration

import (
	"errors"
)

var (
	ErrNoSuchUser             = errors.New("no such user")
	ErrNoSuchProfile          = errors.New("no such profile")
	ErrUserLoginAlreadyExists = errors.New("user with this login already exists")
	ErrBodySignUpValidation   = errors.New("error during validation BodySignUp form")
)

type ErrorRegistration struct {
	Err        error
	StatusCode int
}
