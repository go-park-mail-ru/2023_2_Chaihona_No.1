package authorization

import (
	"net/http"
	"time"
	reg "project/registration"
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

	user, ok := api.Users.CheckUser(r.FormValue("login"))

	if !ok {
		http.Error(w, `no user`, 404)
		return
	}

	if user.Password != r.FormValue("password") {
		http.Error(w, `bad pass`, 400)
		return
	}

	SID := RandStringRunes(32)
	TTL := time.Now().Add(10 * time.Hour)

	session := &Session{
		SessionId: SID,
		UserId:    uint32(user.ID),
		Ttl:       TTL,
	}

	api.Sessions.RegisterNewSession(*session)

	cookie := &http.Cookie{
		Name:    "session_id",
		Value:   SID,
		Expires: time.Now().Add(10 * time.Hour),
	}

	http.SetCookie(w, cookie)
	w.Write([]byte(SID))
}

func (api *RepoHandler) Root(w http.ResponseWriter, r *http.Request) {
	authorized := false
	session, err := r.Cookie("session_id")

	if err == nil && session != nil {
		_, authorized = api.Sessions.CheckSession(session.Value)
	}

	if authorized {
		w.Write([]byte("autrorized"))
	} else {
		w.Write([]byte("not autrorized"))
	}
}
