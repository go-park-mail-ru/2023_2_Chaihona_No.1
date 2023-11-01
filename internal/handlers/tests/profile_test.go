package handlers_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"

	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/handlers"
	mocks "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/handlers/mock_model"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/sessions"
)

var TestProfiles = []model.Profile{
	{
		User: model.User{
			ID:       5,
			Login:    "chert",
			UserType: model.SimpleUserStatus,
			Status:   "На приколе",
		},
		Description: "Первый в России, кто придумал торговать арбузами",
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
		Donated:  100,
		Currency: model.RubCurrency,
	},
	{
		User: model.User{
			ID:       0,
			Login:    "Serezha",
			Password: "12345",
			UserType: model.CreatorStatus,
			Status:   `Двигаюсь на спокойном`,
		},
		Description:     "Первый в России, кто придумал торговать арбузами",
		Subscribers:     1,
		SubscribeLevels: SubscrebeLevels,
		Goals:           Goals,
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

var Goals = []model.Goal{
	{
		ID:          1,
		GoalType:    model.SubcribersGoalType,
		Current:     0,
		GoalValue:   100,
		Description: "Как набирется 100 подпищиков орагинзую сходку",
	},
	{
		ID:          2,
		GoalType:    model.MoneyGoalType,
		Currency:    model.RubCurrency,
		Current:     1000,
		GoalValue:   100000,
		Description: "Куплю много шоколадок",
	},
}

var ProfileTestCases = map[string]TestCase{
	"get simple user's profile": {
		ID: "5",
		Response: JSONEncode(handlers.Result{Body: map[string]interface{}{
			"profiles": TestProfiles[0],
		}}),
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

			repos.Profile.EXPECT().GetProfile(uint(5)).Return(&TestProfiles[0], true).AnyTimes()
		},
	},
	"get creator's profile": {
		ID: "5",
		Response: JSONEncode(handlers.Result{Body: map[string]interface{}{
			"profiles": TestProfiles[1],
		}}),
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

			repos.Profile.EXPECT().GetProfile(uint(5)).Return(&TestProfiles[1], true).AnyTimes()
		},
	},
}

func TestGetProfileInfo(t *testing.T) {
	for caseName, testCase := range ProfileTestCases {
		testCase := testCase
		t.Run(caseName, func(t *testing.T) {
			t.Parallel()
			url := fmt.Sprintf("/api/v1/profile/%s", testCase.ID)
			req := httptest.NewRequest("GET", url, nil)
			w := httptest.NewRecorder()
			req.AddCookie(&testCase.Cookie)

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepos := MockRepos{
				Sessions: mocks.NewMockSessionRepository(ctrl),
				Profile:  mocks.NewMockProfileRepository(ctrl),
			}

			if testCase.Prepare != nil {
				testCase.Prepare(&mockRepos)
			}

			// ProfileHandler := handlers.CreateProfileHandlerViaRepos(
			// 	mockRepos.Sessions,
			// 	mockRepos.Profile,
			// )

			router := mux.NewRouter()
			// router.HandleFunc("/api/v1/profile/{id:[0-9]+}", ProfileHandler.GetInfo).
			// Methods("GET")
			router.ServeHTTP(w, req)

			if w.Code != testCase.StatusCode {
				t.Errorf("[%s] wrong StatusCode: got %d, expected %d",
					caseName, w.Code, testCase.StatusCode)
			}

			resp := w.Result()
			body, _ := ioutil.ReadAll(resp.Body)

			bodyStr := string(body)
			if !reflect.DeepEqual(bodyStr[:len(body)-1], testCase.Response) {
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
		})
	}
}
