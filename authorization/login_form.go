package authorization

import (
	"encoding/json"
	"io"
)

type BodyLogin struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LoginForm struct {
	Body_ BodyLogin `json:"body"`
}

func ParseJSON(r io.Reader) (*BodyLogin, error) {
	decoder := json.NewDecoder(r)
	newUserInput := new(BodyLogin)
	err := decoder.Decode(newUserInput)

	if err != nil {
		return nil, err
	}

	return newUserInput, nil
}
