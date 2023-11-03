package authorization

import (
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"
	sessrep "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/sessions"
	usrep "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/users"
)

const (
	SIDLen      = 32
	TTLDuration = 10 * time.Hour
)

func Authorize(users usrep.UserRepository, form *LoginForm) (*model.User, error) {
	user, ok := users.CheckUser(form.Login)

	if !ok {
		return nil, ErrWrongLogin
	}

	if user.Password != form.Password {
		return nil, ErrWrongPassword
	}

	return user, nil
}

// стоит заменить на мидлвару в будущем думаю
func CheckAuthorization(r *http.Request, sessions sessrep.SessionRepository) bool {
	session, err := r.Cookie("session_id")

	if err == nil && session != nil {
		_, authorized := sessions.CheckSession(session.Value)
		return authorized
	}

	return false
}

// стоит заменить на мидлвару в будущем думаю
func CheckAuthorizationByContext(ctx context.Context, sessions sessrep.SessionRepository) bool {
	session := GetSession(ctx)

	if session != nil {
		_, authorized := sessions.CheckSession(session.Value)
		return authorized
	}

	return false
}

func SetSessionContext(ctx context.Context, sessions sessrep.SessionRepository, userID uint32) {
	w := GetWriter(ctx)

	SID := uuid.New().String()
	TTL := time.Now().Add(TTLDuration)

	sessions.RegisterNewSession(sessrep.Session{
		SessionID: SID,
		UserID:    userID,
		TTL:       TTL,
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

func SetSession(w http.ResponseWriter, sessions sessrep.SessionRepository, userID uint32) {
	SID := uuid.New().String()
	TTL := time.Now().Add(TTLDuration)

	sessions.RegisterNewSession(sessrep.Session{
		SessionID: SID,
		UserID:    userID,
		TTL:       TTL,
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

func RemoveSessionContext(ctx context.Context, sessions sessrep.SessionRepository, sessionID string) error {
	w := GetWriter(ctx)

	EXPIRED := time.Now().Add(-1)

	err := sessions.DeleteSession(sessionID)
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		Expires:  EXPIRED,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		Secure:   false, // пока http
	})

	return err
}

func RemoveSession(w http.ResponseWriter, sessions sessrep.SessionRepository, sessionID string) error {
	EXPIRED := time.Now().Add(-1)

	err := sessions.DeleteSession(sessionID)
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    sessionID,
		Expires:  EXPIRED,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		Secure:   false, // пока http
	})

	return err
}
