package authorization

import (
	"encoding/json"
	"io"
)

type BodyLogin struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (form BodyLogin) Validate() bool {
	isLenCorrect := len(form.Login) > 0 && len(form.Password) > 0

	return isLenCorrect
}

type LoginForm struct {
	Body BodyLogin `json:"body"`
}

func ParseJSON(r io.Reader) (*BodyLogin, error) {
	decoder := json.NewDecoder(r)
	newUserInput := &BodyLogin{}
	err := decoder.Decode(newUserInput)
	if err != nil {
		return nil, err
	}

	return newUserInput, nil
}
