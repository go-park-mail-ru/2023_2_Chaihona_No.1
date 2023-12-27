package posts

import "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"

func PostToPostGRPC(post *model.Post) *PostGRPC {
	comments := post.Comments
	tags := post.Tags

	commentsMap := make(map[int32]*CommentGRPC)
	for i, comment := range comments {
		commentsMap[int32(i)] = &CommentGRPC{
			Id: uint32(comment.ID),
			UserId: uint32(comment.UserId),
			PostId: uint32(comment.PostId),
			Text:         comment.Text,
			IsOwner: comment.IsOwner,
			// CreationDate: comment.CreationDate,
		}
	}

	tagsMap := make(map[int32]*TagGRPC)
	for i, tag := range tags {
		tagsMap[int32(i)] = &TagGRPC{
			Id:   uint32(tag.ID),
			Name: tag.Name,
		}
	}

	return &PostGRPC{
		Id:            uint32(post.ID),
		AuthorId:      uint32(post.AuthorID),
		HasAccess:     post.HasAccess,
		Reason:        post.Reason,
		Payment:       float32(post.Payment),
		Currency:      post.Currency,
		MinSubLevel:   uint32(post.MinSubLevel),
		MinSubLevelId: uint32(post.MinSubLevelId),
		CreationDate:  post.CreationDateSQL.Time.Format("2006-01-02 15:04"),
		UpdatedAt:     post.UpdatedAt,
		Header:        post.Header,
		Body:          post.Body,
		Likes:         uint32(post.Likes),
		Comments:      &CommentsGRPC{CommentsMap: commentsMap},
		Tags:          &TagsGRPC{TagsMap: tagsMap},
		Attaches:      post.Attaches,
		IsLiked:       post.IsLiked,
	}
}

func PostGRPCToPost(post *PostGRPC) *model.Post{
	var comments []model.Comment
	var tags []model.Tag

	for _, comment := range post.Comments.CommentsMap {
		comments = append(comments, model.Comment{
			ID: uint(comment.Id),
			UserId: int(comment.UserId),
			PostId: int(comment.PostId),
			Text:         comment.Text,
			IsOwner: comment.IsOwner,
		})
	}

	for _, tag := range post.Tags.TagsMap {
		tags = append(tags, model.Tag{
			ID:   uint(tag.Id),
			Name: tag.Name,
		})
	}

	return &model.Post{
		ID:            uint(post.Id),
		AuthorID:      uint(post.AuthorId),
		HasAccess:     post.HasAccess,
		Reason:        post.Reason,
		Payment:       float64(post.Payment),
		Currency:      post.Currency,
		MinSubLevel:   uint(post.MinSubLevel),
		MinSubLevelId: uint(post.MinSubLevelId),
		CreationDate:  post.CreationDate,
		UpdatedAt:     post.UpdatedAt,
		Header:        post.Header,
		Body:          post.Body,
		Likes:         uint(post.Likes),
		Comments:      comments,
		Tags:          tags,
		Attaches:      post.Attaches,
		IsLiked:       post.IsLiked,
	}
}