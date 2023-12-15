package comments

import (
	context "context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/configs"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"
)

func InsertCommentSQL(comment model.Comment) squirrel.InsertBuilder {
	return squirrel.Insert(configs.CommentTable).
		Columns("user_id", "post_id", "text").
		Values(comment.UserId, comment.PostId, comment.Text).
		Suffix("RETURNING \"id\"").
		PlaceholderFormat(squirrel.Dollar)
}

func DeleteCommentSQL(commentId uint) squirrel.DeleteBuilder {
	return squirrel.Delete(configs.CommentTable).
		Where(squirrel.Eq{"id": commentId}).
		PlaceholderFormat(squirrel.Dollar)
}

func UpdateCommentSQL(comment model.Comment) squirrel.UpdateBuilder {
	return squirrel.Update(configs.PostTable).
		SetMap(map[string]interface{}{
			"text": comment.Text,
		}).
		Where(squirrel.Eq{"post_id": comment.PostId}).
		PlaceholderFormat(squirrel.Dollar)
}

type CommentStorage struct {
	db *sql.DB
	UnimplementedCommentServiceServer
	// posts.UnimplementedPostsServiceServer
}

type CommentManager struct {
	CLient CommentServiceClient
	// CLient posts.PostsServiceClient
}

func CreateCommentStore(db *sql.DB) *CommentStorage {
	return &CommentStorage{
		db: db,
	}
}

func (manager *CommentManager) CreateComment(comment model.Comment) (int, error) {
	id, err := manager.CLient.CreateCommentCtx(context.Background(), CommentToCommentGRPC(&comment))
	fmt.Println(err)
	if err != nil {
		return 0, err
	}
	return int(id.I), err
}

func (storage *CommentStorage) CreateCommentCtx(ctx context.Context, comment *CommentGRPC) (*Int, error) {
	var commentId int

	err := InsertCommentSQL(*CommentGRPCToComment(comment)).RunWith(storage.db).QueryRow().Scan(&commentId)
	if err != nil {
		return &Int{I: 0}, err
	}

	return &Int{I: int32(commentId)}, nil
}

func (manager *CommentManager) DeleteComment(id uint) error {
	_, err := manager.CLient.DeleteCommentCtx(context.Background(), &UInt{Id: uint32(id)})
	return err
}

func (storage *CommentStorage) DeleteCommentCtx(ctx context.Context, id *UInt) (*Nothing, error) {
	rows, err := DeleteCommentSQL(uint(id.Id)).RunWith(storage.db).Query()
	if err != nil {
		return &Nothing{}, err
	}
	defer rows.Close()
	return &Nothing{}, nil
}

func (manager *CommentManager) ChangeComment(comment model.Comment) error {
	_, err := manager.CLient.ChangeCommentCtx(context.Background(), CommentToCommentGRPC(&comment))

	return err
}

func (storage *CommentStorage) ChangeCommentCtx(ctx context.Context, post *CommentGRPC) (*Nothing, error) {
	rows, err := UpdateCommentSQL(*CommentGRPCToComment(post)).RunWith(storage.db).Query()
	if err != nil {
		return &Nothing{}, err
	}
	defer rows.Close()
	return &Nothing{}, nil
}