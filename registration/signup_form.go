package registration

import (
	"encoding/json"
	"net/http"
)

type BodySignUp struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	UserType string `json:"user_type"`
}

type SignupForm struct {
	Body_ BodySignUp `json:"body"`
}

func ParseJSON(r *http.Request) (*SignupForm, error) {
	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)
	newUserInput := new(SignupForm)
	err := decoder.Decode(newUserInput)

	if err != nil {
		return nil, err
	}

	return newUserInput, nil
}
