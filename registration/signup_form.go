package registration

import (
	"encoding/json"
	"io"
	"project/model"
)

type BodySignUp struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	UserType string `json:"user_type"`
}

func (form BodySignUp) Validate() (*model.User, error) {
	isLenCorrect := len(form.Login) > 0 && len(form.Password) > 0
	isUserTypeCorrect := form.UserType == "simple_user" || form.UserType == "creator" 

	if isLenCorrect && isUserTypeCorrect {
		return &model.User{
			Login: form.Login,
			Password: form.Password,
			UserType: form.UserType,
		}, nil
	}

	return nil, ErrBodySignUpValidation
}

type SignupForm struct {
	Body_ BodySignUp `json:"body"`
}

func ParseJSON(r io.Reader) (*BodySignUp, error) {
	decoder := json.NewDecoder(r)
	newUserInput := &BodySignUp{}
	err := decoder.Decode(newUserInput)

	if err != nil {
		return nil, err
	}

	return newUserInput, nil
}
