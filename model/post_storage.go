package model

import (
	"errors"
	"sync"
)

var ErrNoSuchPost = errors.New("no such post")

type PostRepository interface {
	CreateNewPost(post Post) error
	DeletePost(id uint) error
	GetPostById(id uint) (*Post, error)
	GetPostsByAuthorId(authorId uint) (*[]Post, error)
	GetPosts() ([]Post, error)
}

type PostStorage struct {
	Posts map[uint]Post
	Mu    sync.Mutex
	Size  uint32
}

func CreatePostStorage() PostRepository {
	storage := &PostStorage{
		Posts: make(map[uint]Post),
	}

	return storage
}

func (storage *PostStorage) CreateNewPost(post Post) error {
	storage.Mu.Lock()
	defer storage.Mu.Unlock()

	storage.Size++
	storage.Posts[post.ID] = post

	return nil
}

func (storage *PostStorage) DeletePost(id uint) error {
	storage.Mu.Lock()
	defer storage.Mu.Unlock()

	if _, ok := storage.Posts[id]; !ok {
		return ErrNoSuchPost
	}

	delete(storage.Posts, id)
	storage.Size--

	return nil
}

func (storage *PostStorage) GetPostById(id uint) (*Post, error) {
	storage.Mu.Lock()
	defer storage.Mu.Unlock()

	val, ok := storage.Posts[id]

	if ok {
		copy := val

		return &copy, nil
	}

	return nil, ErrNoSuchPost
}

func (storage *PostStorage) GetPostsByAuthorId(authorId uint) (*[]Post, error) {
	storage.Mu.Lock()
	defer storage.Mu.Unlock()

	posts := make([]Post, 0)
	for _, post := range storage.Posts {
		if post.AuthorID == authorId {
			posts = append(posts, post)
		}
	}

	return &posts, nil
}

func (storage *PostStorage) GetPosts() ([]Post, error) {
	storage.Mu.Lock()
	defer storage.Mu.Unlock()

	posts := make([]Post, 0)

	for _, post := range storage.Posts {
		posts = append(posts, post)
	}

	return posts, nil
}
