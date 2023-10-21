package posts

import (
	"errors"
	"fmt"
)

var ErrNoSuchPost = errors.New("no such post")

type ErrorPost struct {
	Err        error `json:"error"`
	StatusCode int
}

func (e ErrorPost) Error() string {
	return e.Err.Error() + fmt.Sprintf("with status code (%d)", e.StatusCode)
}
