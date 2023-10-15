package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	auth "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/authorization"
	model "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/model"
	reg "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/registration"
	"github.com/gorilla/mux"
)

type BodyPosts struct {
	Posts []model.Post `json:"posts"`
}

type PostHandler struct {
	Sessions auth.SessionRepository
	Posts    model.PostRepository
	Profiles reg.ProfileRepository
}

func CreatePostHandlerViaRepos(session auth.SessionRepository, posts model.PostRepository,
	profiles reg.ProfileRepository,
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

	// authorId, err := strconv.Atoi(strings.Split(r.URL.Path, "/")[3])
	vars := mux.Vars(r)
	authorId, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, `{"error":"bad id"}`, 400)
		return
	}

	posts, errPost := p.Posts.GetPostsByAuthorId(uint(authorId))
	if errPost != nil {
		if errors.Is(ErrorNotAuthor, errPost.Err) {
			res := Result{Err: errPost.Err.Error()}
			errJson, err := json.Marshal(res)
			if err != nil {
				errJson = []byte{}
			}
			http.Error(w, string(errJson), errPost.StatusCode)
			return
		}
		http.Error(w, `{"error":"db"}`, errPost.StatusCode)
		return
	}

	cookie, _ := r.Cookie("session_id")
	session, _ := p.Sessions.CheckSession(cookie.Value)
	userId := session.UserId
	profile, ok := p.Profiles.GetProfile(uint(userId))
	if !ok {
		http.Error(w, `{"error":"very bad"}`, 400)
		return
	}
	subscriptions := profile.Subscriptions
	isSubscirber := false
	for _, user := range subscriptions {
		if user.ID == uint(authorId) {
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
		log.Fatal(err)
	}
}
