package authorization

type LoginForm struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (form LoginForm) Validate() bool {
	return len(form.Login) > 0 && len(form.Password) > 0
}

func (form LoginForm) IsValide() bool {
	return len(form.Login) > 0 && len(form.Password) > 0
}

func (form LoginForm) IsEmpty() bool {
	return false
}
