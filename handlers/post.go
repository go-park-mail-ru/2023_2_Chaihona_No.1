package handlers

import (
	"encoding/json"
	"net/http"
	auth "project/authorization"
	model "project/model"
	reg "project/registration"
	"strconv"
	"strings"
)

type PostHandler struct {
	Sessions *auth.SessionRepository
	Posts    *model.PostRepository
	Profiles *reg.ProfileRepository
}

// func CreatePostHandler() *PostHandler {
// 	return &PostHandler{
// 		auth.CreateSessionStorage(),
// 		model.CreatePostStorage(),
// 		model.CreateProfileStorage(),
// 	}
// }

func CreatePostHandlerViaRepos(session *auth.SessionRepository, posts *model.PostRepository,
	profiles *reg.ProfileRepository) *PostHandler {
	return &PostHandler{
		session,
		posts,
		profiles,
	}
}

func (p *PostHandler) GetAllUserPosts(w http.ResponseWriter, r *http.Request) {
	if !auth.CheckAuthorization(r, *p.Sessions) {
		http.Error(w, `{"error":"unauthorized"}`, 401)
		return
	}

	// vars := mux.Vars(r)
	// authorId, err := strconv.Atoi(vars["id"])
	authorId, err := strconv.Atoi(strings.Split(r.URL.Path, "/")[4])
	if err != nil {
		http.Error(w, `{"error":"bad id"}`, 400)
		return
	}

	posts, err := (*p.Posts).GetPostsByAuthorId(uint(authorId))
	if err != nil {
		if err == NotAuthorError {
			http.Error(w, `{"error":"bad request"}`, 400)
			return
		}
		http.Error(w, `{"error":"db"}`, 500)
		return
	}

	cookie, _ := r.Cookie("session_id")
	session, _ := (*p.Sessions).CheckSession(cookie.Value)
	userId := session.UserId
	profile, _ := (*p.Profiles).GetProfile(uint(userId))
	subscribtions := profile.Subscribtions

	isSubscirber := false
	for _, user := range subscribtions {
		if user.ID == uint(authorId) {
			isSubscirber = true
		}
	}

	// Need to add subscribtion level check logic and one-time payment check logic
	if !isSubscirber {
		for i, _ := range *posts {
			switch (*posts)[i].Access {
			case model.SubscribersAccess:
				(*posts)[i].HasAccess = false
				(*posts)[i].Reason = model.LowLevelReason
				(*posts)[i].Body = ""
			case model.EveryoneAccess:
				(*posts)[i].HasAccess = true
			}
		}
	} else {
		for i, _ := range *posts {
			(*posts)[i].HasAccess = true
		}
	}

	body := map[string]interface{}{
		"posts": posts,
	}
	json.NewEncoder(w).Encode(&Result{Body: body})
}
