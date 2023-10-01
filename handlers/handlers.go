package handlers

import (
	"net/http"
	auth "project/authorization"
	"project/model"
	reg "project/registration"
)

type RepoHandler struct {
	Sessions auth.SessionRepository
	Users    reg.UserRepository
	Profiles reg.ProfileRepository
}

func CreateRepoHandler() *RepoHandler {
	return &RepoHandler{
		auth.CreateSessionStorage(),
		reg.CreateUserStorage(),
		reg.CreateProfileStorage(),
	}
}

func (api *RepoHandler) Signup(w http.ResponseWriter, r *http.Request) {
	AddAllowHeaders(w)
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

	err = api.Users.RegisterNewUser(&user)

	if err != nil {
		http.Error(w, `{"error":"user_registration"}`, 400)
		return
	}

	profile := model.Profile{
		User: user,
	}

	err = api.Profiles.RegisterNewProfile(&profile)

	if err != nil {
		http.Error(w, `{"error":"profile_registration"}`, 401)
		return
	}

	auth.SetSession(w, api.Sessions, uint32(user.ID))

	w.WriteHeader(200)
}

func (api *RepoHandler) Login(w http.ResponseWriter, r *http.Request) {
	AddAllowHeaders(w)
	defer r.Body.Close()
	userForm, err := auth.ParseJSON(r.Body)

	if err != nil {
		http.Error(w, `{"error":"wrong_json"}`, 405)
		return
	}

	// users, _ := api.Profiles.GetProfiles()
	// fmt.Println(users)
	user, err := auth.Authorize(api.Users, userForm)

	if err != nil {
		http.Error(w, `{"error":"user_registration"}`, 400)
		return
	}

	auth.SetSession(w, api.Sessions, uint32(user.ID))

	w.WriteHeader(200)
}

func (api *RepoHandler) Root(w http.ResponseWriter, r *http.Request) {
	AddAllowHeaders(w)
	if auth.CheckAuthorization(r, api.Sessions) {
		w.Write([]byte("autrorized"))
	} else {
		w.Write([]byte("not autrorized"))
	}
}
