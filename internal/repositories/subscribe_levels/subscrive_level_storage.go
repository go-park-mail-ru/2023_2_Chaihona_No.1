package subscribelevels

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
		PlaceholderFormat(squirrel.Dollar)
}

func InsertLevelSQL(level model.SubscribeLevel) squirrel.InsertBuilder {
	return squirrel.Insert(configs.SubscribeLevelTable).
		Columns("level", "name", "description", "cost_integer", "cost_fractional", "currency", "creator_id").
		Values(level.Level, level.Name, level.Description, level.CostInteger, level.CostFractional, level.Currency, level.CreatorID).
		Suffix("RETURNING \"id\"").
		PlaceholderFormat(squirrel.Dollar)
}

type SubscribeLevelStorage struct {
	db *sql.DB
}

func CreateSubscribeLevelStorage(db *sql.DB) *SubscribeLevelStorage {
	return &SubscribeLevelStorage{
		db: db,
	}
}

func (storage *SubscribeLevelStorage) AddNewLevel(level model.SubscribeLevel) (int, error) {
	var postId int
	err := InsertLevelSQL(level).RunWith(storage.db).QueryRow().Scan(&postId)
	if err != nil {
		return 0, err
	}
	return postId, nil
}

func (storage *SubscribeLevelStorage) DeleteLevel(id uint) error {
	return nil
}

func (storage *SubscribeLevelStorage) GetUserLevels(userId uint) ([]model.SubscribeLevel, error) {
	rows, err := SelectUserLevelsSQL(userId).RunWith(storage.db).Query()
	if err != nil {
		return make([]model.SubscribeLevel, 0), err
	}
	var levels []model.SubscribeLevel // probably array of pointers
	err = dbscan.ScanAll(&levels, rows)
	if err != nil {
		return make([]model.SubscribeLevel, 0), err
	}
	return levels, nil
}

func (storage *SubscribeLevelStorage) GetLevel(id uint) (model.SubscribeLevel, error) {
	return model.SubscribeLevel{}, nil
}