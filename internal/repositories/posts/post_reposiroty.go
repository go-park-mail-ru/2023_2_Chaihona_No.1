package posts

import (
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"
)

type PostRepository interface {
	CreateNewPost(post model.Post) (int, error)
	DeletePost(id uint) error
	ChangePost(post model.Post) error
	GetPostById(id uint) (model.Post, error)
	GetPostsByAuthorId(authorID uint, subscriberId uint) ([]model.Post, error)
	GetUsersFeed(userId uint) ([]model.Post, error)
}
