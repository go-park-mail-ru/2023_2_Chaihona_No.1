package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	auth "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/authorization"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/model"
	reg "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/registration"
)

const (
	MAX_BYTES_TO_READ = 1024 * 2
)

type RepoHandler struct {
	sessions auth.SessionRepository
	users    reg.UserRepository
	profiles reg.ProfileRepository
}

func CreateRepoHandler(
	sessions auth.SessionRepository,
	users reg.UserRepository,
	profiles reg.ProfileRepository,
) *RepoHandler {
	return &RepoHandler{
		sessions,
		users,
		profiles,
	}
}

func (api *RepoHandler) Signup(w http.ResponseWriter, r *http.Request) {
	AddAllowHeaders(w)

	body := http.MaxBytesReader(w, r.Body, MAX_BYTES_TO_READ)

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
		return
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
	err = json.NewEncoder(w).Encode(&Result{Body: bodyResponse})
	if err != nil {
		log.Println(err)
	}
}

func (api *RepoHandler) Login(w http.ResponseWriter, r *http.Request) {
	AddAllowHeaders(w)

	body := http.MaxBytesReader(w, r.Body, MAX_BYTES_TO_READ)

	decoder := json.NewDecoder(body)
	loginForm := &auth.BodyLogin{}

	err := decoder.Decode(loginForm)
	if err != nil {
		http.Error(w, `{"error":"wrong_json"}`, http.StatusBadRequest)
		return
	}

	userForm := auth.LoginForm{
		Body: *loginForm,
	}
	user, err := auth.Authorize(api.users, &userForm)
	if err != nil {
		http.Error(w, `{"error":"wrong_input"}`, http.StatusBadRequest)
		return
	}

	auth.SetSession(w, api.sessions, uint32(user.ID))

	w.WriteHeader(http.StatusOK)

	bodyResponse := map[string]interface{}{
		"id": user.ID,
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&Result{Body: bodyResponse})
	if err != nil {
		log.Println(err)
	}
}

func (api *RepoHandler) Logout(w http.ResponseWriter, r *http.Request) {
	AddAllowHeaders(w)
	session, err := r.Cookie("session_id")
	if err != nil {
		http.Error(w, `{"error":"user_logout"}`, http.StatusBadRequest)
		return
	}

	err = auth.RemoveSession(w, api.sessions, session.Value)
	if err == nil {
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, `{"error":"user_logout"}`, http.StatusInternalServerError)
	}
}

func (api *RepoHandler) IsAuthorized(w http.ResponseWriter, r *http.Request) {
	AddAllowHeaders(w)
	w.Header().Add("Content-Type", "application/json")
	if auth.CheckAuthorization(r, api.sessions) {
		w.WriteHeader(http.StatusOK)
		err := json.NewEncoder(w).
			Encode(&Result{Body: map[string]interface{}{"is_authorized": true}})
		if err != nil {
			log.Println(err)
		}
		return
	}
	w.WriteHeader(http.StatusUnauthorized)
	err := json.NewEncoder(w).
		Encode(&Result{Body: map[string]interface{}{"is_authorized": false}})
	if err != nil {
		log.Println(err)
	}
}
