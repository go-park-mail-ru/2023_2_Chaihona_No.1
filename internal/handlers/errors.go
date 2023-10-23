package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

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
	return e.Msg + fmt.Sprintf(" with status code (%d)", e.StatusCode)
}

var (
	ErrValidation = ErrorHttp{
		StatusCode: http.StatusBadRequest,
		Msg:        `{"error":"user_validation"}`,
	}

	ErrDecoding = ErrorHttp{
		StatusCode: http.StatusBadRequest,
		Msg:        `{"error":"wrong_json"}`,
	}

	ErrEncoding = ErrorHttp{
		StatusCode: http.StatusInternalServerError,
		Msg:        `{"error":"encoding_json"}`,
	}

	ErrLogoutCookie = ErrorHttp{
		StatusCode: http.StatusBadRequest,
		Msg:        `{"error":"user_logout_no_cookie"}`,
	}

	ErrLogoutDeleteSession = ErrorHttp{
		StatusCode: http.StatusInternalServerError,
		Msg:        `{"error":"user_logout_cant_delete_session"}`,
	}

	ErrUnathorized = ErrorHttp{
		StatusCode: http.StatusUnauthorized,
		Msg:        `{"error":"user_unauthorized"}`,
	}
)

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
