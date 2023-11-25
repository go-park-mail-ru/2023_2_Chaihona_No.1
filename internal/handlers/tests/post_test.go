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
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/posts"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/sessions"
)

type PostTestCase struct {
	ID         string
	Response   handlers.Result
	Prepare    func(repos *MockRepos)
	StatusCode int
	Cookie     http.Cookie
}

var PostTestCases = map[string]PostTestCase{
	"get simple post": {
		ID: "5",
		Response: handlers.Result{Body: map[string]interface{}{
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
		}},
		StatusCode: 200,
		Cookie: http.Cookie{
			Name:     "session_id",
			Value:    "chertila",
			Expires:  time.Now().Add(10 * time.Hour),
			HttpOnly: true,
		},
		Prepare: func(repos *MockRepos) {
			repos.Posts.EXPECT().GetPostsByAuthorId(uint(5), uint(7)).Return([]model.Post{
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

			repos.Sessions.EXPECT().CheckSession("chertila").Return(&sessions.SessionOld{
				SessionID: "chertila",
				UserID:    9,
				TTL:       time.Now().Add(10 * time.Hour),
			}, true).AnyTimes()

			// repos.Profile.EXPECT().GetProfile(uint(9)).Return(&model.Profile{
			// 	User: model.User{
			// 		ID:       9,
			// 		Login:    "chert",
			// 		UserType: model.SimpleUserStatus,
			// 	},
			// 	Subscriptions: []model.User{{ID: 3}},
			// }, true).AnyTimes()
		},
	},
	"get post for subscriber without access": {
		ID: "6",
		Response: handlers.Result{Body: map[string]interface{}{
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
		}},
		StatusCode: 200,
		Cookie: http.Cookie{
			Name:     "session_id",
			Value:    "chertila",
			Expires:  time.Now().Add(10 * time.Hour),
			HttpOnly: true,
		},
		Prepare: func(repos *MockRepos) {
			repos.Posts.EXPECT().GetPostsByAuthorId(uint(6), uint(7)).Return([]model.Post{
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

			repos.Sessions.EXPECT().CheckSession("chertila").Return(&sessions.SessionOld{
				SessionID: "chertila",
				UserID:    9,
				TTL:       time.Now().Add(10 * time.Hour),
			}, true).AnyTimes()

			// repos.Profile.EXPECT().GetProfile(uint(9)).Return(&model.Profile{
			// 	User: model.User{
			// 		ID:       9,
			// 		Login:    "chert",
			// 		UserType: model.SimpleUserStatus,
			// 	},
			// 	Subscriptions: []model.User{{ID: 3}},
			// }, true).AnyTimes()
		},
	},
	"get post for subscribers with access": {
		ID: "7",
		Response: handlers.Result{Body: map[string]interface{}{
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
		}},
		StatusCode: 200,
		Cookie: http.Cookie{
			Name:     "session_id",
			Value:    "chertila",
			Expires:  time.Now().Add(10 * time.Hour),
			HttpOnly: true,
		},
		Prepare: func(repos *MockRepos) {
			repos.Posts.EXPECT().GetPostsByAuthorId(uint(7), uint(7)).Return([]model.Post{
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

			repos.Sessions.EXPECT().CheckSession("chertila").Return(&sessions.SessionOld{
				SessionID: "chertila",
				UserID:    9,
				TTL:       time.Now().Add(10 * time.Hour),
			}, true).AnyTimes()

			// repos.Profile.EXPECT().GetProfile(uint(9)).Return(&model.Profile{
			// 	User: model.User{
			// 		ID:       9,
			// 		Login:    "chert",
			// 		UserType: model.SimpleUserStatus,
			// 	},
			// 	Subscriptions: []model.User{{ID: 7}},
			// }, true).AnyTimes()
		},
	},
	"get simple user's post": {
		ID: "4",
		// Response:   `{"error":"user isn't author"}`,
		Response:   handlers.Result{Err: "user isn't author"},
		StatusCode: 400,
		Cookie: http.Cookie{
			Name:     "session_id",
			Value:    "chertila",
			Expires:  time.Now().Add(10 * time.Hour),
			HttpOnly: true,
		},
		Prepare: func(repos *MockRepos) {
			repos.Posts.EXPECT().
				GetPostsByAuthorId(uint(4), uint(7)).
				Return([]model.Post{}, &posts.ErrorPost{Err: handlers.ErrorNotAuthor, StatusCode: http.StatusBadRequest}).
				AnyTimes()
			repos.Sessions.EXPECT().CheckSession("chertila").Return(&sessions.SessionOld{
				SessionID: "chertila",
				UserID:    9,
				TTL:       time.Now().Add(10 * time.Hour),
			}, true).AnyTimes()
		},
	},
}

func TestGetPosts(t *testing.T) {
	for caseName, testCase := range PostTestCases {
		testCase := testCase
		caseName := caseName
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
				Likes:    mocks.NewMockLikeRepository(ctrl),
			}

			if testCase.Prepare != nil {
				testCase.Prepare(&mockRepos)
			}

			PostHandler := handlers.CreatePostHandlerViaRepos(
				mockRepos.Sessions,
				mockRepos.Posts,
				mockRepos.Likes,
			)

			router := mux.NewRouter()
			router.HandleFunc("/api/v1/profile/{id:[0-9]+}/post", handlers.NewWrapper(PostHandler.GetAllUserPostsStrategy).ServeHTTP).
				Methods("GET")
			router.ServeHTTP(w, req)

			if w.Code != testCase.StatusCode {
				t.Errorf("[%s] wrong StatusCode: got %d, expected %d",
					caseName, w.Code, testCase.StatusCode)
			}

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
