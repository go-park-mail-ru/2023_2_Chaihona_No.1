package subscriptions

import (
	model "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"
)

type SubscriptionRepository interface {
	AddNewSubscription(subscription model.Subscription) (int, error)
	DeleteSubscription(levelId int, subId int) error
	GetSubscription(id int) (model.Subscription, error)
	GetUserSubscriptions(id int) ([]model.User, error)
	ChangeSubscription(subscription model.Subscription) error
	CountSubscribers(id int) (int, error)
	GetAllNotFreeSubscriptions() ([]model.Subscription, error)
}