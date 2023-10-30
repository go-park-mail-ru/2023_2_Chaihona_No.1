package likes

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/configs"
)

func InsertLikeSQL(userId int, postId int) squirrel.InsertBuilder {
	return squirrel.Insert(configs.LikeTable).
		Columns("user_id", "post_id").
		Values(userId, postId).
		PlaceholderFormat(squirrel.Dollar)
}

func DeleteLikeSQL(userId int, postId int) squirrel.DeleteBuilder {
	return squirrel.Delete(configs.LikeTable).
		Where(squirrel.Eq{"user_id": userId, "post_id": postId}).
		PlaceholderFormat(squirrel.Dollar)
}

type LikeStorage struct {
	db *sql.DB
}

func CreateLikeStorage(db *sql.DB) *LikeStorage {
	return &LikeStorage{
		db: db,
	}
}

func (storage *LikeStorage) CreateNewLike(userId int, postId int) error {
	_, err := InsertLikeSQL(userId, postId).RunWith(storage.db).Query()
	if err != nil {
		return err
	}
	return nil
}

func (storage *LikeStorage) DeleteLike(userId int, postId int) error {
	_, err := DeleteLikeSQL(userId, postId).RunWith(storage.db).Query()
	if err != nil {
		return err
	}
	return nil
}
