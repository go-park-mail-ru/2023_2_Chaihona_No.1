package authorization

import (
	"net/http"
	"project/model"
	reg "project/registration"
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

	SetSession(w, api.Sessions, uint32(user.ID))

	w.WriteHeader(200)
}

func (api *RepoHandler) Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	userForm, err := ParseJSON(r.Body)

	if err != nil {
		http.Error(w, `{"error":"wrong_json"}`, 405)
		return
	}

	user, err := Authorize(api.Users, userForm)

	if err != nil {
		http.Error(w, `{"error":"user_registration"}`, 400)
		return
	}

	SetSession(w, api.Sessions, uint32(user.ID))

	w.WriteHeader(200)
}

func (api *RepoHandler) Root(w http.ResponseWriter, r *http.Request) {
	if CheckAuthorization(r, api.Sessions) {
		w.Write([]byte("autrorized"))
	} else {
		w.Write([]byte("not autrorized"))
	}
}
