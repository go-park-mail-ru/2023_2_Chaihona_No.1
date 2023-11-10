package posts

import (
	"database/sql"
	"fmt"

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

// func SelectUserPostsSQL(authorId uint) squirrel.SelectBuilder {
// 	return squirrel.Select("*").
// 		From(configs.PostTable).
// 		Where(squirrel.Eq{"creator_id": authorId}).
// 		PlaceholderFormat(squirrel.Dollar)
// }

func SelectPostByIdSQL(postId uint) squirrel.SelectBuilder {
	return squirrel.Select("*").
		From(configs.PostTable).
		Where(squirrel.Eq{"id": postId}).
		PlaceholderFormat(squirrel.Dollar)
}

// For followed users
// func SelectUserPostsSQL(authorId uint, subscriberId uint) squirrel.SelectBuilder {
// 	return squirrel.Select("p.*, CASE WHEN sl1.level > sl2.level THEN FALSE ELSE TRUE END AS has_access, "+
// 		"sl1.level as min_sub_level, "+
// 		"array_agg(pa.file_path) as attaches, "+
// 		"coalesce(array_length(array_agg(distinct pl.id) FILTER (WHERE pl IS NOT NULL), 1), 0) as likes, "+
// 		fmt.Sprintf("CASE WHEN coalesce(array_length(array_agg(distinct pl.id) FILTER (WHERE pl.user_id = %d), 1), 0) > 0 THEN TRUE ELSE FALSE END AS is_liked", subscriberId)).
// 		From(configs.PostTable+" p").
// 		CrossJoin(configs.SubscriptionTable+" s").
// 		LeftJoin(configs.AttachTable+" pa ON p.id = pa.post_id").
// 		LeftJoin(configs.LikeTable+" pl ON p.id = pl.post_id").
// 		InnerJoin(configs.SubscribeLevelTable+" sl1 ON p.min_subscription_level_id = sl1.id").
// 		InnerJoin(configs.SubscribeLevelTable+" sl2 ON s.subscription_level_id = sl2.id").
// 		Where(squirrel.Eq{
// 			"p.creator_id":    authorId,
// 			"s.subscriber_id": subscriberId,
// 			"s.creator_id":    authorId,
// 		}).
// 		GroupBy("p.id", "sl1.level", "sl2.level").
// 		PlaceholderFormat(squirrel.Dollar)
// }

//For no followed users
func SelectUserPostsSQL(authorId uint, subscriberId uint) squirrel.SelectBuilder {
	return squirrel.Select("p.*, TRUE AS has_access, "+ //change on CASE WHEN sl1.level = 0 THEN TRUE ELSE FALSE END AS has_access, "+
		"sl1.level as min_sub_level, "+
		"array_agg(pa.file_path) as attaches, "+
		"coalesce(array_length(array_agg(distinct pl.id) FILTER (WHERE pl IS NOT NULL), 1), 0) as likes, "+
		fmt.Sprintf("CASE WHEN coalesce(array_length(array_agg(distinct pl.id) FILTER (WHERE pl.user_id = %d), 1), 0) > 0 THEN TRUE ELSE FALSE END AS is_liked", subscriberId)).
		From(configs.PostTable+" p").
		LeftJoin(configs.AttachTable+" pa ON p.id = pa.post_id").
		LeftJoin(configs.LikeTable+" pl ON p.id = pl.post_id").
		InnerJoin(configs.SubscribeLevelTable+" sl1 ON p.min_subscription_level_id = sl1.id").
		Where(squirrel.Eq{
			"p.creator_id":    authorId,
		}).
		GroupBy("p.id", "sl1.level").
		PlaceholderFormat(squirrel.Dollar)
}

//For owner 
// func SelectUserPostsSQL(authorId uint, subscriberId uint) squirrel.SelectBuilder {
// 	return squirrel.Select("p.*, TRUE AS has_access, "+
// 		"sl1.level as min_sub_level, "+
// 		"array_agg(pa.file_path) as attaches, "+
// 		"coalesce(array_length(array_agg(distinct pl.id) FILTER (WHERE pl IS NOT NULL), 1), 0) as likes, "+
// 		fmt.Sprintf("CASE WHEN coalesce(array_length(array_agg(distinct pl.id) FILTER (WHERE pl.user_id = %d), 1), 0) > 0 THEN TRUE ELSE FALSE END AS is_liked", subscriberId)).
// 		From(configs.PostTable+" p").
// 		LeftJoin(configs.AttachTable+" pa ON p.id = pa.post_id").
// 		LeftJoin(configs.LikeTable+" pl ON p.id = pl.post_id").
// 		InnerJoin(configs.SubscribeLevelTable+" sl1 ON p.min_subscription_level_id = sl1.id").
// 		Where(squirrel.Eq{
// 			"p.creator_id":    authorId,
// 		}).
// 		GroupBy("p.id", "sl1.level").
// 		PlaceholderFormat(squirrel.Dollar)
// }

// func SelectAvailiblePostsSQL(userId uint) squirrel.SelectBuilder {
// 	return squirrel.Select("p.*, CASE WHEN sl1.level > sl2.level THEN FALSE ELSE TRUE END AS has_access, "+
// 		"array_agg(pa.file_path) as attaches, "+
// 		"coalesce(array_length(array_agg(distinct pl.id) FILTER (WHERE pl IS NOT NULL), 1), 0) as likes, "+
// 		fmt.Sprintf("CASE WHEN coalesce(array_length(array_agg(distinct pl.id) FILTER (WHERE pl.user_id = %d), 1), 0) > 0 THEN TRUE ELSE FALSE END AS is_liked", userId)).
// 		From(configs.PostTable+" p").
// 		InnerJoin(configs.SubscriptionTable+" s ON p.creator_id = s.creator_id").
// 		LeftJoin(configs.AttachTable+" pa ON p.id = pa.post_id").
// 		LeftJoin(configs.LikeTable+" pl ON p.id = pl.post_id").
// 		InnerJoin(configs.SubscribeLevelTable+" sl1 ON p.min_subscription_level_id = sl1.id").
// 		InnerJoin(configs.SubscribeLevelTable+" sl2 ON s.subscription_level_id = sl2.id").
// 		Where(squirrel.Eq{
// 			"s.subscriber_id": userId,
// 		}, "sl1.level <= sl2.level").
// 		GroupBy("p.id", "sl1.level", "sl2.level").
// 		PlaceholderFormat(squirrel.Dollar)
// }


func SelectAvailiblePostsSQL(userId uint) squirrel.SelectBuilder {
	return squirrel.Select("p.*, TRUE AS has_access, "+
		"array_agg(pa.file_path) as attaches, "+
		"coalesce(array_length(array_agg(distinct pl.id) FILTER (WHERE pl IS NOT NULL), 1), 0) as likes, "+
		fmt.Sprintf("CASE WHEN coalesce(array_length(array_agg(distinct pl.id) FILTER (WHERE pl.user_id = %d), 1), 0) > 0 THEN TRUE ELSE FALSE END AS is_liked", userId)).
		From(configs.PostTable+" p").
		InnerJoin(configs.SubscriptionTable+" s ON p.creator_id = s.creator_id").
		LeftJoin(configs.AttachTable+" pa ON p.id = pa.post_id").
		LeftJoin(configs.LikeTable+" pl ON p.id = pl.post_id").
		Where(squirrel.Eq{
			"s.subscriber_id": userId,
		}).
		GroupBy("p.id").
		PlaceholderFormat(squirrel.Dollar)
}

func UpdatePostSQL(post model.Post) squirrel.UpdateBuilder {
	return squirrel.Update(configs.PostTable).
		SetMap(map[string]interface{}{
			"header":                    post.Header,
			"body":                      post.Body,
			"min_subscription_level_id": post.MinSubLevelId,
		}).
		Where(squirrel.Eq{"id": post.ID, "creator_id": post.AuthorID}).
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
	err := InsertPostSQL(post).RunWith(storage.db).QueryRow().Scan(&postId)
	if err != nil {
		return 0, err
	}
	return postId, nil
}

func (storage *PostStorage) DeletePost(id uint) error {
	rows, err := DeletePostSQL(id).RunWith(storage.db).Query()
	defer rows.Close()
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

func (storage *PostStorage) ChangePost(post model.Post) error {
	rows, err := UpdatePostSQL(post).RunWith(storage.db).Query()
	defer rows.Close()
	if err != nil {
		return err
	}
	return nil
}

func (storage *PostStorage) GetPostsByAuthorId(authorId uint, subscriberId uint) ([]model.Post, error) {
	rows, err := SelectUserPostsSQL(authorId, subscriberId).RunWith(storage.db).Query()
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

func (storage *PostStorage) GetUsersFeed(userId uint) ([]model.Post, error) {
	rows, err := SelectAvailiblePostsSQL(userId).RunWith(storage.db).Query()
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