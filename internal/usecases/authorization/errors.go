package authorization

import (
	"errors"
)

var (
	ErrWrongLogin    = errors.New("wrong login")
	ErrWrongPassword = errors.New("wrong password")
)
