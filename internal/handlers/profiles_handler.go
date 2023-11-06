package handlers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/payments"
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
	Payments payments.PaymentRepository
}

func CreateProfileHandlerViaRepos(
	session sessions.SessionRepository,
	users users.UserRepository,
	levels subscribelevels.SubscribeLevelRepository,
	subscriptions subscriptions.SubscriptionRepository,
	payments payments.PaymentRepository,
) *ProfileHandler {
	return &ProfileHandler{
		session,
		users,
		levels,
		subscriptions,
		payments,
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
			IsFollowed: true,
		}
		donated, err := p.Payments.GetPaymentsByAuthorId(user.ID)
		if err != nil {
			return Result{}, err
		}
		profile.Donated = strconv.Itoa(int(donated.PaymentInteger)) + "," + strconv.Itoa(int(donated.PaymentFractional))
		profile.Currency = "RUB"
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

func (p *ProfileHandler) ChangeUserStratagy(ctx context.Context, form FileForm) (Result, error) {
	if !auth.CheckAuthorizationByContext(ctx, p.Session) {
		return Result{}, ErrUnathorized
	}

	cookie := auth.GetSession(ctx)
	session, ok := p.Session.CheckSession(cookie.Value)
	if !ok {
		return Result{}, ErrNoSession
	}
	// formUserId, err := strconv.Atoi(GetFirst[string](form.Form.Value["id"]))
	// if err != nil {
	// 	return Result{}, ErrBadID
	// }
	// if session.UserID != uint32(formUserId) {
	// 	return Result{}, fmt.Errorf("%s", "wrong_change")
	// }

	fmt.Println(session.UserID)
	user := model.User{
		ID: uint(session.UserID),
		Nickname: GetFirst[string](form.Form.Value["nickname"]),
		Login: GetFirst[string](form.Form.Value["login"]),
		Status: GetFirst[string](form.Form.Value["status"]),
		Background: GetFirst[string](form.Form.Value["background"]),
		Description: GetFirst[string](form.Form.Value["description"]),
	}
	currentUser, err := p.Users.GetUser(int(user.ID))
	if err != nil {
		return Result{}, ErrDataBase
	}
	fileArray, ok := form.Form.File["avatar"]
	if ok&&len(fileArray)>0 {
		path, err := saveFile(fileArray[0], GetFirst[string](form.Form.Value["id"]))
		if err != nil {
			return Result{}, err
		}
		user.Avatar = path
	} else {
		user.Avatar = currentUser.Avatar
	}
	if GetFirst[string](form.Form.Value["login"]) == "" {
		user.Login = currentUser.Login 
	} else {
		user.Login = GetFirst[string](form.Form.Value["login"])
	}
	user.Password = currentUser.Password
	if GetFirst[string](form.Form.Value["new_password"]) != "" {
		if GetFirst[string](form.Form.Value["old_password"]) != currentUser.Password {
			return Result{}, ErrValidation //change error
		}
		user.Password = GetFirst[string](form.Form.Value["new_password"])
	}
	err = p.Users.ChangeUser(user)
	if err != nil {
		return Result{}, ErrDataBase
	}
	return Result{}, nil
}

// func (p *ProfileHandler) ChangeUserStratagy(ctx context.Context, form UserForm) (Result, error) {
// 	if !auth.CheckAuthorizationByContext(ctx, p.Session) {
// 		return Result{}, ErrUnathorized
// 	}

// 	cookie := auth.GetSession(ctx)
// 	session, ok := p.Session.CheckSession(cookie.Value)
// 	if !ok {
// 		return Result{}, ErrNoSession
// 	}
// 	if session.UserID != uint32(form.Body.User.ID) {
// 		return Result{}, fmt.Errorf("%s", "wrong_change")
// 	}

// 	user := model.User{
// 		ID: form.Body.User.ID,
// 		Nickname: form.Body.User.Nickname,
// 		Login: form.Body.User.Login,
// 		Status: form.Body.User.Status,
// 		Avatar: form.Body.User.Avatar,
// 		Background: form.Body.User.Background,
// 		Description: form.Body.User.Description,
// 		Is_author: form.Body.User.IsAuthor,
// 	}
// 	if form.Body.User.NewPassword != "" {
// 		currentUser, err := p.Users.GetUser(int(user.ID))
// 		if err != nil {
// 			return Result{}, ErrDataBase
// 		}
// 		if form.Body.User.OldPassword != currentUser.Password {
// 			return Result{}, ErrValidation //change error
// 		}
// 		user.Password = form.Body.User.NewPassword
// 	}
// 	err := p.Users.ChangeUser(user)
// 	if err != nil {
// 		return Result{}, ErrDataBase
// 	}
// 	return Result{}, nil
// }

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

func (p *ProfileHandler) FollowStratagy(ctx context.Context, form FollowForm) (Result, error) {
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

	subscription := model.Subscription{
		Subscriber_id: uint(session.UserID),
		Creator_id: uint(id),
		Subscription_level_id: uint(form.Body.SubscriptionLevelId),
	}
	_, err = p.Subscriptions.AddNewSubscription(subscription)
	if err != nil {
		fmt.Println(err)
		return Result{}, ErrDataBase
	}
	return Result{}, nil
}

func (p *ProfileHandler) UnfollowStratagy(ctx context.Context, form FollowForm) (Result, error) {
	if !auth.CheckAuthorizationByContext(ctx, p.Session) {
		return Result{}, ErrUnathorized
	}


	cookie := auth.GetSession(ctx)
	session, ok := p.Session.CheckSession(cookie.Value)
	if !ok {
		return Result{}, ErrNoSession
	}

	err := p.Subscriptions.DeleteSubscription(form.Body.SubscriptionLevelId, int(session.UserID))
	if err != nil {
		return Result{}, ErrDataBase
	}
	return Result{}, nil
}