package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	auth "project/authorization"
	model "project/model"
	"strconv"

	"github.com/gorilla/mux"
)

type PostHandler struct {
	Sessions auth.SessionRepository
	Posts    model.PostRepository
}

func CreatePostHandler() *PostHandler {
	return &PostHandler{
		auth.CreateSessionStorage(),
		model.CreatePostStorage(),
	}
}

func (p *PostHandler) GetAllUserPosts(w http.ResponseWriter, r *http.Request) {
	if auth.CheckAuthorization(w, r, p.Sessions) {
		http.Error(w, `{"error":"unauthorized"}`, 401)
		return
	}

	vars := mux.Vars(r)
	authorId, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, `{"error":"bad id"}`, 400)
		return
	}

	posts, err := p.Posts.GetPostsByAuthorId(uint(authorId))
	if err != nil {
		if err == errors.New("user isn't author") {
			http.Error(w, `{"error":"bad request"}`, 400)
		}
		http.Error(w, `{"error":"db"}`, 500)
		return
	}

	body := map[string]interface{}{
		"posts": posts,
	}
	json.NewEncoder(w).Encode(&Result{Body: body})
}
