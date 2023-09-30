package authorization

import (
	"fmt"
	"net/http"
	model "project/model"
	reg "project/registration"
	"time"
)

const (
	SID_LEN      = 32
	TTL_DURATION = 10 * time.Hour
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
		fmt.Println(session)
		return authorized
	}

	return false
}

func SetSession(w http.ResponseWriter, sessions SessionRepository, userId uint32) {
	SID := RandStringRunes(SID_LEN)
	TTL := time.Now().Add(TTL_DURATION)

	sessions.RegisterNewSession(Session{
		SessionId: SID,
		UserId:    userId,
		Ttl:       TTL,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    SID,
		Expires:  TTL,
		HttpOnly: true,
	})
}
