package authorization

import (
	"errors"
)

var ErrNoSuchSession = errors.New("no such session")
var ErrWrongLogin = errors.New("wrong login")
var ErrWrongPassword = errors.New("wrong password")
