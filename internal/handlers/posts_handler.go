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
		return
	}

	// probably get user subscriptions from cookie!!!
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
