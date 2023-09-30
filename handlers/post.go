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
	Profiles model.ProfileRepository
}

func CreatePostHandler() *PostHandler {
	return &PostHandler{
		auth.CreateSessionStorage(),
		model.CreatePostStorage(),
		model.CreateProfileStorage(),
	}
}

func (p *PostHandler) GetAllUserPosts(w http.ResponseWriter, r *http.Request) {
	if !auth.CheckAuthorization(r, p.Sessions) {
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
			return
		}
		http.Error(w, `{"error":"db"}`, 500)
		return
	}

	cookie, _ := r.Cookie("session_id")
	session, _ := p.Sessions.CheckSession(cookie.Value)
	userId := session.UserId
	profile, _ := p.Profiles.GetProfile(uint(userId))
	subscribtions := profile.Subscribtions

	isSubscirber := false
	for _, user := range subscribtions {
		if user.ID == userId {
			isSubscirber = true
		}
	}
	// Need to add subscribtion level check logic and one-time payment check logic
	if !isSubscirber {
		for post := range posts {
			if post.Access == model.SubscribersAccess {
				post.HasAccess = false
				post.Reason = model.LowLevelReason
				post.body = ""
			}
		}
	}

	body := map[string]interface{}{
		"posts": posts,
	}
	json.NewEncoder(w).Encode(&Result{Body: body})
}
