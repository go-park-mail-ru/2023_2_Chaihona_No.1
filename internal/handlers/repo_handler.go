package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"
	profsrep "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/profiles"
	sessrep "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/sessions"
	usrep "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/users"
	auth "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/usecases/authorization"
	reg "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/usecases/registration"
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

func (api *RepoHandler) SignupStrategy(ctx context.Context, form reg.SignupForm) (*Result, error) {
	user := &model.User{
		Login:    form.Login,
		Password: form.Password,
		UserType: form.UserType,
	}

	errReg := api.users.RegisterNewUser(user)
	if errReg != nil {
		return nil, errReg
	}

	profile := model.Profile{
		User: *user,
	}

	errReg = api.profiles.RegisterNewProfile(&profile)
	if errReg != nil {
		return nil, errReg
	}

	auth.SetSessionContext(ctx, api.sessions, uint32(user.ID))

	bodyResponse := map[string]interface{}{
		"id": user.ID,
	}

	return &Result{Body: bodyResponse}, nil
}

func (api *RepoHandler) LoginStrategy(ctx context.Context, form auth.LoginForm) (*Result, error) {
	loginForm := &auth.LoginForm{}

	user, err := auth.Authorize(api.users, loginForm)
	if err != nil {
		return nil, err
	}

	auth.SetSessionContext(ctx, api.sessions, uint32(user.ID))

	bodyResponse := map[string]interface{}{
		"id": user.ID,
	}

	return &Result{Body: bodyResponse}, nil
}

type EmptyForm struct{}

func (f EmptyForm) IsValide() bool {
	return true
}

func (api *RepoHandler) LogoutStrategy(ctx context.Context, form EmptyForm) (Result, error) {
	session := ctx.Value(sessionIDKey{}).(*http.Cookie)
	if session == nil {
		return Result{}, ErrLogoutCookie
	}

	err := auth.RemoveSessionContext(ctx, api.sessions, session.Value)
	if err == nil {
		return Result{}, nil
	} else {
		return Result{}, ErrLogoutDeleteSession
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

func (api *RepoHandler) IsAuthorizedStrategy(ctx context.Context, form EmptyForm) (Result, error) {
	if auth.CheckAuthorizationByCookie(ctx.Value(sessionIDKey{}).(*http.Cookie), api.sessions) {
		return Result{Body: map[string]interface{}{"is_authorized": true}}, nil
	}

	return Result{Body: map[string]interface{}{"is_authorized": false}}, ErrUnathorized
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
