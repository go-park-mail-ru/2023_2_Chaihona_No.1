package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

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

func (p *ProfileHandler) GetInfo(w http.ResponseWriter, r *http.Request) {
	AddAllowHeaders(w)
	if !auth.CheckAuthorization(r, p.Session) {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, `{"error":"bad id"}`, 400)
		return
	}

	user, err := p.Users.GetUserWithSubscribers(id)
	if err != nil {
		http.Error(w, `{"error":"db1"}`, 500)
		return
	}
	var profile model.Profile
	if user.Is_author {
		levels, err := p.Levels.GetUserLevels(uint(id))
		if err != nil {
			http.Error(w, `{"error":"db"}`, 500)
			return
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
			http.Error(w, `{"error":"db"}`, 500)
			return
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

	result := Result{Body: BodyProfile{Profile: profile}}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&result)
	if err != nil {
		log.Println(err)
	}
}

func (p *ProfileHandler) ChangeUser(w http.ResponseWriter, r *http.Request) {
	AddAllowHeaders(w)
	if !auth.CheckAuthorization(r, p.Session) {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}

	// vars := mux.Vars(r)
	// _, err := strconv.Atoi(vars["id"])
	// if err != nil {
	// 	http.Error(w, `{"error":"bad id"}`, 400)
	// 	return
	// }

	body := http.MaxBytesReader(w, r.Body, maxBytesToRead)

	decoder := json.NewDecoder(body)
	user := &model.User{}

	err := decoder.Decode(user)
	if err != nil {
		http.Error(w, `{"error":"wrong_json"}`, http.StatusBadRequest)
		return
	}
	cookie, err := r.Cookie("session_id")
	if err != nil {
		http.Error(w, `{"error":"wrong_cookie"}`, http.StatusBadRequest)
		return
	}
	session, ok := p.Session.CheckSession(cookie.Value)
	if !ok {
		http.Error(w, `{"error":"wrong_session"}`, http.StatusBadRequest)
		return
	}
	if session.UserID != uint32(user.ID) {
		http.Error(w, `{"error":"wrong_change"}`, http.StatusBadRequest)
		return
	}

	err = p.Users.ChangeUser(*user)
	if err != nil {
		http.Error(w, `{"error":"db"}`, 500)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (p *ProfileHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	AddAllowHeaders(w)
	if !auth.CheckAuthorization(r, p.Session) {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, `{"error":"bad id"}`, 400)
		return
	}

	cookie, err := r.Cookie("session_id")
	if err != nil {
		http.Error(w, `{"error":"wrong_cookie"}`, http.StatusBadRequest)
		return
	}
	session, ok := p.Session.CheckSession(cookie.Value)
	if !ok {
		http.Error(w, `{"error":"wrong_session"}`, http.StatusBadRequest)
		return
	}
	if session.UserID != uint32(id) {
		http.Error(w, `{"error":"wrong_change"}`, http.StatusBadRequest)
		return
	}

	err = p.Users.DeleteUser(id)
	if err != nil {
		http.Error(w, `{"error":"db"}`, 500)
		return
	}

	w.WriteHeader(http.StatusOK)
}
