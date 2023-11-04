package handlers_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"

	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/handlers"
	mocks "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/handlers/tests/mock_model"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/sessions"
)

type ProfileTestCase struct {
	ID         string
	Response   handlers.Result
	User       model.User
	Prepare    func(repos *MockRepos)
	StatusCode int
	Cookie     http.Cookie
	MethodHTTP string
	MethodAPI  string
}

var TestProfiles = []model.Profile{
	{
		User: model.User{
			ID:          5,
			Login:       "chert",
			UserType:    model.SimpleUserStatus,
			Is_author:   false,
			Status:      "На приколе",
			Description: "Первый в России, кто придумал торговать арбузами",
			Subscribers: 0,
		},
		Subscriptions: []model.User{
			{
				ID:       0,
				Login:    "Serezha",
				UserType: model.CreatorStatus,
				Status:   `Двигаюсь на спокойном`,
			},
			{
				ID:       1,
				Login:    "Gosha",
				UserType: model.CreatorStatus,
				Status:   `Шурую на фронте`,
			},
			{
				ID:       2,
				Login:    "Sanya",
				UserType: model.CreatorStatus,
				Status:   `Кайфую на бэке`,
			},
			{
				ID:       3,
				Login:    "Umalat",
				UserType: model.CreatorStatus,
				Status:   `Кефтеме`,
			},
		},
		// Donated:  100,
		// Currency: model.RubCurrency,
	},
	{
		User: model.User{
			ID:          5,
			Login:       "Serezha",
			Password:    "12345",
			UserType:    model.CreatorStatus,
			Is_author:   true,
			Status:      `Двигаюсь на спокойном`,
			Description: "Первый в России, кто придумал торговать арбузами",
			Subscribers: 1,
		},
		Subscribers:     1,
		SubscribeLevels: SubscrebeLevels,
	},
	{
		User: model.User{
			ID:          5,
			Login:       "Serezha",
			UserType:    model.CreatorStatus,
			Status:      `Двигаюсь на спокойном`,
			Description: "Первый в России, кто придумал торговать арбузами",
		},
	},
	{
		User: model.User{
			ID:          5,
			Login:       "chert",
			UserType:    model.SimpleUserStatus,
			Is_author:   false,
			Status:      "На приколе",
			Description: "Первый в России, кто придумал торговать арбузами",
		},
	},
}

var SubscrebeLevels = []model.SubscribeLevel{
	{
		ID:          1,
		Level:       1,
		Name:        "Почитатель",
		Description: "У тебя будет доступ ко всем моим произведениям",
		Payment:     1000,
		Currency:    model.RubCurrency,
	},
	{
		ID:          2,
		Level:       2,
		Name:        "Истинный ценитель",
		Description: "У тебя будет доступ к моим черновым вариантам и неиспользованным кускам",
		Payment:     5000,
		Currency:    model.RubCurrency,
	},
}

