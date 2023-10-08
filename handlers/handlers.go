package handlers

import (
	"encoding/json"
	"net/http"
	auth "project/authorization"
	"project/model"
	reg "project/registration"
)

type RepoHandler struct {
	sessions auth.SessionRepository
	users    reg.UserRepository
	profiles reg.ProfileRepository
}

func CreateRepoHandler(sessions auth.SessionRepository, users reg.UserRepository, profiles reg.ProfileRepository) *RepoHandler {
	return &RepoHandler{
		sessions,
		users,
		profiles,
	}
}

func (api *RepoHandler) Signup(w http.ResponseWriter, r *http.Request) {
	AddAllowHeaders(w)
	defer r.Body.Close()
	regForm, err := reg.ParseJSON(r.Body)

	if err != nil {
		http.Error(w, `{"error":"wrong_json"}`, http.StatusBadRequest)
		return
	}

	user, err := regForm.Validate()

	if err != nil {
		http.Error(w, `{"error":"user_validation"}`, http.StatusBadRequest)
	}

	err = api.users.RegisterNewUser(user)
	
	if err != nil {
		http.Error(w, `{"error":"user_registration"}`, 400)
		return
	}

	profile := model.Profile{
		User: *user,
	}

	err = api.profiles.RegisterNewProfile(&profile)
	if err != nil {
		http.Error(w, `{"error":"profile_registration"}`, 401)
		return
	}

	auth.SetSession(w, api.sessions, uint32(user.ID))

	body := map[string]interface{}{
		"id": user.ID,
	}

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&Result{Body: body})
}

func (api *RepoHandler) Login(w http.ResponseWriter, r *http.Request) {
	AddAllowHeaders(w)
	defer r.Body.Close()
	bodyForm, err := auth.ParseJSON(r.Body)

	if err != nil {
		http.Error(w, `{"error":"wrong_json"}`, http.StatusBadRequest)
		return
	}
	userForm := auth.LoginForm{
		Body_: *bodyForm,
	}
	user, err := auth.Authorize(api.users, &userForm)

	if err != nil {
		http.Error(w, `{"error":"user_registration"}`, 400)
		return
	}

	auth.SetSession(w, api.sessions, uint32(user.ID))

	w.WriteHeader(200)

	body := map[string]interface{}{
		"id": user.ID,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&Result{Body: body})
}

func (api *RepoHandler) Logout(w http.ResponseWriter, r *http.Request) {
	AddAllowHeaders(w)
	session, err := r.Cookie("session_id")

	if err != nil {
		http.Error(w, `{"error":"user_logout"}`, 400)
	}

	auth.RemoveSession(w, api.sessions, session.Value)

	w.WriteHeader(200)
}

func (api *RepoHandler) IsAuthorized(w http.ResponseWriter, r *http.Request) {
	AddAllowHeaders(w)
	w.Header().Add("Content-Type", "application/json")
	if auth.CheckAuthorization(r, api.sessions) {
		json.NewEncoder(w).Encode(&Result{Body: map[string]interface{}{"is_authorized": true}})
	} else {
		json.NewEncoder(w).Encode(&Result{Body: map[string]interface{}{"is_authorized": false}})
	}
}
