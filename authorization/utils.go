package authorization

import (
	"net/http"
	model "project/model"
	reg "project/registration"
	"time"

	"github.com/google/uuid"
)

const (
	SID_LEN      = 32
	TTL_DURATION = 10 * time.Hour
)

func Authorize(users reg.UserRepository, form *LoginForm) (*model.User, error) {
	user, ok := users.CheckUser(form.Body.Login)

	if !ok {
		return nil, ErrWrongLogin
	}

	if user.Password != form.Body.Password {
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

func SetSession(w http.ResponseWriter, sessions SessionRepository, userId uint32) {
	SID := uuid.New().String()
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
		SameSite: http.SameSiteLaxMode,
		Secure:   false,
	})
}

func RemoveSession(w http.ResponseWriter, sessions SessionRepository, sessionId string) error {
	EXPIRED := time.Now().Add(-1)

	err := sessions.DeleteSession(sessionId)

	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    sessionId,
		Expires:  EXPIRED,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		Secure:   false, // пока http
	})

	return err
}
