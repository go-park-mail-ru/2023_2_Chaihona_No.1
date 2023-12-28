package users_test

import (
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/configs"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/users"
)

var testUsers = []model.User{
	{
		Nickname:    "Aabbcc",
		Login:       "Aabbcc",
		Password:    "Aabbcc88**",
		Is_author:   false,
		Status:      "cheel",
		Description: "kefteme",
	},
}

type io struct {
	Id    int
	Login string
	User  model.User
	Users []model.User
}

type testCase struct {
	Columns  []string
	DBMethod string
	Prepare  func(mock sqlmock.Sqlmock)
	In       io
	Out      io
}

var testCases = map[string]testCase{
	"Register user": {
		DBMethod: "Register",
		Columns:  []string{"id"},
		Out: io{
			Id: 5,
		},
		In: io{
			User: testUsers[0],
		},
		Prepare: func(mock sqlmock.Sqlmock) {
			mock.ExpectExec("INSERT INTO "+configs.UserTable).
				WithArgs(testUsers[0].Nickname,
					testUsers[0].Login,
					testUsers[0].Password,
					testUsers[0].Is_author,
					testUsers[0].Status,
					testUsers[0].Description).
				WillReturnResult(sqlmock.NewResult(1, 1))
		},
	},
}

func TestUserStorage(t *testing.T) {
	for caseName, testCase := range testCases {
		testCase := testCase
		caseName := caseName
		t.Run(caseName, func(t *testing.T) {
			t.Parallel()
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			if testCase.Prepare != nil {
				testCase.Prepare(mock)
			}

			userStorage := users.CreateUserStorage(db)

			userId, err := userStorage.RegisterNewUser(&testCase.In.User)
			if err != nil {
				t.Errorf("unexpected error %s", err)
			}

			if !reflect.DeepEqual(userId, testCase.Out.Id) {
				t.Errorf("[%s] wrong Response: got %d, expected %d",
					caseName, userId, testCase.Out.Id)
			}
		})
	}
}
