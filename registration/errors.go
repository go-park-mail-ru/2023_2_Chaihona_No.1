package registration

import (
	"errors"
)

var ErrNoSuchUser = errors.New("No such user!")
var ErrNoSuchProfile = errors.New("No such profile!")
