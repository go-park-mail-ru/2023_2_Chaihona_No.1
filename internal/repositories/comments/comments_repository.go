package comments

import (
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"
)

type CommentRepository interface {
	CreateComment(comment model.Comment) (int, error)
	// GetPostComments(postID int) ([]model.Comment, error)
	DeleteComment(commentId int) error
	EditComment(comment model.Comment) (error)
}