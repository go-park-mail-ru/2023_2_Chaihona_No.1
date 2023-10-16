package authorization

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
