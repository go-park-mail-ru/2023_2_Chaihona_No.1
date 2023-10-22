package handlers_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"

	// "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/authorization"
	// "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/handlers"
	// mocks "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/handlers/mock_model"
	// "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/model"
)

type AuthorizathionTestCase struct {
	Response   string
	User       TestUser
	APIMethod  string
	Prepare    func(repos *MockRepos)
	StatusCode int
	Cookie     http.Cookie
}

type TestUser struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	UserType string `json:"user_type,omitempty"`
}

var TestUsers = []TestUser{
	{
		Login:    "chertila",
		Password: "AaBb888**",
		UserType: model.SimpleUserStatus,
	},
	{
		Login:    "chertila",
		Password: "",
		UserType: model.SimpleUserStatus,
	},
	{
		Login:    "chertila",
		Password: "AaBb888**",
	},
}

var SignupTestCases = map[string]AuthorizathionTestCase{
	"successful simple user sign up": {
		User:      TestUsers[0],
		APIMethod: "registration",
		Response: JSONEncode(handlers.Result{Body: map[string]interface{}{
			"id": 0,
		}}),
		StatusCode: http.StatusOK,
		Prepare: func(repos *MockRepos) {
			repos.Users.EXPECT().RegisterNewUser(&model.User{
				Login:    TestUsers[0].Login,
				Password: TestUsers[0].Password,
				UserType: TestUsers[0].UserType,
			}).Return(nil).AnyTimes()

			repos.Profile.EXPECT().RegisterNewProfile(&model.Profile{
				User: model.User{
					Login:    TestUsers[0].Login,
					Password: TestUsers[0].Password,
					UserType: TestUsers[0].UserType,
				},
			}).Return(nil).AnyTimes()

			repos.Sessions.EXPECT().RegisterNewSession(gomock.Any()).AnyTimes()
		},
	},
	"unsuccessful simple user sign up": {
		User:       TestUsers[1],
		APIMethod:  "registration",
		Response:   `{"error":"user_validation"}`,
		StatusCode: http.StatusBadRequest,
	},
	"successful simple user login": {
		User:      TestUsers[2],
		APIMethod: "login",
		Response: JSONEncode(handlers.Result{Body: map[string]interface{}{
			"id": 0,
		}}),
		StatusCode: http.StatusOK,
		Prepare: func(repos *MockRepos) {
			repos.Users.EXPECT().CheckUser("chertila").Return(&model.User{
				Login:    TestUsers[2].Login,
				Password: TestUsers[2].Password,
				UserType: TestUsers[2].UserType,
			}, true).AnyTimes()

			repos.Sessions.EXPECT().RegisterNewSession(gomock.Any()).AnyTimes()
		},
	},
	"unsuccessful non-existent user login": {
		User:       TestUsers[2],
		APIMethod:  "login",
		Response:   `{"error":"wrong_input"}`,
		StatusCode: http.StatusBadRequest,
		Prepare: func(repos *MockRepos) {
			repos.Users.EXPECT().CheckUser("chertila").Return(nil, false).AnyTimes()
		},
	},
	"unsuccessful simple user login": {
		User:       TestUsers[2],
		APIMethod:  "login",
		Response:   `{"error":"wrong_input"}`,
		StatusCode: http.StatusBadRequest,
		Prepare: func(repos *MockRepos) {
			repos.Users.EXPECT().CheckUser("chertila").Return(&model.User{
				Login:    TestUsers[2].Login,
				Password: "bla bla bla",
				UserType: TestUsers[2].UserType,
			}, true).AnyTimes()
		},
	},
	"successful simple user logout": {
		User:      TestUsers[2],
		APIMethod: "logout",
		Response: JSONEncode(handlers.Result{Body: map[string]interface{}{
			"id": 0,
		}}),
		StatusCode: http.StatusOK,
		Prepare: func(repos *MockRepos) {
			repos.Sessions.EXPECT().DeleteSession("chertila").Return(nil).AnyTimes()
		},
		Cookie: http.Cookie{
			Name:     "session_id",
			Value:    "chertila",
			Expires:  time.Now().Add(10 * time.Hour),
			HttpOnly: true,
		},
	},
	"unsuccessful simple user logout: wrong Cookie": {
		User:       TestUsers[2],
		APIMethod:  "logout",
		Response:   `{"error":"user_logout"}`,
		StatusCode: http.StatusBadRequest,
		Cookie: http.Cookie{
			Name:     "bla bla bla",
			Value:    "chertila",
			Expires:  time.Now().Add(10 * time.Hour),
			HttpOnly: true,
		},
	},
	"unsuccessful simple user logout: non-existent session": {
		User:       TestUsers[2],
		APIMethod:  "logout",
		Response:   `{"error":"user_logout"}`,
		StatusCode: http.StatusInternalServerError,
		Prepare: func(repos *MockRepos) {
			repos.Sessions.EXPECT().
				DeleteSession("chertila").
				Return(authorization.ErrNoSuchSession).
				AnyTimes()
		},
		Cookie: http.Cookie{
			Name:     "session_id",
			Value:    "chertila",
			Expires:  time.Now().Add(10 * time.Hour),
			HttpOnly: true,
		},
	},
}

func TestAuthorization(t *testing.T) {
	for caseName, testCase := range SignupTestCases {
		testCase := testCase
		t.Run(caseName, func(t *testing.T) {
			t.Parallel()
			url := "/api/v1/" + testCase.APIMethod
			postBody := httptest.NewRecorder().Body
			postBody.Write([]byte(JSONEncode(testCase.User)))
			req := httptest.NewRequest("POST", url, postBody)
			w := httptest.NewRecorder()

			if testCase.Cookie.Name != "" {
				req.AddCookie(&testCase.Cookie)
			}

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepos := MockRepos{
				Users:    mocks.NewMockUserRepository(ctrl),
				Sessions: mocks.NewMockSessionRepository(ctrl),
				Profile:  mocks.NewMockProfileRepository(ctrl),
			}

			if testCase.Prepare != nil {
				testCase.Prepare(&mockRepos)
			}

			authHandler := handlers.CreateRepoHandler(
				mockRepos.Sessions,
				mockRepos.Users,
				mockRepos.Profile,
			)

			router := mux.NewRouter()
			router.HandleFunc("/api/v1/registration", authHandler.Signup).Methods("POST")
			router.HandleFunc("/api/v1/login", authHandler.Login).Methods("POST")
			router.HandleFunc("/api/v1/logout", authHandler.Logout).Methods("POST")
			router.ServeHTTP(w, req)

			if w.Code != testCase.StatusCode {
				t.Errorf("[%s] wrong StatusCode: got %d, expected %d",
					caseName, w.Code, testCase.StatusCode)
			}

			resp := w.Result()
			body, err := ioutil.ReadAll(resp.Body)

			bodyStr := string(body)
			if err != nil && !reflect.DeepEqual(bodyStr[:len(body)-1], testCase.Response) {
				t.Errorf("[%s] wrong Response: got %+v, expected %+v",
					caseName, bodyStr, testCase.Response)
			}
			if cookieHeader := resp.Header.Get("Set-Cookie"); testCase.StatusCode < 400 &&
				cookieHeader == "" {
				t.Errorf(
					"[%s] wrong Set-Cookie header: header is empty",
					caseName,
				)
			}
		})
	}
}
