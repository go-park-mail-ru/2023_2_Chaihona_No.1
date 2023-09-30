package authorization

import (
	"net/http"
	"project/model"
	reg "project/registration"
	"time"
)

type RepoHandler struct {
	Sessions SessionRepository
	Users    reg.UserRepository
	Profiles reg.ProfileRepository
}

func CreateRepoHandler() *RepoHandler {
	return &RepoHandler{
		CreateSessionStorage(),
		reg.CreateUserStorage(),
		reg.CreateProfileStorage(),
	}
}

// http://127.0.0.1:8080/login?login=12&password=12

func (api *RepoHandler) Signup(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	regForm, err := reg.ParseJSON(r.Body)

	if err != nil {
		http.Error(w, `{"error":"wrong_json"}`, 405)
		return
	}

	user := model.User{
		Login:    regForm.Body_.Login,
		Password: regForm.Body_.Password,
		UserType: regForm.Body_.UserType,
	}

	err = api.Users.RegisterNewUser(user)

	if err != nil {
		http.Error(w, `{"error":"user_registration"}`, 400)
		return
	}

	if regForm.Body_.UserType == "creator" {
		profile := model.Profile{
			ID:   user.ID,
			User: user,
		}

		err = api.Profiles.RegisterNewProfile(profile)

		if err != nil {
			http.Error(w, `{"error":"profile_registration"}`, 401)
			return
		}
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

func (api *RepoHandler) Login(w http.ResponseWriter, r *http.Request) {
	userForm, err := ParseJSON(r)

	if err != nil {
		http.Error(w, `{"error":"wrong_json"}`, 405)
		return
	}

	user, err := Authorize(api.Users, userForm)

	if err != nil {
		http.Error(w, `{"error":"user_registration"}`, 400)
		return
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

func (api *RepoHandler) Root(w http.ResponseWriter, r *http.Request) {
	if CheckAuthorization(r, api.Sessions) {
		w.Write([]byte("autrorized"))
	} else {
		w.Write([]byte("not autrorized"))
	}
}
