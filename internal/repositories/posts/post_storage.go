package posts

import (
	"net/http"
	"sync"

	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"
)

type PostStorage struct {
	posts map[uint]model.Post
	mu    sync.Mutex
}

func CreatePostStorage() PostRepository {
	storage := &PostStorage{
		posts: make(map[uint]model.Post),
	}

	return storage
}

func (storage *PostStorage) CreateNewPost(post model.Post) error {
	storage.mu.Lock()
	defer storage.mu.Unlock()

	storage.posts[post.ID] = post

	return nil
}

func (storage *PostStorage) DeletePost(id uint) error {
	storage.mu.Lock()
	defer storage.mu.Unlock()

	if _, ok := storage.posts[id]; !ok {
		return &ErrorPost{ErrNoSuchPost, http.StatusBadRequest}
	}

	delete(storage.posts, id)

	return nil
}

func (storage *PostStorage) GetPostById(id uint) (model.Post, error) {
	storage.mu.Lock()
	defer storage.mu.Unlock()

	val, ok := storage.posts[id]

	if !ok {
		return model.Post{}, &ErrorPost{ErrNoSuchPost, http.StatusBadRequest}
	}

	return val, nil
}

func (storage *PostStorage) GetPostsByAuthorId(authorId uint) ([]model.Post, error) {
	storage.mu.Lock()
	defer storage.mu.Unlock()

	posts := make([]model.Post, 0)
	for _, post := range storage.posts {
		if post.AuthorID == authorId {
			posts = append(posts, post)
		}
	}

	return posts, nil
}

func (storage *PostStorage) GetPosts() ([]model.Post, error) {
	storage.mu.Lock()
	defer storage.mu.Unlock()

	posts := make([]model.Post, 0)

	for _, post := range storage.posts {
		posts = append(posts, post)
	}

	return posts, nil
}
