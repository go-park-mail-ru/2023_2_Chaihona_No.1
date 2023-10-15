package handlers_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"

	auth "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/authorization"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/handlers"
	mocks "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/handlers/mock_model"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/model"
)

type MockRepos struct {
	Users    *mocks.MockUserRepository
	Sessions *mocks.MockSessionRepository
	Posts    *mocks.MockPostRepository
	Profile  *mocks.MockProfileRepository
}

type TestCase struct {
	ID         string
	Response   string
	Prepare    func(repos *MockRepos)
	StatusCode int
	Cookie     http.Cookie
}

func JSONEncode(posts interface{}) string {
	res, _ := json.Marshal(posts)
	return string(res)
}

var PostTestCases = map[string]TestCase{
	"get simple post": {
		ID: "5",
		Response: JSONEncode(handlers.Result{Body: map[string]interface{}{
			"posts": []model.Post{
				{
					ID:           9,
					AuthorID:     5,
					HasAccess:    true,
					Access:       model.EveryoneAccess,
					CreationDate: "15:08 30.09.2023",
					Header:       "Header",
					Body:         "Body",
					Likes:        10,
				},
			},
		}}),
		StatusCode: 200,
		Cookie: http.Cookie{
			Name:     "session_id",
			Value:    "chertila",
			Expires:  time.Now().Add(10 * time.Hour),
			HttpOnly: true,
		},
		Prepare: func(repos *MockRepos) {
			repos.Posts.EXPECT().GetPostsByAuthorId(uint(5)).Return([]model.Post{
				{
					ID:           9,
					AuthorID:     5,
					Access:       model.EveryoneAccess,
					CreationDate: "15:08 30.09.2023",
					Header:       "Header",
					Body:         "Body",
					Likes:        10,
				},
			}, nil).AnyTimes()

			repos.Sessions.EXPECT().CheckSession("chertila").Return(&auth.Session{
				SessionId: "chertila",
				UserId:    9,
				Ttl:       time.Now().Add(10 * time.Hour),
			}, true).AnyTimes()

			repos.Profile.EXPECT().GetProfile(uint(9)).Return(&model.Profile{
				User: model.User{
					ID:       9,
					Login:    "chert",
					UserType: model.SimpleUserStatus,
				},
				Subscriptions: []model.User{{ID: 3}},
			}, true).AnyTimes()
		},
	},
	"get post for subscriber without access": {
		ID: "6",
		Response: JSONEncode(handlers.Result{Body: map[string]interface{}{
			"posts": []model.Post{
				{
					ID:           13,
					AuthorID:     6,
					HasAccess:    false,
					Access:       model.SubscribersAccess,
					Reason:       model.LowLevelReason,
					MinSubLevel:  3,
					CreationDate: "15:08 30.09.2023",
					Header:       "Header",
					Likes:        10,
				},
			},
		}}),
		StatusCode: 200,
		Cookie: http.Cookie{
			Name:     "session_id",
			Value:    "chertila",
			Expires:  time.Now().Add(10 * time.Hour),
			HttpOnly: true,
		},
		Prepare: func(repos *MockRepos) {
			repos.Posts.EXPECT().GetPostsByAuthorId(uint(6)).Return([]model.Post{
				{
					ID:           13,
					AuthorID:     6,
					Access:       model.SubscribersAccess,
					MinSubLevel:  3,
					CreationDate: "15:08 30.09.2023",
					Header:       "Header",
					Body:         "Body",
					Likes:        10,
				},
			}, nil).AnyTimes()

			repos.Sessions.EXPECT().CheckSession("chertila").Return(&auth.Session{
				SessionId: "chertila",
				UserId:    9,
				Ttl:       time.Now().Add(10 * time.Hour),
			}, true).AnyTimes()

			repos.Profile.EXPECT().GetProfile(uint(9)).Return(&model.Profile{
				User: model.User{
					ID:       9,
					Login:    "chert",
					UserType: model.SimpleUserStatus,
				},
				Subscriptions: []model.User{{ID: 3}},
			}, true).AnyTimes()
		},
	},
	"get post for subscribers with access": {
		ID: "7",
		Response: JSONEncode(handlers.Result{Body: map[string]interface{}{
			"posts": []model.Post{
				{
					ID:           17,
					AuthorID:     7,
					HasAccess:    true,
					Access:       model.SubscribersAccess,
					MinSubLevel:  2,
					CreationDate: "15:08 30.09.2023",
					Header:       "Header",
					Body:         "Body",
					Likes:        10,
				},
			},
		}}),
		StatusCode: 200,
		Cookie: http.Cookie{
			Name:     "session_id",
			Value:    "chertila",
			Expires:  time.Now().Add(10 * time.Hour),
			HttpOnly: true,
		},
		Prepare: func(repos *MockRepos) {
			repos.Posts.EXPECT().GetPostsByAuthorId(uint(7)).Return([]model.Post{
				{
					ID:           17,
					AuthorID:     7,
					Access:       model.SubscribersAccess,
					MinSubLevel:  2,
					CreationDate: "15:08 30.09.2023",
					Header:       "Header",
					Body:         "Body",
					Likes:        10,
				},
			}, nil).AnyTimes()

			repos.Sessions.EXPECT().CheckSession("chertila").Return(&auth.Session{
				SessionId: "chertila",
				UserId:    9,
				Ttl:       time.Now().Add(10 * time.Hour),
			}, true).AnyTimes()

			repos.Profile.EXPECT().GetProfile(uint(9)).Return(&model.Profile{
				User: model.User{
					ID:       9,
					Login:    "chert",
					UserType: model.SimpleUserStatus,
				},
				Subscriptions: []model.User{{ID: 7}},
			}, true).AnyTimes()
		},
	},
	"get simple user's post": {
		ID:         "4",
		Response:   `{"error":"user isn't author"}`,
		StatusCode: 400,
		Cookie: http.Cookie{
			Name:     "session_id",
			Value:    "chertila",
			Expires:  time.Now().Add(10 * time.Hour),
			HttpOnly: true,
		},
		Prepare: func(repos *MockRepos) {
			repos.Posts.EXPECT().
				GetPostsByAuthorId(uint(4)).
				Return([]model.Post{}, &model.ErrorPost{Err: handlers.ErrorNotAuthor, StatusCode: http.StatusBadRequest}).
				AnyTimes()
			repos.Sessions.EXPECT().CheckSession("chertila").Return(&auth.Session{
				SessionId: "chertila",
				UserId:    9,
				Ttl:       time.Now().Add(10 * time.Hour),
			}, true).AnyTimes()
		},
	},
}

