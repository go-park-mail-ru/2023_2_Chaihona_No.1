package registration

import (
	"errors"
)

var ErrNoSuchUser = errors.New("No such user!")
var ErrNoSuchProfile = errors.New("No such profile!")
var ErrUserLoginAlreadyExists = errors.New("User with this login already exists")
var ErrBodySignUpValidation = errors.New("error during validation BodySignUp form")

type ErrorRegistration struct {
	Err        error
	StatusCode int
}
