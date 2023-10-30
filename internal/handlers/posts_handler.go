package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"
	likesrep "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/likes"
	postsrep "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/posts"
	profsrep "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/profiles"
	sessrep "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/sessions"
	auth "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/usecases/authorization"
	"github.com/gorilla/mux"
)

type BodyPosts struct {
	Posts []model.Post `json:"posts"`
}

type BodyLike struct {
	PostId int `json:"post_id"`
}

type PostHandler struct {
	Sessions sessrep.SessionRepository
	Posts    postsrep.PostRepository
	Profiles profsrep.ProfileRepository
	Likes    likesrep.LikeRepository
}

func CreatePostHandlerViaRepos(session sessrep.SessionRepository, posts postsrep.PostRepository,
	profiles profsrep.ProfileRepository, likes likesrep.LikeRepository,
) *PostHandler {
	return &PostHandler{
		session,
		posts,
		profiles,
		likes,
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

	posts, errPost := p.Posts.GetPostsByAuthorId(uint(authorID))

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
	}

	cookie, _ := r.Cookie("session_id")
	session, _ := p.Sessions.CheckSession(cookie.Value)
	userID := session.UserID
	profile, ok := p.Profiles.GetProfile(uint(userID))
	if !ok {
		http.Error(w, `{"error":"very bad"}`, 400)
		return
	}
	subscriptions := profile.Subscriptions
	isSubscirber := false
	for _, user := range subscriptions {
		if user.ID == uint(authorID) {
			isSubscirber = true
		}
	}

	if !isSubscirber {
		for i := range posts {
			switch posts[i].Access {
			case model.SubscribersAccess:
				posts[i].HasAccess = false
				posts[i].Reason = model.LowLevelReason
				posts[i].Body = ""
			case model.EveryoneAccess:
				posts[i].HasAccess = true
			}
		}
	} else {
		for i := range posts {
			posts[i].HasAccess = true
		}
	}

	result := Result{Body: BodyPosts{Posts: posts}}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&result)
	if err != nil {
		log.Println(err)
	}
}

func (p *PostHandler) LikePost(w http.ResponseWriter, r *http.Request) {
	AddAllowHeaders(w)
	if !auth.CheckAuthorization(r, p.Sessions) {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}
	body := http.MaxBytesReader(w, r.Body, maxBytesToRead)

	decoder := json.NewDecoder(body)
	bodyLike := &BodyLike{}
	err := decoder.Decode(bodyLike)
	if err != nil {
		http.Error(w, `{"error":"wrong_json"}`, http.StatusBadRequest)
		return
	}

	cookie, err := r.Cookie("session_id")
	if err != nil {
		http.Error(w, `{"error": "wrong cookie"}`, http.StatusBadRequest)
	}
	session, ok := p.Sessions.CheckSession(cookie.Value)
	if !ok {
		http.Error(w, `{"error" : "wrong cookie"}`, http.StatusBadRequest)
	}
	err = p.Likes.CreateNewLike(int(session.UserID), bodyLike.PostId)
	if err != nil {
		http.Error(w, `{"error":"wrong_json"}`, http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (p *PostHandler) UnlikePost(w http.ResponseWriter, r *http.Request) {
	AddAllowHeaders(w)
	if !auth.CheckAuthorization(r, p.Sessions) {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
		return
	}
	body := http.MaxBytesReader(w, r.Body, maxBytesToRead)

	decoder := json.NewDecoder(body)
	bodyLike := &BodyLike{}
	err := decoder.Decode(bodyLike)
	if err != nil {
		http.Error(w, `{"error":"wrong_json"}`, http.StatusBadRequest)
		return
	}

	cookie, err := r.Cookie("session_id")
	if err != nil {
		http.Error(w, `{"error": "wrong cookie"}`, http.StatusBadRequest)
	}
	session, ok := p.Sessions.CheckSession(cookie.Value)
	if !ok {
		http.Error(w, `{"error" : "wrong cookie"}`, http.StatusBadRequest)
	}
	err = p.Likes.DeleteLike(int(session.UserID), bodyLike.PostId)
	if err != nil {
		http.Error(w, `{"error":"wrong_json"}`, http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
