package handlers

import (
	"encoding/json"
	"net/http"
	auth "project/authorization"
	model "project/model"
	"strconv"

	"github.com/gorilla/mux"
)

type Result struct {
	Body interface{} `json:"body,omitempty"`
	Err  string      `json:"err,omitempty"`
}

type ProfileHandler struct {
	Session  auth.SessionRepository
	Profiles model.ProfileRepository
}

func CreateProfileHandler() *ProfileHandler {
	return &ProfileHandler{
		auth.CreateSessionStorage(),
		model.CreateProfileStorage(),
	}
}

func (p *ProfileHandler) GetInfo(w http.ResponseWriter, r *http.Request) {
	if auth.CheckAuthorization(r, p.Session) {
		http.Error(w, `{"error":"unauthorized"}`, 401)
		return
	}

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, `{"error":"bad id"}`, 400)
		return
	}

	profile, ok := p.Profiles.GetProfile(uint(id))
	if !ok {
		http.Error(w, `{"error":"db"}`, 500)
		return
	}

	body := map[string]interface{}{
		"profiles": profile,
	}
	json.NewEncoder(w).Encode(&Result{Body: body})
}
