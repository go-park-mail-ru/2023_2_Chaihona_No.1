package registration

import (
	"encoding/json"
	"io"
)

type BodySignUp struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	UserType string `json:"user_type"`
}

type SignupForm struct {
	Body_ BodySignUp `json:"body"`
}

func ParseJSON(r io.Reader) (*SignupForm, error) {
	decoder := json.NewDecoder(r)
	newUserInput := new(SignupForm)
	err := decoder.Decode(newUserInput)

	if err != nil {
		return nil, err
	}

	return newUserInput, nil
}
