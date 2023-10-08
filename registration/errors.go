package registration

import (
	"errors"
)

var ErrNoSuchUser = errors.New("no such user")
var ErrNoSuchProfile = errors.New("no such profile")
var ErrUserLoginAlreadyExists = errors.New("user with this login already exists")
var ErrBodySignUpValidation = errors.New("error during validation BodySignUp form")

type ErrorRegistration struct {
	Err        error
	StatusCode int
}
