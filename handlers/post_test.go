package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	auth "project/authorization"
	mocks "project/handlers/mock_model"
	"project/model"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
)

type MockRepos struct {
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

var TestCases = map[string]TestCase{
	"get simple post": TestCase{
		ID: "5",
		Response: JSONEncode(Result{Body: map[string]interface{}{
			"posts": []model.Post{
				model.Post{
					ID:           9,
					AuthorID:     5,
					HasAccess:    true,
					Access:       model.EveryoneAccess,
					CreationDate: "15:08 30.09.2023",
					Header:       "Header",
					Body:         "Body",
					Likes:        10,
				},
			}}}),
		StatusCode: 200,
		Cookie: http.Cookie{
			Name:     "session_id",
			Value:    "chertila",
			Expires:  time.Now().Add(10 * time.Hour),
			HttpOnly: true,
		},
		Prepare: func(repos *MockRepos) {
			repos.Posts.EXPECT().GetPostsByAuthorId(uint(5)).Return(&[]model.Post{
				model.Post{
					ID:           9,
					AuthorID:     5,
					Access:       model.EveryoneAccess,
					CreationDate: "15:08 30.09.2023",
					Header:       "Header",
					Body:         "Body",
					Likes:        10,
				},
			}, nil).AnyTimes()

			repos.Sessions.EXPECT().CheckSession("chertila").Return(auth.Session{
				SessionId: "chertila",
				UserId:    9,
				Ttl:       time.Now().Add(10 * time.Hour),
			}, true).AnyTimes()

			repos.Profile.EXPECT().GetProfile(uint(9)).Return(&model.Profile{
				ID: 9,
				User: model.User{
					ID:       9,
					Login:    "chert",
					UserType: model.SimpleUserStatus,
				},
				Subscribtions: []model.User{model.User{ID: 3}},
			}, true).AnyTimes()
		},
	},
	// "get one-time payment post without access": TestCase{
	// 	ID: "5",
	// 	Response: JSONEncode(model.Post{
	// 		ID:           11,
	// 		AuthorID:     5,
	// 		HasAccess:    false,
	// 		Access:       model.OneTimePaymentAccess,
	// 		Reason:       model.UnpaidReason,
	// 		Payment:      100,
	// 		Currency:     model.Currency,
	// 		CreationDate: "15:08 30.09.2023",
	// 		Header:       "Header",
	// 		Likes:        10,
	// 	}),
	// 	StatusCode: 200,
	// },
	"get post for subscriber without access": TestCase{
		ID: "6",
		Response: JSONEncode(Result{Body: map[string]interface{}{
			"posts": []model.Post{
				model.Post{
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
			}}}),
		StatusCode: 200,
		Cookie: http.Cookie{
			Name:     "session_id",
			Value:    "chertila",
			Expires:  time.Now().Add(10 * time.Hour),
			HttpOnly: true,
		},
		Prepare: func(repos *MockRepos) {
			repos.Posts.EXPECT().GetPostsByAuthorId(uint(6)).Return(&[]model.Post{
				model.Post{
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

			repos.Sessions.EXPECT().CheckSession("chertila").Return(auth.Session{
				SessionId: "chertila",
				UserId:    9,
				Ttl:       time.Now().Add(10 * time.Hour),
			}, true).AnyTimes()

			repos.Profile.EXPECT().GetProfile(uint(9)).Return(&model.Profile{
				ID: 9,
				User: model.User{
					ID:       9,
					Login:    "chert",
					UserType: model.SimpleUserStatus,
				},
				Subscribtions: []model.User{model.User{ID: 3}},
			}, true).AnyTimes()
		},
	},
	// // "get one-time payment post with access": TestCase{
	// // 	ID: "5",
	// // 	Response: JSONEncode(model.Post{
	// // 		ID:           15,
	// // 		AuthorID:     5,
	// // 		HasAccess:    true,
	// // 		Access:       model.OneTimePaymentAccess,
	// // 		Payment:      100,
	// // 		Currency:     model.Currency,
	// // 		CreationDate: "15:08 30.09.2023",
	// // 		Header:       "Header",
	// // 		Body:         "Body",
	// // 		Likes:        10,
	// // 	}),
	// // 	StatusCode: 200,
	// // },
	"get post for subscribers with access": TestCase{
		ID: "7",
		Response: JSONEncode(Result{Body: map[string]interface{}{
			"posts": []model.Post{
				model.Post{
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
			}}}),
		StatusCode: 200,
		Cookie: http.Cookie{
			Name:     "session_id",
			Value:    "chertila",
			Expires:  time.Now().Add(10 * time.Hour),
			HttpOnly: true,
		},
		Prepare: func(repos *MockRepos) {
			repos.Posts.EXPECT().GetPostsByAuthorId(uint(7)).Return(&[]model.Post{
				model.Post{
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

			repos.Sessions.EXPECT().CheckSession("chertila").Return(auth.Session{
				SessionId: "chertila",
				UserId:    9,
				Ttl:       time.Now().Add(10 * time.Hour),
			}, true).AnyTimes()

			repos.Profile.EXPECT().GetProfile(uint(9)).Return(&model.Profile{
				ID: 9,
				User: model.User{
					ID:       9,
					Login:    "chert",
					UserType: model.SimpleUserStatus,
				},
				Subscribtions: []model.User{model.User{ID: 7}},
			}, true).AnyTimes()
		},
	},
	"get simple user's post": TestCase{
		ID:         "4",
		Response:   `{"error":"bad request"}`,
		StatusCode: 400,
		Cookie: http.Cookie{
			Name:     "session_id",
			Value:    "chertila",
			Expires:  time.Now().Add(10 * time.Hour),
			HttpOnly: true,
		},
		Prepare: func(repos *MockRepos) {
			repos.Posts.EXPECT().GetPostsByAuthorId(uint(4)).Return(&[]model.Post{}, NotAuthorError).AnyTimes()
		},
	},

	"get non-existent user posts": TestCase{
		ID:         "bla",
		Response:   `{"error":"bad id"}`,
		StatusCode: 400,
		Cookie: http.Cookie{
			Name:     "session_id",
			Value:    "chertila",
			Expires:  time.Now().Add(10 * time.Hour),
			HttpOnly: true,
		},
	},
}

func TestGetPosts(t *testing.T) {
	for caseName, testCase := range TestCases {
		url := "https://api/v1/profile/" + testCase.ID + "/post"
		req := httptest.NewRequest("GET", url, nil)
		w := httptest.NewRecorder()
		req.AddCookie(&testCase.Cookie)

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockRepos := MockRepos{
			mocks.NewMockSessionRepository(ctrl),
			mocks.NewMockPostRepository(ctrl),
			mocks.NewMockProfileRepository(ctrl),
		}

		if testCase.Prepare != nil {
			testCase.Prepare(&mockRepos)
		}

		PostHandler := &PostHandler{
			mockRepos.Sessions,
			mockRepos.Posts,
			mockRepos.Profile,
		}

		PostHandler.GetAllUserPosts(w, req)

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
	}
}
