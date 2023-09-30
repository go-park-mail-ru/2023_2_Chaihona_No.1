package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"project/model"
	"testing"
)

type TestCase struct {
	ID         string
	Response   string
	StatusCode int
}

func JSONEncode(posts []model.Post) string {
	res, _ := json.Marshal(posts)
	return string(res)
}

var TestCases = map[string]TestCase{
	"get simple post": TestCase{
		ID: "5",
		Response: JSONEncode(model.Post{
			ID:           9,
			AuthorID:     5,
			HasAccess:    true,
			Access:       model.EveryoneAccess,
			CreationDate: "15:08 30.09.2023",
			Header:       "Header",
			Body:         "Body",
			Likes:        10,
		}),
		StatusCode: 200,
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
		Response: JSONEncode([]model.Post{model.Post{
			ID:           13,
			AuthorID:     6,
			HasAccess:    false,
			Access:       model.SubscribersAccess,
			Reason:       model.LowLevelReason,
			MinSubLevel:  3,
			CreationDate: "15:08 30.09.2023",
			Header:       "Header",
			Likes:        10,
		}}),
		StatusCode: 200,
	},
	// "get one-time payment post with access": TestCase{
	// 	ID: "5",
	// 	Response: JSONEncode(model.Post{
	// 		ID:           15,
	// 		AuthorID:     5,
	// 		HasAccess:    true,
	// 		Access:       model.OneTimePaymentAccess,
	// 		Payment:      100,
	// 		Currency:     model.Currency,
	// 		CreationDate: "15:08 30.09.2023",
	// 		Header:       "Header",
	// 		Body:         "Body",
	// 		Likes:        10,
	// 	}),
	// 	StatusCode: 200,
	// },
	"get post for subscribers with access": TestCase{
		ID: "5",
		Response: JSONEncode([]model.Post{model.Post{
			ID:           17,
			AuthorID:     5,
			HasAccess:    true,
			Access:       model.SubscribersAccess,
			MinSubLevel:  2,
			CreationDate: "15:08 30.09.2023",
			Header:       "Header",
			Body:         "Body",
			Likes:        10,
		}}),
		StatusCode: 200,
	},
	"get simple user's post": TestCase{
		ID:         "4",
		Response:   `{"error":"bad request"}`,
		StatusCode: 400,
	},
	"get non-existent user posts": TestCase{
		ID:         "bla",
		Response:   `{"error":"bad id"}`,
		StatusCode: 400,
	},
}

func TestGetPosts(t *testing.T) {
	for caseName, testCase := range TestCases {
		url := "htpps://api/v1/profile/" + testCase.ID + "/post"
		req := httptest.NewRequest("GET", url, nil)
		w := httptest.NewRecorder()

		PostHandler := CreateTestPostHandler() //Make this func using mocks

		PostHandler.GetAllUserPosts(w, req)

		if w.Code != testCase.StatusCode {
			t.Errorf("[%d] wrong StatusCode: got %d, expected %d",
				caseName, w.Code, testCase.StatusCode)
		}

		resp := w.Result()
		body, _ := ioutil.ReadAll(resp.Body)

		bodyStr := string(body)
		if bodyStr != testCase.Response {
			t.Errorf("[%d] wrong Response: got %+v, expected %+v",
				caseName, bodyStr, testCase.Response)
		}
	}
}
