package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"
	profsrep "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/profiles"
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
	profiles profsrep.ProfileRepository
}

func CreateRepoHandler(
	sessions sessrep.SessionRepository,
	users usrep.UserRepository,
	profiles profsrep.ProfileRepository,
) *RepoHandler {
	return &RepoHandler{
		sessions,
		users,
		profiles,
	}
}

// swagger:route OPTIONS /api/v1/registration Auth SignUpOptions
// Handle OPTIONS request
// Responses:
//
//	200: result

// swagger:route POST /api/v1/registration Auth SignUp
// SignUp user
//
// Responses:
//
//	200: result
//	400: result
//	500: result
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

	errReg := api.users.RegisterNewUser(user)
	if err != nil {
		switch errReg.(type) {
		case usrep.ErrorUserRegistration:
			errReg := errReg.(usrep.ErrorUserRegistration)
			http.Error(w, fmt.Sprintf(`{"error":"%v"}`, errReg.Err), errReg.StatusCode)
			return
		}
	}

	profile := model.Profile{
		User: *user,
	}

	errReg = api.profiles.RegisterNewProfile(&profile)
	if errReg != nil {
		switch errReg.(type) {
		case profsrep.ErrorProfileRegistration:
			errReg := errReg.(profsrep.ErrorProfileRegistration)
			http.Error(w, fmt.Sprintf(`{"error":"%v"}`, errReg.Err), errReg.StatusCode)
			return
		}
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

// swagger:route OPTIONS /api/v1/login Auth LoginOptions
// Handle OPTIONS request
// Responses:
//
//	200: result

// swagger:route POST /api/v1/login Auth Login
// Login user
//
// Responses:
//
//	200: result
//	400: result
//	500: result
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

// swagger:route OPTIONS /api/v1/logout Auth LogoutOptions
// Handle OPTIONS request
// Responses:
//
//	200: result

// swagger:route POST /api/v1/logout Auth Logout
// Logout user
//
// Responses:
//
//	200: result
//	400: result
//	500: result
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

// swagger:route OPTIONS /api/v1/is_authorized Auth IsAuthorizedOptions
// Handle OPTIONS request
// Responses:
//
//	200: result

// swagger:route GET /api/v1/is_authorized Auth IsAuthorized
// Check Authorization
//
// Responses:
//
//	200: result
//	400: result
//	500: result
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