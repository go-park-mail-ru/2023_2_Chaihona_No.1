package authorization

import (
	"encoding/json"
	"net/http"
)

type BodyLogin struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginForm struct {
	Body_ BodyLogin `json:"body"`
}

func ParseJSON(r *http.Request) (*LoginForm, error) {
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	newUserInput := new(LoginForm)
	err := decoder.Decode(newUserInput)

	if err != nil {
		return nil, err
	}

	return newUserInput, nil
}
