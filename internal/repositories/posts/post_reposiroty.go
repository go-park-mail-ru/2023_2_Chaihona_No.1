package posts

import (
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"
)

type PostRepository interface {
	CreateNewPost(post model.Post) (int, error)
	DeletePost(id uint) error
	GetPostById(id uint) (model.Post, error)
	GetPostsByAuthorId(authorID uint) ([]model.Post, error)
}
