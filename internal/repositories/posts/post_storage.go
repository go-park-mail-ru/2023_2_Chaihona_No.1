package posts

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/dbscan"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/configs"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"
)

func SelectUserPostsSQL(authorId uint) squirrel.SelectBuilder {
	return squirrel.Select("*").
		From(configs.PostTable).
		Where(squirrel.Eq{"creator_id": authorId}).
		PlaceholderFormat(squirrel.Dollar)
}

type PostStorage struct {
	db *sql.DB
}

func CreatePostStorage(db *sql.DB) PostRepository {
	return &PostStorage{
		db: db,
	}
}

func (storage *PostStorage) CreateNewPost(post model.Post) error {
	return nil
}

func (storage *PostStorage) DeletePost(id uint) error {
	return nil
}

func (storage *PostStorage) GetPostById(id uint) (model.Post, error) {
	return model.Post{}, nil
}

func (storage *PostStorage) GetPostsByAuthorId(authorId uint) ([]model.Post, error) {
	rows, err := SelectUserPostsSQL(authorId).RunWith(storage.db).Query()
	if err != nil {
		return []model.Post{}, err
	}
	var posts []model.Post
	err = dbscan.ScanAll(&posts, rows)
	if err != nil {
		return []model.Post{}, err
	}
	return posts, nil
}

func (storage *PostStorage) GetPosts() ([]model.Post, error) {
	return []model.Post{}, nil
}
