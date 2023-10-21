package profiles

import (
	"errors"
	"fmt"
)

var (
	ErrNoSuchProfile = errors.New("no such profile")
)

type ErrorProfileRegistration struct {
	Err        error
	StatusCode int
}

func (e ErrorProfileRegistration) Error() string {
	return e.Err.Error() + fmt.Sprintf("with status code (%d)", e.StatusCode)
}
