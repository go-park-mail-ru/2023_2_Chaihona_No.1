package test_data

import (
	"project/model"
)

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
	{
		ID:       6,
		Login:    "Simple_chel_s_podpiskoy",
		Password: "12345",
		UserType: model.SimpleUserStatus,
		Status:   `Кефтеме`,
	},
}

var Profiles = []model.Profile{
	{
		User: model.User{
			ID:       1,
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
	{
		User: model.User{
			ID:       2,
			Login:    "Gosha",
			Password: "12345",
			UserType: model.CreatorStatus,
			Status:   `Шурую на фронте`,
		},
		Description: "Второй в России, кто придумал торговать арбузами",
		Subscribers: 1,
	},
	{
		User: model.User{
			ID:       3,
			Login:    "Sanya",
			Password: "12345",
			UserType: model.CreatorStatus,
			Status:   `Кайфую на бэке`,
		},
		Description: "Третий в России, кто придумал торговать арбузами",
		Subscribers: 1,
	},
	{
		User: model.User{
			ID:       1,
			Login:    "Serezha",
			Password: "12345",
			UserType: model.CreatorStatus,
			Status:   `Двигаюсь на спокойном`,
		},
		Description: "Первый в России, кто придумал торговать арбузами",
		Subscribers: 1,
	},
	{
		User: model.User{
			ID:       4,
			Login:    "Umalat",
			Password: "12345",
			UserType: model.CreatorStatus,
			Status:   `Кефтеме`,
		},
		Description: "Четвертый в России, кто придумал торговать арбузами",
		Subscribers: 1,
	},
	{
		User: model.User{
			ID:       5,
			Login:    "Simple_chel",
			Password: "12345",
			UserType: model.SimpleUserStatus,
			Status:   `Кефтеме`,
		},
		Subscribtions: []model.User{
			{
				ID:       1,
				Login:    "Serezha",
				UserType: model.CreatorStatus,
				Status:   `Двигаюсь на спокойном`,
			},
			{
				ID:       2,
				Login:    "Gosha",
				UserType: model.CreatorStatus,
				Status:   `Шурую на фронте`,
			},
			{
				ID:       3,
				Login:    "Sanya",
				UserType: model.CreatorStatus,
				Status:   `Кайфую на бэке`,
			},
			{
				ID:       4,
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
			ID:       4,
			Login:    "Umalat",
			Password: "12345",
			UserType: model.CreatorStatus,
			Status:   `Кефтеме`,
		},
		Description: "Четвертый в России, кто придумал торговать арбузами",
		Subscribers: 1,
	},
	{
		User: model.User{
			ID:       6,
			Login:    "Simple_chel_s_podpiskoy",
			Password: "12345",
			UserType: model.SimpleUserStatus,
			Status:   `Кефтеме`,
		},
		Donated:  50,
		Currency: model.RubCurrency,
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
