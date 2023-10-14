package authorization

import (
	"errors"
)

var (
	ErrNoSuchSession = errors.New("no such session")
	ErrWrongLogin    = errors.New("wrong login")
	ErrWrongPassword = errors.New("wrong password")
)
