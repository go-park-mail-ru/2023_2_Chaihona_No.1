package model

import (
	"errors"
	"net/http"
	"sync"
)

var ErrNoSuchPost = errors.New("no such post")

type ErrorPost struct {
	Err        error `json:"error"`
	StatusCode int
}

type PostRepository interface {
	CreateNewPost(post Post) *ErrorPost
	DeletePost(id uint) *ErrorPost
	GetPostById(id uint) (Post, *ErrorPost)
	GetPostsByAuthorId(authorId uint) ([]Post, *ErrorPost)
	GetPosts() ([]Post, *ErrorPost)
}

type PostStorage struct {
	posts map[uint]Post
	mu    sync.Mutex
}

func CreatePostStorage() PostRepository {
	storage := &PostStorage{
		posts: make(map[uint]Post),
	}

	return storage
}

func (storage *PostStorage) CreateNewPost(post Post) *ErrorPost {
	storage.mu.Lock()
	defer storage.mu.Unlock()

	storage.posts[post.ID] = post

	return nil
}

func (storage *PostStorage) DeletePost(id uint) *ErrorPost {
	storage.mu.Lock()
	defer storage.mu.Unlock()

	if _, ok := storage.posts[id]; !ok {
		return &ErrorPost{ErrNoSuchPost, http.StatusBadRequest}
	}

	delete(storage.posts, id)

	return nil
}

func (storage *PostStorage) GetPostById(id uint) (Post, *ErrorPost) {
	storage.mu.Lock()
	defer storage.mu.Unlock()

	val, ok := storage.posts[id]

	if !ok {
		return Post{}, &ErrorPost{ErrNoSuchPost, http.StatusBadRequest}
	}

	return val, nil
}

func (storage *PostStorage) GetPostsByAuthorId(authorId uint) ([]Post, *ErrorPost) {
	storage.mu.Lock()
	defer storage.mu.Unlock()

	posts := make([]Post, 0)
	for _, post := range storage.posts {
		if post.AuthorID == authorId {
			posts = append(posts, post)
		}
	}

	return posts, nil
}

func (storage *PostStorage) GetPosts() ([]Post, *ErrorPost) {
	storage.mu.Lock()
	defer storage.mu.Unlock()

	posts := make([]Post, 0)

	for _, post := range storage.posts {
		posts = append(posts, post)
	}

	return posts, nil
}
