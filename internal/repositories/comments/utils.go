package comments

import "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"

func CommentToCommentGRPC(comment *model.Comment) *CommentGRPC {
	return &CommentGRPC{
		Id:            uint32(comment.ID),
		PostId: uint32(comment.PostId),
		UserId: uint32(comment.UserId),
		Text: comment.Text,
	}
}

func CommentGRPCToComment(comment *CommentGRPC) *model.Comment{
	return &model.Comment{
		ID:            uint(comment.Id),
		UserId: int(comment.UserId),
		PostId: int(comment.PostId),
		Text: comment.Text,
	}
}