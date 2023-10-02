package handlers

import (
	"encoding/json"
	"fmt"
	"io"
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
		Login:    regForm.Login,
		Password: regForm.Password,
		UserType: regForm.UserType,
	}

	err = api.Users.RegisterNewUser(&user)

	if err != nil {
		http.Error(w, `{"error":"user_registration"}`, 400)
		return
	}

	profile := model.Profile{
		User: user,
	}

	abc, _ := io.ReadAll(r.Body)
	fmt.Println((string)(abc))

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
	bodyForm, err := auth.ParseJSON(r.Body)

	if err != nil {
		http.Error(w, `{"error":"wrong_json"}`, 405)
		return
	}

	// users, _ := api.Profiles.GetProfiles()
	// fmt.Println(users)
	userForm := auth.LoginForm{
		Body_: *bodyForm,
	}
	user, err := auth.Authorize(api.Users, &userForm)

	if err != nil {
		fmt.Println(err)
		http.Error(w, `{"error":"user_registration"}`, 400)
		return
	}

	auth.SetSession(w, api.Sessions, uint32(user.ID))

	w.WriteHeader(200)
}

func (api *RepoHandler) IsAuthorized(w http.ResponseWriter, r *http.Request) {
	AddAllowHeaders(w)
	w.Header().Add("Content-Type", "application/json")
	if auth.CheckAuthorization(r, api.Sessions) {
		json.NewEncoder(w).Encode(&Result{Body: map[string]interface{}{"is_authorized": true}})
	} else {
		json.NewEncoder(w).Encode(&Result{Body: map[string]interface{}{"is_authorized": false}})
	}
}
