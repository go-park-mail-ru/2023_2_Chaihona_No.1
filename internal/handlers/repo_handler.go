package handlers

import (
	"context"

	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"
	sessrep "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/sessions"
	usrep "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/users"
	auth "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/usecases/authorization"
	reg "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/usecases/registration"
)

type RepoHandler struct {
	sessions sessrep.SessionRepository
	users    usrep.UserRepository
}

type EmptyForm struct{}

func (f EmptyForm) IsValide() bool {
	return true
}

func (f EmptyForm) IsEmpty() bool {
	return true
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
func (api *RepoHandler) SignupStrategy(ctx context.Context, form reg.SignupForm) (*Result, error) {
	user := &model.User{
		Login:    form.Body.Login,
		Password: form.Body.Password,
		UserType: form.Body.UserType,
	}

	id, errReg := api.users.RegisterNewUser(user)
	if errReg != nil {
		return nil, errReg
	}

	auth.SetSessionContext(ctx, api.sessions, uint32(user.ID))

	bodyResponse := map[string]interface{}{
		"id": id,
	}

	return &Result{Body: bodyResponse}, nil
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
func (api *RepoHandler) LoginStrategy(ctx context.Context, form auth.LoginForm) (*Result, error) {
	user, err := auth.Authorize(api.users, &form)
	if err != nil {
		return nil, err
	}

	auth.SetSessionContext(ctx, api.sessions, uint32(user.ID))

	bodyResponse := map[string]interface{}{
		"id": user.ID,
	}

	return &Result{Body: bodyResponse}, nil
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
func (api *RepoHandler) LogoutStrategy(ctx context.Context, form EmptyForm) (Result, error) {
	session := auth.GetSession(ctx)
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
func (api *RepoHandler) IsAuthorizedStrategy(ctx context.Context, form EmptyForm) (Result, error) {
	if auth.CheckAuthorizationByContext(ctx, api.sessions) {
		return Result{Body: map[string]interface{}{"is_authorized": true}}, nil
	}

	return Result{Body: map[string]interface{}{"is_authorized": false}}, nil
}
