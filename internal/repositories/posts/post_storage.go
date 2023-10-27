package posts

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/dbscan"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/configs"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"
)

func InsertPostSQL(post model.Post) squirrel.InsertBuilder {
	return squirrel.Insert(configs.PostTable).
		Columns("header", "body", "creator_id", "min_subscription_level_id").
		Values(post.Header, post.Body, post.AuthorID, post.MinSubLevelId).
		Suffix("RETURNING \"id\"").
		PlaceholderFormat(squirrel.Dollar)
}

func DeletePostSQL(postId uint) squirrel.DeleteBuilder {
	return squirrel.Delete(configs.PostTable).
		Where(squirrel.Eq{"id": postId}).
		PlaceholderFormat(squirrel.Dollar)
}

func SelectUserPostsSQL(authorId uint) squirrel.SelectBuilder {
	return squirrel.Select("*").
		From(configs.PostTable).
		Where(squirrel.Eq{"creator_id": authorId}).
		PlaceholderFormat(squirrel.Dollar)
}

func SelectPostByIdSQL(postId uint) squirrel.SelectBuilder {
	return squirrel.Select("*").
		From(configs.PostTable).
		Where(squirrel.Eq{"id": postId}).
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

func (storage *PostStorage) CreateNewPost(post model.Post) (int, error) {
	var postId int
	err := InsertPostSQL(post).RunWith(storage.db).QueryRow().Scan(postId)
	if err != nil {
		return 0, err
	}
	return postId, nil
}

func (storage *PostStorage) DeletePost(id uint) error {
	_, err := DeletePostSQL(id).RunWith(storage.db).Query()
	if err != nil {
		return err
	}
	return nil
}

func (storage *PostStorage) GetPostById(postId uint) (model.Post, error) {
	rows, err := SelectPostByIdSQL(postId).RunWith(storage.db).Query()
	if err != nil {
		return model.Post{}, err
	}
	var posts []model.Post
	if err = dbscan.ScanAll(&posts, rows); err != nil {
		return model.Post{}, err
	}
	if len(posts) > 0 {
		return posts[0], nil
	}
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
