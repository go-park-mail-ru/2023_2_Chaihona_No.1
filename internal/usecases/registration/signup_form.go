package registration

import (
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"
)

// swagger:parameters SignUp
type SignupForm struct {
	// in:body
	Body struct {
		Login    string `json:"login"`
		Password string `json:"password"`
		UserType string `json:"-"`
		IsAuthor bool `json:"isAuthor"`
	} `json:"body"`
}

func (form SignupForm) Validate() (*model.User, error) {
	isLenCorrect := len(form.Body.Login) > 0 && len(form.Body.Password) > 0
	// isUserTypeCorrect := form.Body.UserType == "simple_user" || form.Body.UserType == "creator"
	isUserTypeCorrect := true
	if isLenCorrect && isUserTypeCorrect {
		return &model.User{
			Login:    form.Body.Login,
			Password: form.Body.Password,
			UserType: form.Body.UserType,
		}, nil
	}

	return nil, ErrBodySignUpValidation
}

func (form SignupForm) IsValide() bool {
	isLenCorrect := len(form.Body.Login) > 0 && len(form.Body.Password) > 0
	// isUserTypeCorrect := form.Body.UserType == "simple_user" || form.Body.UserType == "creator"
	isUserTypeCorrect := true
	return isLenCorrect && isUserTypeCorrect
}

func (form SignupForm) IsEmpty() bool {
	return false
}