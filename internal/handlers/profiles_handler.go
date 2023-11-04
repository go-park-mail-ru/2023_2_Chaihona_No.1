package handlers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/sessions"
	subscribelevels "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/subscribe_levels"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/subscriptions"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/users"
	auth "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/usecases/authorization"
)

type BodyProfile struct {
	Profile model.Profile `json:"profile"`
}

type ProfileHandler struct {
	Session       sessions.SessionRepository
	Users         users.UserRepository
	Levels        subscribelevels.SubscribeLevelRepository
	Subscriptions subscriptions.SubscriptionRepository
}

func CreateProfileHandlerViaRepos(
	session sessions.SessionRepository,
	users users.UserRepository,
	levels subscribelevels.SubscribeLevelRepository,
	subscriptions subscriptions.SubscriptionRepository,
) *ProfileHandler {
	return &ProfileHandler{
		session,
		users,
		levels,
		subscriptions,
	}
}

// swagger:route OPTIONS /api/v1/profile/{id} Profile GetInfoOptions
// Handle OPTIONS request
// Responses:
//
//	200: result

// swagger:route GET /api/v1/profile/{id} Profile GetInfo
// Get profile info
//
// Parameters:
//   - name: id
//     in: path
//     description: ID of user
//     required: true
//     type: integer
//     format: int
//
// Responses:
//
//	200: result
//	400: result
//	401: result
//	500: result
func (p *ProfileHandler) GetInfoStrategy(ctx context.Context, form EmptyForm) (Result, error) {
	if !auth.CheckAuthorizationByContext(ctx, p.Session) {
		return Result{}, ErrUnathorized
	}

	vars := auth.GetVars(ctx)
	if vars == nil {
		return Result{}, ErrNoVars
	}
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return Result{}, ErrBadID
	}

	user, err := p.Users.GetUserWithSubscribers(id)
	if err != nil {
		return Result{}, err
	}
	var profile model.Profile
	if user.Is_author {
		levels, err := p.Levels.GetUserLevels(uint(id))
		if err != nil {
			return Result{}, err
		}
		user.UserType = model.CreatorStatus
		profile = model.Profile{
			User:            user,
			Subscribers:     user.Subscribers,
			SubscribeLevels: levels,
		}
	} else {
		subscriptions, err := p.Subscriptions.GetUserSubscriptions(id)
		if err != nil {
			return Result{}, err
		}

		// donated, err := p.Payments.SumUserPayments(id)
		// if err != nil {
		// 	http.Error(w, `{"error":"db"}`, 500)
		// 	return
		// }

		user.UserType = model.SimpleUserStatus
		profile = model.Profile{
			User:          user,
			Subscriptions: subscriptions,
		}
	}
	return Result{Body: BodyProfile{Profile: profile}}, nil

}

func (p *ProfileHandler) ChangeUserStratagy(ctx context.Context, form UserForm) (Result, error) {
	if !auth.CheckAuthorizationByContext(ctx, p.Session) {
		return Result{}, ErrUnathorized
	}

	cookie := auth.GetSession(ctx)
	session, ok := p.Session.CheckSession(cookie.Value)
	if !ok {
		return Result{}, ErrNoSession
	}
	if session.UserID != uint32(form.Body.ID) {
		return Result{}, fmt.Errorf("%s", "wrong_change")
	}

	user := model.User{
		ID: form.Body.ID,
		Nickname: form.Body.Nickname,
		Login: form.Body.Login,
		UserType: form.Body.UserType,
		Status: form.Body.Status,
		Avatar: form.Body.Avatar,
		Background: form.Body.Background,
		Description: form.Body.Description,
	}
	err := p.Users.ChangeUser(user)
	if err != nil {
		return Result{}, ErrDataBase
	}
	return Result{}, nil
}

func (p *ProfileHandler) DeleteUserStratagy(ctx context.Context, form EmptyForm) (Result, error) {
	if !auth.CheckAuthorizationByContext(ctx, p.Session) {
		return Result{}, ErrUnathorized
	}

	vars := auth.GetVars(ctx)
	if vars == nil {
		return Result{}, ErrNoVars
	}
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return Result{}, ErrBadID
	}

	cookie := auth.GetSession(ctx)
	session, ok := p.Session.CheckSession(cookie.Value)
	if !ok {
		return Result{}, ErrNoSession
	}

	if session.UserID != uint32(id) {
		return Result{}, fmt.Errorf("%s", "wrong_delete")
	}

	err = p.Users.DeleteUser(id)
	if err != nil {
		return Result{}, ErrDataBase
	}
	return Result{}, nil
}