var ProfileTestCases = map[string]ProfileTestCase{
	"get simple user's profile": {
		MethodHTTP: "GET",
		MethodAPI:  "GetInfo",
		ID:         "5",
		Response: handlers.Result{Body: map[string]interface{}{
			"profile": TestProfiles[0],
		}},
		StatusCode: 200,
		Cookie: http.Cookie{
			Name:     "session_id",
			Value:    "chertila",
			Expires:  time.Now().Add(10 * time.Hour),
			HttpOnly: true,
		},
		Prepare: func(repos *MockRepos) {
			repos.Sessions.EXPECT().CheckSession("chertila").Return(&sessions.Session{
				SessionID: "chertila",
				UserID:    9,
				TTL:       time.Now().Add(10 * time.Hour),
			}, true).AnyTimes()

			repos.Users.EXPECT().GetUserWithSubscribers(5).Return(TestProfiles[0].User, nil).AnyTimes()
			repos.Subscriptions.EXPECT().GetUserSubscriptions(5).Return(TestProfiles[0].Subscriptions, nil).AnyTimes()
		},
	},
	"get creator's profile": {
		MethodHTTP: "GET",
		MethodAPI:  "GetInfo",
		ID:         "5",
		Response: handlers.Result{Body: map[string]interface{}{
			"profile": TestProfiles[1],
		}},
		StatusCode: 200,
		Cookie: http.Cookie{
			Name:     "session_id",
			Value:    "chertila",
			Expires:  time.Now().Add(10 * time.Hour),
			HttpOnly: true,
		},
		Prepare: func(repos *MockRepos) {
			repos.Sessions.EXPECT().CheckSession("chertila").Return(&sessions.Session{
				SessionID: "chertila",
				UserID:    9,
				TTL:       time.Now().Add(10 * time.Hour),
			}, true).AnyTimes()
			repos.Users.EXPECT().GetUserWithSubscribers(5).Return(TestProfiles[1].User, nil).AnyTimes()
			repos.Subscription_levels.EXPECT().GetUserLevels(uint(5)).Return(TestProfiles[1].SubscribeLevels, nil).AnyTimes()
		},
	},
	"successful change creator": {
		MethodHTTP: "POST",
		MethodAPI:  "ChangeUser",
		ID:         "5",
		StatusCode: 200,
		Cookie: http.Cookie{
			Name:     "session_id",
			Value:    "chertila",
			Expires:  time.Now().Add(10 * time.Hour),
			HttpOnly: true,
		},
		Prepare: func(repos *MockRepos) {
			repos.Sessions.EXPECT().CheckSession("chertila").Return(&sessions.Session{
				SessionID: "chertila",
				UserID:    5,
				TTL:       time.Now().Add(10 * time.Hour),
			}, true).AnyTimes()
			repos.Users.EXPECT().ChangeUser(TestProfiles[2].User).Return(nil).AnyTimes()
		},
		User: TestProfiles[2].User,
	},
	"unsuccessful change user": {
		MethodHTTP: "POST",
		MethodAPI:  "ChangeUser",
		ID:         "5",
		StatusCode: 400,
		Response: handlers.Result{Err: "wrong_change"},
		Cookie: http.Cookie{
			Name:     "session_id",
			Value:    "chertila",
			Expires:  time.Now().Add(10 * time.Hour),
			HttpOnly: true,
		},
		Prepare: func(repos *MockRepos) {
			repos.Sessions.EXPECT().CheckSession("chertila").Return(&sessions.Session{
				SessionID: "chertila",
				UserID:    9,
				TTL:       time.Now().Add(10 * time.Hour),
			}, true).AnyTimes()
			repos.Users.EXPECT().ChangeUser(TestProfiles[2].User).Return(nil).AnyTimes()
		},
		User: TestProfiles[2].User,
	},
	"successful change simple user": {
		MethodHTTP: "POST",
		MethodAPI:  "ChangeUser",
		ID:         "5",
		StatusCode: 200,
		Cookie: http.Cookie{
			Name:     "session_id",
			Value:    "chertila",
			Expires:  time.Now().Add(10 * time.Hour),
			HttpOnly: true,
		},
		Prepare: func(repos *MockRepos) {
			repos.Sessions.EXPECT().CheckSession("chertila").Return(&sessions.Session{
				SessionID: "chertila",
				UserID:    5,
				TTL:       time.Now().Add(10 * time.Hour),
			}, true).AnyTimes()
			repos.Users.EXPECT().ChangeUser(TestProfiles[3].User).Return(nil).AnyTimes()
		},
		User: TestProfiles[3].User,
	},
	"successful delete simple user": {
		MethodHTTP: "DELETE",
		MethodAPI:  "ChangeUser",
		ID:         "5",
		StatusCode: 200,
		Cookie: http.Cookie{
			Name:     "session_id",
			Value:    "chertila",
			Expires:  time.Now().Add(10 * time.Hour),
			HttpOnly: true,
		},
		Prepare: func(repos *MockRepos) {
			repos.Sessions.EXPECT().CheckSession("chertila").Return(&sessions.Session{
				SessionID: "chertila",
				UserID:    5,
				TTL:       time.Now().Add(10 * time.Hour),
			}, true).AnyTimes()
			repos.Users.EXPECT().DeleteUser(int(TestProfiles[3].User.ID)).Return(nil).AnyTimes()
		},
		User: TestProfiles[3].User,
	},
	"unsuccessful delete simple user": {
		MethodHTTP: "DELETE",
		MethodAPI:  "ChangeUser",
		ID:         "5",
		Response: handlers.Result{Err: "wrong_change"},
		StatusCode: 400,
		Cookie: http.Cookie{
			Name:     "session_id",
			Value:    "chertila",
			Expires:  time.Now().Add(10 * time.Hour),
			HttpOnly: true,
		},
		Prepare: func(repos *MockRepos) {
			repos.Sessions.EXPECT().CheckSession("chertila").Return(&sessions.Session{
				SessionID: "chertila",
				UserID:    9,
				TTL:       time.Now().Add(10 * time.Hour),
			}, true).AnyTimes()
			repos.Users.EXPECT().DeleteUser(int(TestProfiles[3].User.ID)).Return(nil).AnyTimes()
		},
		User: TestProfiles[3].User,
	},
}

