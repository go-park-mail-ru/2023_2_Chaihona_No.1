package registration

import (
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/model"
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
			Login:    form.Login,
			Password: form.Password,
			UserType: form.UserType,
		}, nil
	}

	return nil, ErrBodySignUpValidation
}

type SignupForm struct {
	Body_ BodySignUp `json:"body"`
}
