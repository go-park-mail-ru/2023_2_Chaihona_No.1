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

func SelectAllNotFreeSubscriptionsSQL() squirrel.SelectBuilder {
	return squirrel.Select("s.*").
		From(configs.SubscriptionTable +" s").
		InnerJoin(configs.SubscribeLevelTable + " sl ON s.subscription_level_id = sl.id").
		Where("sl.level <> 0").
		PlaceholderFormat(squirrel.Dollar)
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

func InsertSubscriptionSQL(subsciption model.Subscription) squirrel.InsertBuilder {
	return squirrel.Insert(configs.SubscriptionTable).
		Columns("creator_id", "subscriber_id", "subscription_level_id").
		Values(subsciption.Creator_id, subsciption.Subscriber_id, subsciption.Subscription_level_id).
		Suffix("RETURNING \"id\"").
		PlaceholderFormat(squirrel.Dollar)
}

func InsertOrUpdateSubscriptionSQL(subsciption model.Subscription) squirrel.InsertBuilder {
	return squirrel.Insert(configs.SubscriptionTable).
		Columns("id", "creator_id", "subscriber_id", "subscription_level_id").
		Values(subsciption.Id, subsciption.Creator_id, subsciption.Subscriber_id, subsciption.Subscription_level_id).
		Suffix(fmt.Sprintf(" ON CONFLICT(id) DO UPDATE SET creator_id=%d, subscriber_id=%d, subscription_level_id=%d", 
												subsciption.Creator_id, subsciption.Subscriber_id, subsciption.Subscription_level_id)).
		// Suffix("RETURNING \"id\"").
		PlaceholderFormat(squirrel.Dollar)
}

func DeleteSubscriptionSQL(levelId, subId int) squirrel.DeleteBuilder {
	return squirrel.Delete(configs.SubscriptionTable).
		Where(squirrel.Eq{"subscription_level_id": levelId, "subscriber_id":subId}).
		PlaceholderFormat(squirrel.Dollar)
}

func CreateSubscriptionsStorage(db *sql.DB) *SubscriptionsStorage {
	return &SubscriptionsStorage{
		db: db,
	}
}

func (storage *SubscriptionsStorage) AddNewSubscription(subscription model.Subscription) (int, error) {
	var subId int
	var err error
	if (subscription.Id != 0) {
		rows, err := InsertOrUpdateSubscriptionSQL(subscription).RunWith(storage.db).Query()
		defer rows.Close()
		if err != nil {
			return 0, err
		}
		subId = int(subscription.Id)
	} else {
		err = InsertSubscriptionSQL(subscription).RunWith(storage.db).QueryRow().Scan(&subId)
	}
	if err != nil {
		return 0, err
	}
	return subId, nil
}

func (storage *SubscriptionsStorage) DeleteSubscription(levelId, subId int) error {
	rows, err := DeleteSubscriptionSQL(levelId, subId).RunWith(storage.db).Query()
	defer rows.Close()
	if err != nil {
		return err
	}
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

func (storage *SubscriptionsStorage) GetAllNotFreeSubscriptions() ([]model.Subscription, error) {
	rows, err := SelectAllNotFreeSubscriptionsSQL().RunWith(storage.db).Query()
	if err != nil {
		return []model.Subscription{},err
	}
	var subscriptions []model.Subscription
	err = dbscan.ScanAll(&subscriptions, rows)
	if err != nil {
		return []model.Subscription{},err
	}
	return subscriptions, nil
}