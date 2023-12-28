package subscriptionlevels

import (
	model "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"
)

type SubscribeLevelRepository interface {
	AddNewLevel(level *model.SubscribeLevel) (model.SubscribeLevel, error)
	DeleteLevel(id uint) error
	GetUserLevels(userId uint) ([]model.SubscribeLevel, error)
	GetLevel(id uint) (model.SubscribeLevel, error)
}
