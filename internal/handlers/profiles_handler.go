package handlers

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/analytics"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/payments"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/sessions"
	subscribelevels "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/subscribe_levels"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/subscriptions"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/users"
	auth "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/usecases/authorization"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/usecases/files"
	pay "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/usecases/payment"
	"github.com/robfig/cron/v3"
)

type BodyProfile struct {
	Profile model.Profile `json:"profile"`
}

type Profiles struct {
	Profiles []model.Profile `json:"profiles"`
}

type Analytics struct {
	Analytics model.Analitycs `json:"analytics"`
}

type ProfileHandler struct {
	SessionManager *sessions.RedisManager
	// Session       sessions.SessionRepository
	Users         users.UserRepository
	Levels        subscribelevels.SubscribeLevelRepository
	Subscriptions subscriptions.SubscriptionRepository
	Payments      *payments.PaymentManager
	Analytics	analytics.AnalyticsRepository
}

func CreateProfileHandlerViaRepos(
	sessionManager *sessions.RedisManager,
	users users.UserRepository,
	levels subscribelevels.SubscribeLevelRepository,
	subscriptions subscriptions.SubscriptionRepository,
	payments *payments.PaymentManager,
	analytic analytics.AnalyticsRepository,
) *ProfileHandler {
	return &ProfileHandler{
		sessionManager,
		users,
		levels,
		subscriptions,
		payments,
		analytic,
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
	// if !auth.CheckAuthorizationManager(ctx, p.SessionManager) {
	// 	return Result{}, ErrUnathorized
	// }
	vars := auth.GetVars(ctx)
	if vars == nil {
		return Result{}, ErrNoVars
	}
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return Result{}, ErrBadID
	}

	cookie := auth.GetSession(ctx)
	userID := 0
	session, ok := p.SessionManager.CheckSessionCtxWrapper(ctx, cookie.Value)
	if ok {
		//return Result{}, ErrNoSession
		userID = int(session.UserID)
	}

	user, err := p.Users.GetUserWithSubscribers(id, userID)
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
			IsFollowed:      user.IsFollowed,
			VisiterSubscriptionLevelId: user.VisiterSubscriptionLevelId,
			SubscriptionId: user.SubscriptionId,
		}
		donated, err := p.Payments.GetPaymentsByAuthorId(user.ID)
		if err != nil {
			return Result{}, err
		}
		var sumInteger int
		var sumFractional int
		for _, payment := range donated {
			sumInteger += int(payment.PaymentInteger)
			sumFractional += int(payment.PaymentFractional)
		}
		sumInteger += sumFractional / 100
		sumFractional = sumFractional % 100
		profile.Donated = strconv.Itoa(int(sumInteger)) + "," + strconv.Itoa(int(sumFractional))
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
	if !auth.CheckAuthorizationManager(ctx, p.SessionManager) {
		return Result{}, ErrUnathorized
	}

	cookie := auth.GetSession(ctx)
	session, ok := p.SessionManager.CheckSessionCtxWrapper(ctx, cookie.Value)
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
	user := model.User{
		ID:          uint(session.UserID),
		Nickname:    GetFirst[string](form.Form.Value["nickname"]),
		Login:       GetFirst[string](form.Form.Value["login"]),
		Status:      GetFirst[string](form.Form.Value["status"]),
		Background:  GetFirst[string](form.Form.Value["background"]),
		Description: GetFirst[string](form.Form.Value["description"]),
	}
	currentUser, err := p.Users.GetUser(int(user.ID))
	if err != nil {
		return Result{}, ErrDataBase
	}
	fileArray, ok := form.Form.File["avatar"]
	if ok && len(fileArray) > 0 {
		path, err := files.SaveFile(fileArray[0], GetFirst[string](form.Form.Value["id"]))
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
			return Result{}, ErrMissmatchPassword //change error
		}
		user.Password = GetFirst[string](form.Form.Value["new_password"])
	}
	err = p.Users.ChangeUser(user)
	if err != nil {
		return Result{}, ErrDataBase
	}
	return Result{}, nil
}

