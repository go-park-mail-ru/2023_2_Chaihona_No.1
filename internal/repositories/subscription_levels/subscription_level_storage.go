package subscriptionlevels

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/dbscan"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/configs"
	model "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"
)

func SelectUserLevelsSQL(userId uint) squirrel.SelectBuilder {
	return squirrel.Select("*").
		From(configs.SubscribeLevelTable).
		Where(squirrel.Eq{"creator_id": userId}).
		PlaceholderFormat(squirrel.Dollar).
		OrderBy("cost_integer DESC", "cost_fractional DESC")
}

type SubscribeLevelStorage struct {
	db *sql.DB
}

func CreateSubscribeLevelStorage(db *sql.DB) *SubscribeLevelStorage {
	return &SubscribeLevelStorage{
		db: db,
	}
}

func (storage *SubscribeLevelStorage) AddNewLevel(level *model.SubscribeLevel) (model.SubscribeLevel, error) {
	return model.SubscribeLevel{}, nil
}

func (storage *SubscribeLevelStorage) DeleteLevel(id uint) error {
	return nil
}

func (storage *SubscribeLevelStorage) GetUserLevels(userId uint) ([]model.SubscribeLevel, error) {
	rows, err := SelectUserLevelsSQL(userId).RunWith(storage.db).Query()
	if err != nil {
		return make([]model.SubscribeLevel, 0), err
	}
	var levels []model.SubscribeLevel
	err = dbscan.ScanAll(&levels, rows)
	if err != nil {
		return make([]model.SubscribeLevel, 0), err
	}
	return levels, nil
}

func (storage *SubscribeLevelStorage) GetLevel(id uint) (model.SubscribeLevel, error) {
	return model.SubscribeLevel{}, nil
}
