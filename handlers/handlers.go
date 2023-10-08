package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	auth "project/authorization"
	"project/model"
	reg "project/registration"
)

const (
	MAX_BYTES_TO_READ = 1024 * 2
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

	body := http.MaxBytesReader(w, r.Body, MAX_BYTES_TO_READ)
	defer body.Close()

	decoder := json.NewDecoder(body)
	regForm := &reg.BodySignUp{}

	err := decoder.Decode(regForm)
	if err != nil {
		http.Error(w, `{"error":"wrong_json"}`, http.StatusBadRequest)
		return
	}

	user, err := regForm.Validate()
	if err != nil {
		http.Error(w, `{"error":"user_validation"}`, http.StatusBadRequest)
	}

	errReg := api.users.RegisterNewUser(user)
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error":"%v"}`, errReg.Err), errReg.StatusCode)
		return
	}

	profile := model.Profile{
		User: *user,
	}

	errReg = api.profiles.RegisterNewProfile(&profile)
	if errReg != nil {
		http.Error(w, fmt.Sprintf(`{"error":"%v"}`, errReg.Err), errReg.StatusCode)
		return
	}

	auth.SetSession(w, api.sessions, uint32(user.ID))

	bodyResponse := map[string]interface{}{
		"id": user.ID,
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&Result{Body: bodyResponse})
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
		http.Error(w, `{"error":"user_registration"}`, http.StatusBadRequest)
		return
	}

	auth.SetSession(w, api.sessions, uint32(user.ID))

	w.WriteHeader(http.StatusOK)

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
		http.Error(w, `{"error":"user_logout"}`, http.StatusBadRequest)
	}

	auth.RemoveSession(w, api.sessions, session.Value)

	w.WriteHeader(http.StatusOK)
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
