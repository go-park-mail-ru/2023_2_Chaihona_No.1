package test_data

import (
	"project/model"
)

var Users = []model.User{
	{
		ID:       1,
		Login:    "Serezha",
		Password: "12345",
		UserType: model.CreatorStatus,
		Status:   `Двигаюсь на спокойном`,
	},
	{
		ID:       2,
		Login:    "Gosha",
		Password: "12345",
		UserType: model.CreatorStatus,
		Status:   `Шурую на фронте`,
	},
	{
		ID:       3,
		Login:    "Sanya",
		Password: "12345",
		UserType: model.CreatorStatus,
		Status:   `Кайфую на бэке`,
	},
	{
		ID:       4,
		Login:    "Umalat",
		Password: "12345",
		UserType: model.CreatorStatus,
		Status:   `Кефтеме`,
	},
	{
		ID:       5,
		Login:    "Simple_chel",
		Password: "12345",
		UserType: model.SimpleUserStatus,
		Status:   `Кефтеме`,
	},
}

var Posts = []model.Post{
	{
		ID:           1,
		AuthorID:     1,
		Access:       model.EveryoneAccess,
		CreationDate: "01.10.2023 18:46",
		Header:       "Всем привет, это моя страница в Копилке!",
		Body:         "Здесь будет очень много интересного, я буду выкладывать свои стихи и прозу, а также выкладывать видео с литературными разборами. Подписывайся!",
		Likes:        19,
		Comments: []model.Comment{
			{
				ID: 1,
				User: model.User{
					ID:       5,
					Login:    "Simple_chel",
					Password: "12345",
					UserType: model.SimpleUserStatus,
					Status:   `Кефтеме`,
				},
				Text:         "Вау! Круто! Уже купил подписку.",
				CreationDate: "01.10.2023 18:51",
			},
		},
		Tags: []model.Tag{
			{
				ID:   1,
				Name: "Литература",
			},
		},
	},
	{
		ID:           2,
		AuthorID:     1,
		Access:       model.SubscribersAccess,
		CreationDate: "01.10.2023 19:00",
		Header:       "Новый рассказ!",
		Body: `У американца, немца и армянина спрашивают, чтобы вы сделали, если бы у вас была машина времени. Американец:
		— Я бы побывал в 1945 и сделал бы так, чтобы не было атомных бомбежек Японии, это такой позор для моей нации
		Немец:
		— Я бы перебрался бы во времена 1939 года и сделал бы так, чтобы не было Второй мировой, это такой позор для моей нации
		Армянин:
		— Я бы походил по-другому`,
		Likes: 190,
		Comments: []model.Comment{
			{
				ID: 2,
				User: model.User{
					ID:       5,
					Login:    "Simple_chel",
					Password: "12345",
					UserType: model.SimpleUserStatus,
					Status:   `Кефтеме`,
				},
				Text:         "Впечатляет",
				CreationDate: "01.10.2023 18:51",
			},
		},
		Tags: []model.Tag{
			{
				ID:   1,
				Name: "Литература",
			},
		},
	},
}
