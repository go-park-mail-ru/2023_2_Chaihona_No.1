package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	auth "project/authorization"
	reg "project/registration"
	"strconv"

	"github.com/gorilla/mux"
)

type ProfileHandler struct {
	Session  *auth.SessionRepository
	Profiles *reg.ProfileRepository
}

// func CreateProfileHandler() *ProfileHandler {
// 	return &ProfileHandler{
// 		auth.CreateSessionStorage(),
// 		reg.CreateProfileStorage(),
// 	}
// }

func CreateProfileHandlerViaRepos(session *auth.SessionRepository, profiles *reg.ProfileRepository) *ProfileHandler {
	return &ProfileHandler{
		session,
		profiles,
	}
}

func (p *ProfileHandler) GetInfo(w http.ResponseWriter, r *http.Request) {
	AddAllowHeaders(w)
	if !auth.CheckAuthorization(r, *p.Session) {
		http.Error(w, `{"error":"unauthorized"}`, 401)
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, `{"error":"bad id"}`, 400)
		return
	}

	profile, ok := (*p.Profiles).GetProfile(uint(id))
	fmt.Println(profile)
	if !ok {
		http.Error(w, `{"error":"db"}`, 500)
		return
	}

	body := map[string]interface{}{
		"profiles": profile,
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&Result{Body: body})
}