func TestGetProfileInfo(t *testing.T) {
	for caseName, testCase := range ProfileTestCases {
		testCase := testCase
		caseName := caseName
		t.Run(caseName, func(t *testing.T) {
			t.Parallel()
			url := fmt.Sprintf("/api/v1/profile/%s", testCase.ID)
			w := httptest.NewRecorder()
			postBody := httptest.NewRecorder().Body
			if testCase.MethodHTTP == "POST" {
				body, err := json.Marshal(testCase.User)
				if err != nil {
					t.Errorf("%s", err)
				}
				postBody.Write([]byte(body))
			}
			req := httptest.NewRequest(testCase.MethodHTTP, url, postBody)
			req.AddCookie(&testCase.Cookie)

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepos := MockRepos{
				Sessions:            mocks.NewMockSessionRepository(ctrl),
				Users:               mocks.NewMockUserRepository(ctrl),
				Subscriptions:       mocks.NewMockSubscriptionRepository(ctrl),
				Subscription_levels: mocks.NewMockSubscribeLevelRepository(ctrl),
			}

			if testCase.Prepare != nil {
				testCase.Prepare(&mockRepos)
			}

			ProfileHandler := handlers.CreateProfileHandlerViaRepos(
				mockRepos.Sessions,
				mockRepos.Users,
				mockRepos.Subscription_levels,
				mockRepos.Subscriptions,
			)

			router := mux.NewRouter()
			router.HandleFunc("/api/v1/profile/{id:[0-9]+}", handlers.NewWrapper(ProfileHandler.GetInfoStrategy).ServeHTTP).
				Methods("GET")
			router.HandleFunc("/api/v1/profile/{id:[0-9]+}",  handlers.NewWrapper(ProfileHandler.ChangeUserStratagy).ServeHTTP).
				Methods("POST")
			router.HandleFunc("/api/v1/profile/{id:[0-9]+}",  handlers.NewWrapper(ProfileHandler.DeleteUserStratagy).ServeHTTP).
				Methods("DELETE")

			router.ServeHTTP(w, req)

			if w.Code != testCase.StatusCode {
				t.Errorf("[%s] wrong StatusCode: got %d, expected %d",
					caseName, w.Code, testCase.StatusCode)
			}

			if testCase.MethodAPI == "GetInfo" {
				resp := w.Result()
				body, _ := io.ReadAll(resp.Body)
				expected, err := json.Marshal(testCase.Response)
				if err != nil {
					t.Errorf("%s", err)
				}

				bodyStr := string(body)
				if !reflect.DeepEqual(bodyStr[:len(body)-1], expected) {
					t.Errorf("[%s] wrong Response: got %+v, expected %+v",
						caseName, bodyStr, testCase.Response)
				}
				if contentHeader := resp.Header.Get("Content-Type"); contentHeader != "application/json" {
					t.Errorf(
						"[%s] wrong Content-Type header: got %s, expected application/json",
						caseName,
						contentHeader,
					)
				}
			}

		})
	}
}