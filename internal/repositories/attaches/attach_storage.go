package attaches

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/dbscan"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/configs"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"
)

//add metadata
func InsertAttachSQL(attach model.Attach) squirrel.InsertBuilder {
	return squirrel.Insert(configs.AttachTable).
		Columns("file_path", "post_id", "name", "is_media").
		Values(attach.FilePath, attach.PostId, attach.Name, attach.IsMedia).
		Suffix("RETURNING \"id\"").
		PlaceholderFormat(squirrel.Dollar)
}

func SelectPostAttachesSQL(postID int) squirrel.SelectBuilder {
	return squirrel.Select("file_path", "name", "is_media").
		From(configs.AttachTable + " a").
		Where(squirrel.Eq{"a.post_id": postID}).
		PlaceholderFormat(squirrel.Dollar)
}

func CountPostAttachesSQL(postID int) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(*) as attaches").
		From(configs.AttachTable + " a").
		Where(squirrel.Eq{"a.post_id": postID}).
		PlaceholderFormat(squirrel.Dollar)
}
	
func DeleteAttachSQL(path string) squirrel.DeleteBuilder {
	return squirrel.Delete(configs.AttachTable).
		Where(squirrel.Eq{"file_path": path}).
		PlaceholderFormat(squirrel.Dollar)
}

type AttachStorage struct {
	db *sql.DB
}

func CreateAttachStorage(db *sql.DB) AttachRepository {
	return &AttachStorage{
		db: db,
	}
}


func (storage *AttachStorage) PinAttach(attach model.Attach) (int, error) {
	var attachId int
	err := InsertAttachSQL(attach).RunWith(storage.db).QueryRow().Scan(&attachId)
	if err != nil {
		return 0, err
	}
	return attachId, nil
}

func (storage *AttachStorage) GetPostAttaches(postID int) ([]model.Attach, error) {
	rows, err := SelectPostAttachesSQL(postID).RunWith(storage.db).Query()
	if err != nil {
		return []model.Attach{}, err
	}
	var attaches []model.Attach
	err = dbscan.ScanAll(&attaches, rows)
	if err != nil {
		return []model.Attach{}, err
	}
	return attaches, nil
}


func (storage *AttachStorage) DeleteAttach(path string) error {
	rows, err := DeleteAttachSQL(path).RunWith(storage.db).Query()
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}

func (storage *AttachStorage) CountAttaches(postId int) (int, error) {
	var countAttaches int
	err := CountPostAttachesSQL(postId).RunWith(storage.db).QueryRow().Scan(&countAttaches)
	if err != nil {
		return 0, err
	}
	return countAttaches, nil
}