func TestGetPosts(t *testing.T) {
	for caseName, testCase := range PostTestCases {
		testCase := testCase
		t.Run(caseName, func(t *testing.T) {
			t.Parallel()
			url := fmt.Sprintf("/api/v1/profile/%s/post", testCase.ID)
			req := httptest.NewRequest("GET", url, nil)
			w := httptest.NewRecorder()
			req.AddCookie(&testCase.Cookie)

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepos := MockRepos{
				Sessions: mocks.NewMockSessionRepository(ctrl),
				Posts:    mocks.NewMockPostRepository(ctrl),
				Profile:  mocks.NewMockProfileRepository(ctrl),
			}

			if testCase.Prepare != nil {
				testCase.Prepare(&mockRepos)
			}

			PostHandler := handlers.CreatePostHandlerViaRepos(
				mockRepos.Sessions,
				mockRepos.Posts,
				mockRepos.Profile,
			)

			router := mux.NewRouter()
			router.HandleFunc("/api/v1/profile/{id:[0-9]+}/post", PostHandler.GetAllUserPosts).
				Methods("GET")
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

			if contentHeader := resp.Header.Get("Content-Type"); testCase.StatusCode < 400 &&
				contentHeader != "application/json" {
				t.Errorf(
					"[%s] wrong Content-Type header: got %s, expected application/json",
					caseName,
					contentHeader,
				)
			}
		})
	}
}