func (p *ProfileHandler) ChangeUserStatusStratagy(ctx context.Context, form StatusForm) (Result, error) {
	if !auth.CheckAuthorizationManager(ctx, p.SessionManager) {
		return Result{}, ErrUnathorized
	}

	cookie := auth.GetSession(ctx)
	session, ok := p.SessionManager.CheckSessionCtxWrapper(ctx, cookie.Value)
	if !ok {
		return Result{}, ErrNoSession
	}

	err := p.Users.ChangeUserStatus(form.Body.Status, int(session.UserID))
	if err != nil {
		return Result{}, ErrDataBase
	}
	return Result{}, nil
}

func (p *ProfileHandler) ChangeUserDescriptionStratagy(ctx context.Context, form DescriptionForm) (Result, error) {
	if !auth.CheckAuthorizationManager(ctx, p.SessionManager) {
		return Result{}, ErrUnathorized
	}

	cookie := auth.GetSession(ctx)
	session, ok := p.SessionManager.CheckSessionCtxWrapper(ctx, cookie.Value)
	if !ok {
		return Result{}, ErrNoSession
	}

	err := p.Users.ChangeUserDescription(form.Body.Description, int(session.UserID))
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
	if !auth.CheckAuthorizationManager(ctx, p.SessionManager) {
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
	session, ok := p.SessionManager.CheckSessionCtxWrapper(ctx, cookie.Value)
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
	if !auth.CheckAuthorizationManager(ctx, p.SessionManager) {
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
	session, ok := p.SessionManager.CheckSessionCtxWrapper(ctx, cookie.Value)
	if !ok {
		return Result{}, ErrNoSession
	}

	levels, err := p.Levels.GetUserLevels(uint(id))
	if err != nil {
		return Result{}, ErrDataBase
	}
	subLevel := model.SubscribeLevel{}
	subLevel.Level = 9999
	for _, level := range levels {
		if level.ID == uint(form.Body.SubscriptionLevelId) {
			subLevel = level
			break
		}
	}

	switch subLevel.Level {
	case 9999:
		return Result{}, ErrNoLevelId

	case 0:
		subscription := model.Subscription{
			Id: uint(form.Body.SubscriptionId),
			Subscriber_id:         uint(session.UserID),
			Creator_id:            uint(id),
			Subscription_level_id: uint(form.Body.SubscriptionLevelId),
		}
		subId, err := p.Subscriptions.AddNewSubscription(subscription)
		if err != nil {
			return Result{}, ErrDataBase
		}
		return Result{Body: map[string]interface{}{"id": subId}}, nil

	default:
		value := strconv.Itoa(int(subLevel.CostInteger))
		if subLevel.CostFractional < 10 {
			value += ".0" + strconv.Itoa(int(subLevel.CostFractional))
		} else {
			value += "." +strconv.Itoa(int(subLevel.CostFractional))
		}

		payment := model.Payment{
			DonaterId: uint(session.UserID),
			CreatorId: uint(id),
			Currency: subLevel.Currency,
			Value: value,
			Type: model.PaymentTypeSubscription,
		}
		responseUkassa, err := pay.Subscribe(payment)
		if err != nil {
			//think
			log.Println(err)
			return Result{}, ErrPayment
		}

		switch responseUkassa.Status {
		case "pending":
			payment.Status = model.PaymentWaitingStatus
		case "succeeded":
			payment.Status = model.PaymentSucceededStatus
		case "canceled":
			payment.Status = model.PaymentCanceledStatus
		}

		payment.UUID = responseUkassa.Id
		payment.PaymentInteger = subLevel.CostInteger
		payment.PaymentFractional = subLevel.CostFractional
		id, err := p.Payments.CreateNewPayment(payment)
		if err != nil {
			//think
			log.Println(err)
			return Result{}, ErrDataBase
		}
		payment.Id = id

		c := cron.New()
		_, err = c.AddFunc("* * * * *", func () {
			fmt.Println("joba2")
			payment, err := pay.CheckPaymentStatusAPI(p.Payments, payment)
			if err != nil {
				log.Println(err)
				return
			}
			if payment.Status == model.PaymentSucceededStatus {
					subscription := model.Subscription{
						Id: uint(form.Body.SubscriptionId),
						Subscriber_id:         payment.DonaterId,
						Creator_id:            payment.CreatorId,
						Subscription_level_id: subLevel.ID,
					}
					_, err := p.Subscriptions.AddNewSubscription(subscription)
					if err != nil {
						log.Println(err)
						return
					}
				c.Stop()
			}
			if payment.Status == model.PaymentCanceledStatus {
				c.Stop()
			}
		})
		c.Start()

		if err != nil {
			//think
			log.Println(err)
			return Result{}, ErrDataBase
		}
		return Result{Body: BodyPayments{RedirectURL: responseUkassa.Confirmation.ConfirmationURL}}, nil
	}
}

func (p *ProfileHandler) UnfollowStratagy(ctx context.Context, form FollowForm) (Result, error) {
	if !auth.CheckAuthorizationManager(ctx, p.SessionManager) {
		return Result{}, ErrUnathorized
	}

	cookie := auth.GetSession(ctx)
	session, ok := p.SessionManager.CheckSessionCtxWrapper(ctx, cookie.Value)
	if !ok {
		return Result{}, ErrNoSession
	}

	err := p.Subscriptions.DeleteSubscription(form.Body.SubscriptionLevelId, int(session.UserID))
	if err != nil {
		return Result{}, ErrDataBase
	}
	return Result{}, nil
}

func (p *ProfileHandler) GetTopUsersStratagy(ctx context.Context, form EmptyForm) (Result, error) {
	vars := auth.GetVars(ctx)
	if vars == nil {
		return Result{}, ErrNoVars
	}
	limit, err := strconv.Atoi(vars["limit"])
	if err != nil {
		return Result{}, ErrBadID
	}
	users, err := p.Users.GetTopUsers(limit)

	if err != nil {
		return Result{}, ErrDataBase
	}

	var profiles []model.Profile
	for _, user := range users {
		profiles = append(profiles, model.Profile{
			User:        user,
			Subscribers: user.Subscribers,
		})
	}

	return Result{Body: Profiles{Profiles: profiles}}, nil
}

func (p *ProfileHandler) Search(ctx context.Context, form EmptyForm) (Result, error) {
	vars := auth.GetVars(ctx)
	if vars == nil {
		return Result{}, ErrNoVars
	}
	nickname := vars["nickname"]
	users, err := p.Users.Search(nickname)

	if err != nil {
		return Result{}, ErrDataBase
	}

	var profiles []model.Profile
	for _, user := range users {
		profiles = append(profiles, model.Profile{
			User:        user,
			Subscribers: user.Subscribers,
		})
	}

	return Result{Body: Profiles{Profiles: profiles}}, nil
}

func  (p *ProfileHandler) Analitycs(ctx context.Context, form EmptyForm) (Result, error) {
	if !auth.CheckAuthorizationManager(ctx, p.SessionManager) {
		return Result{}, ErrUnathorized
	}
	
	cookie := auth.GetSession(ctx)
	session, ok := p.SessionManager.CheckSessionCtxWrapper(ctx, cookie.Value)
	if !ok {
		return Result{}, ErrNoSession
	}

	analytic, err := p.Analytics.GetLastAnalytics(int(session.UserID))
	if err != nil {
		fmt.Println("HERE", err)
		return Result{}, ErrDataBase
	}

	return Result{Body: Analytics{Analytics: analytic}}, nil
}