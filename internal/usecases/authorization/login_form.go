package authorization

// swager:parameters Login
type LoginForm struct {
	// in:body
	Body struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	} `json:"body"`
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
