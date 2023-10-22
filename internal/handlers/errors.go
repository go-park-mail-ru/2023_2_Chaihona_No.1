package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"encoding/json"

	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/posts"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/profiles"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/users"
)

var ErrorNotAuthor = errors.New("user isn't author")

type ErrorHttp struct {
	StatusCode int
	Msg        string
}

func (e ErrorHttp) Error() string {
	return e.Msg + fmt.Sprintf("(with status code %d)", e.StatusCode)
}

func WriteHttpError(w http.ResponseWriter, err error) {
	switch err.(type) {
	case ErrorHttp:
		err := err.(ErrorHttp)
		http.Error(w, err.Msg, err.StatusCode)

	case posts.ErrorPost:
		err := err.(posts.ErrorPost)
		jsonErr, _ := json.Marshal(err)
		http.Error(w, string(jsonErr), err.StatusCode)

	case users.ErrorUserRegistration:
		err := err.(users.ErrorUserRegistration)
		jsonErr, _ := json.Marshal(err)
		http.Error(w, string(jsonErr), err.StatusCode)
	
	case profiles.ErrorProfileRegistration:
		err := err.(profiles.ErrorProfileRegistration)
		jsonErr, _ := json.Marshal(err)
		http.Error(w, string(jsonErr), err.StatusCode)

	default:
		jsonErr, _ := json.Marshal(err)
		http.Error(w, string(jsonErr), http.StatusInternalServerError)
	}	
}
