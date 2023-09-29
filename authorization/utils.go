package authorization

import (
	"net/http"
	model "project/model"
	reg "project/registration"
)

func Authorize(users reg.UserRepository, form *LoginForm) (*model.User, error) {
	user, ok := users.CheckUser(form.Body_.Login)

	if !ok {
		return nil, ErrWrongLogin
	}

	if user.Password != form.Body_.Password {
		return nil, ErrWrongPassword
	}

	return user, nil
}

func CheckAuthorization(r *http.Request, sessions SessionRepository) bool {
	session, err := r.Cookie("session_id")

	if err == nil && session != nil {
		_, authorized := sessions.CheckSession(session.Value)

		return authorized
	}

	return false
}
