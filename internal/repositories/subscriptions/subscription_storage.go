package subscriptions

import (
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/dbscan"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/configs"
	model "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"
)

type SubscriptionsStorage struct {
	db *sql.DB
}

func CountSubscribersSQL(id int) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(*)").
		From(configs.SubscriptionTable).
		Where(squirrel.Eq{"creator_id": id}).
		PlaceholderFormat(squirrel.Dollar)
}

func SelectUserSubscriptionsSQL(id int) squirrel.SelectBuilder {
	return squirrel.Select().
		Columns(configs.UserTable+".id", configs.UserTable+".nickname", configs.UserTable+".avatar_path").
		From(configs.SubscriptionTable + ", " + configs.UserTable).
		Suffix(fmt.Sprintf("WHERE %s.creator_id = %s.id and %s.subscriber_id = %d",
			configs.SubscriptionTable, configs.UserTable, configs.SubscriptionTable, id))
}

func CreateSubscriptionsStorage(db *sql.DB) *SubscriptionsStorage {
	return &SubscriptionsStorage{
		db: db,
	}
}

func (storage *SubscriptionsStorage) AddNewSubscription(subscription model.Subscription) (int, error) {
	return 0, nil
}

func (storage *SubscriptionsStorage) DeleteSubscription(id int) error {
	return nil
}

func (storage *SubscriptionsStorage) GetSubscription(id int) (model.Subscription, error) {
	return model.Subscription{}, nil
}

func (storage *SubscriptionsStorage) GetUserSubscriptions(id int) ([]model.User, error) {
	rows, err := SelectUserSubscriptionsSQL(id).RunWith(storage.db).Query()
	if err != nil {
		return []model.User{}, err
	}
	var subsciptions []model.User // probably pointer
	err = dbscan.ScanAll(&subsciptions, rows)
	if err != nil {
		return []model.User{}, err
	}
	return subsciptions, nil
}

func (storage *SubscriptionsStorage) ChangeSubscription(subscription model.Subscription) error {
	return nil
}

func (storage *SubscriptionsStorage) CountSubscribers(id int) (int, error) {
	rows, err := CountSubscribersSQL(id).RunWith(storage.db).Query()
	if err != nil {
		return 0, err
	}
	var counts []*int
	err = dbscan.ScanAll(&counts, rows)
	if err != nil {
		return 0, err
	}
	return *counts[0], nil
}
