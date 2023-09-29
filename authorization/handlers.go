package authorization

import (
	"net/http"
	reg "project/registration"
	"time"
)

type RepoHandler struct {
	Sessions SessionRepository
	Users    reg.UserRepository
}

func CreateRepoHandler() *RepoHandler {
	return &RepoHandler{
		CreateSessionStorage(),
		reg.CreateUserStorage(),
	}
}

// http://127.0.0.1:8080/login?login=12&password=12

func (api *RepoHandler) Login(w http.ResponseWriter, r *http.Request) {
	userForm, err := ParseJSON(r)

	if err != nil {
		http.Error(w, `wrong json`, 405)
	} else {
		user, err := Authorize(api.Users, userForm)

		if err != nil {
			http.Error(w, `wrong pass or login`, 400)
		}

		SID := RandStringRunes(32)
		TTL := time.Now().Add(10 * time.Hour)

		api.Sessions.RegisterNewSession(Session{
			SessionId: SID,
			UserId:    uint32(user.ID),
			Ttl:       TTL,
		})

		http.SetCookie(w, &http.Cookie{
			Name:     "session_id",
			Value:    SID,
			Expires:  time.Now().Add(10 * time.Hour),
			HttpOnly: true,
		})

		w.WriteHeader(200)
	}
}

func (api *RepoHandler) Root(w http.ResponseWriter, r *http.Request) {
	if CheckAuthorization(r, api.Sessions) {
		w.Write([]byte("autrorized"))
	} else {
		w.Write([]byte("not autrorized"))
	}
}
