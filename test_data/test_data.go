package test_data

import (
	"project/model"
)

var Users = []model.User{
	model.User{
		ID:       1,
		Login:    "Serezha",
		Password: "12345",
		UserType: model.CreatorStatus,
		Status:   `Двигаюсь на спокойном`,
	},
}
