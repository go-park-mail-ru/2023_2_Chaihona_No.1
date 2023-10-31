package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"
	postsrep "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/posts"
	profsrep "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/profiles"
	sessrep "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/sessions"
	auth "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/usecases/authorization"
	"github.com/gorilla/mux"
)

type BodyPosts struct {
	Posts []model.Post `json:"posts"`
}

type PostHandler struct {
	Sessions sessrep.SessionRepository
	Posts    postsrep.PostRepository
	Profiles profsrep.ProfileRepository
}

func CreatePostHandlerViaRepos(session sessrep.SessionRepository, posts postsrep.PostRepository,
	profiles profsrep.ProfileRepository,
) *PostHandler {
	return &PostHandler{
		session,
		posts,
		profiles,
	}
}

func (p *PostHandler) GetAllUserPosts(w http.ResponseWriter, r *http.Request) {
	AddAllowHeaders(w)
	w.Header().Add("Content-Type", "application/json")
	if !auth.CheckAuthorization(r, p.Sessions) {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	authorID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, `{"error":"bad id"}`, http.StatusBadRequest)
		return
	}

	cookie, _ := r.Cookie("session_id")
	session, ok := p.Sessions.CheckSession(cookie.Value)
	if !ok {
		http.Error(w, `{"error" : "wrong cookie"}`, http.StatusBadRequest)
	}
	posts, errPost := p.Posts.GetPostsByAuthorId(uint(authorID), uint(session.UserID))

	// сделал по примеру из 6-ой лекции, возможно, стоит добавить обработку по дефолту в свиче
	if errPost != nil {
		switch err.(type) {
		case postsrep.ErrorPost:
			errPost := errPost.(postsrep.ErrorPost)
			if errors.Is(ErrorNotAuthor, errPost.Err) {
				res := Result{Err: errPost.Err.Error()}
				errJSON, err := json.Marshal(res)
				if err != nil {
					errJSON = []byte{}
				}
				http.Error(w, string(errJSON), errPost.StatusCode)
				return
			}
			http.Error(w, `{"error":"db"}`, errPost.StatusCode)
			return
		}
		return
	}

	result := Result{Body: BodyPosts{Posts: posts}}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&result)
	if err != nil {
		log.Println(err)
	}
}

func (p *PostHandler) ChangePost(w http.ResponseWriter, r *http.Request) {
	AddAllowHeaders(w)
	if !auth.CheckAuthorization(r, p.Sessions) {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	_, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, `{"error":"bad id"}`, http.StatusBadRequest)
		return
	}

	body := http.MaxBytesReader(w, r.Body, maxBytesToRead)

	decoder := json.NewDecoder(body)
	post := &model.Post{}

	err = decoder.Decode(post)
	if err != nil {
		http.Error(w, `{"error":"wrong_json"}`, http.StatusBadRequest)
		return
	}

	cookie, _ := r.Cookie("session_id")
	session, ok := p.Sessions.CheckSession(cookie.Value)
	if !ok {
		http.Error(w, `{"error" : "wrong cookie"}`, http.StatusBadRequest)
	}
	post.AuthorID = uint(session.UserID)
	errPost := p.Posts.ChangePost(*post)

	// сделал по примеру из 6-ой лекции, возможно, стоит добавить обработку по дефолту в свиче
	if errPost != nil {
		switch err.(type) {
		case postsrep.ErrorPost:
			errPost := errPost.(postsrep.ErrorPost)
			if errors.Is(ErrorNotAuthor, errPost.Err) {
				res := Result{Err: errPost.Err.Error()}
				errJSON, err := json.Marshal(res)
				if err != nil {
					errJSON = []byte{}
				}
				http.Error(w, string(errJSON), errPost.StatusCode)
				return
			}
			http.Error(w, `{"error":"db"}`, errPost.StatusCode)
			return
		}
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (p *PostHandler) CreateNewPost(w http.ResponseWriter, r *http.Request) {
	AddAllowHeaders(w)
	if !auth.CheckAuthorization(r, p.Sessions) {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}

	body := http.MaxBytesReader(w, r.Body, maxBytesToRead)

	decoder := json.NewDecoder(body)
	post := &model.Post{}

	err := decoder.Decode(post)
	if err != nil {
		http.Error(w, `{"error":"wrong_json"}`, http.StatusBadRequest)
		return
	}

	postId, errPost := p.Posts.CreateNewPost(*post)

	// сделал по примеру из 6-ой лекции, возможно, стоит добавить обработку по дефолту в свиче
	if errPost != nil {
		switch err.(type) {
		case postsrep.ErrorPost:
			errPost := errPost.(postsrep.ErrorPost)
			if errors.Is(ErrorNotAuthor, errPost.Err) {
				res := Result{Err: errPost.Err.Error()}
				errJSON, err := json.Marshal(res)
				if err != nil {
					errJSON = []byte{}
				}
				http.Error(w, string(errJSON), errPost.StatusCode)
				return
			}
			http.Error(w, `{"error":"db"}`, errPost.StatusCode)
			return
		}
		return
	}

	bodyResponse := map[string]interface{}{
		"id": postId,
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&Result{Body: bodyResponse})
	if err != nil {
		log.Println(err)
	}
}

func (p *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	AddAllowHeaders(w)
	if !auth.CheckAuthorization(r, p.Sessions) {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, `{"error":"bad id"}`, 400)
		return
	}

	err = p.Posts.DeletePost(uint(id))
	if err != nil {
		http.Error(w, `{"error":"db"}`, 500)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (p *PostHandler) GetFeed(w http.ResponseWriter, r *http.Request) {
	AddAllowHeaders(w)
	w.Header().Add("Content-Type", "application/json")
	if !auth.CheckAuthorization(r, p.Sessions) {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}

	cookie, _ := r.Cookie("session_id")
	session, ok := p.Sessions.CheckSession(cookie.Value)
	if !ok {
		http.Error(w, `{"error" : "wrong cookie"}`, http.StatusBadRequest)
	}
	posts, errPost := p.Posts.GetUsersFeed(uint(session.UserID))

	// сделал по примеру из 6-ой лекции, возможно, стоит добавить обработку по дефолту в свиче
	if errPost != nil {
		switch errPost.(type) {
		case postsrep.ErrorPost:
			errPost := errPost.(postsrep.ErrorPost)
			if errors.Is(ErrorNotAuthor, errPost.Err) {
				res := Result{Err: errPost.Err.Error()}
				errJSON, err := json.Marshal(res)
				if err != nil {
					errJSON = []byte{}
				}
				http.Error(w, string(errJSON), errPost.StatusCode)
				return
			}
			http.Error(w, `{"error":"db"}`, errPost.StatusCode)
			return
		}
		return
	}

	result := Result{Body: BodyPosts{Posts: posts}}

	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(&result)
	if err != nil {
		log.Println(err)
	}
}
