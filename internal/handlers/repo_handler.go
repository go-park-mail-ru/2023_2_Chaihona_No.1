package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	sessrep "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/sessions"
	usrep "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/users"
	auth "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/usecases/authorization"
	reg "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/usecases/registration"
)

const (
	maxBytesToRead = 1024 * 2
)

type RepoHandler struct {
	sessions sessrep.SessionRepository
	users    usrep.UserRepository
}

func CreateRepoHandler(
	sessions sessrep.SessionRepository,
	users usrep.UserRepository,
) *RepoHandler {
	return &RepoHandler{
		sessions,
		users,
	}
}

func (api *RepoHandler) Signup(w http.ResponseWriter, r *http.Request) {
	AddAllowHeaders(w)

	body := http.MaxBytesReader(w, r.Body, maxBytesToRead)

	decoder := json.NewDecoder(body)
	regForm := &reg.SignupForm{}

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

	id, errReg := api.users.RegisterNewUser(user)
	if err != nil {
		switch errReg.(type) {
		case usrep.ErrorUserRegistration:
			errReg := errReg.(usrep.ErrorUserRegistration)
			http.Error(w, fmt.Sprintf(`{"error":"%v"}`, errReg.Err), errReg.StatusCode)
			return
		}
	}

	auth.SetSession(w, api.sessions, uint32(user.ID))

	bodyResponse := map[string]interface{}{
		"id": id,
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

	body := http.MaxBytesReader(w, r.Body, maxBytesToRead)

	decoder := json.NewDecoder(body)
	loginForm := &auth.LoginForm{}

	err := decoder.Decode(loginForm)
	if err != nil {
		http.Error(w, `{"error":"wrong_json"}`, http.StatusBadRequest)
		return
	}

	user, err := auth.Authorize(api.users, loginForm)
	if err != nil {
		http.Error(w, `{"error":"wrong_input"}`, http.StatusBadRequest)
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
