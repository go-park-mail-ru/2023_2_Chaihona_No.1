package authorization

import (
	"errors"
)

var ErrNoSuchSession = errors.New("No such session!")
var ErrWrongLogin = errors.New("Wrong login!")
var ErrWrongPassword = errors.New("Wrong password!")
