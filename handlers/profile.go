package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	auth "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/authorization"
	reg "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/registration"
)

type ProfileHandler struct {
	Session  auth.SessionRepository
	Profiles reg.ProfileRepository
}

func CreateProfileHandlerViaRepos(
	session auth.SessionRepository,
	profiles reg.ProfileRepository,
) *ProfileHandler {
	return &ProfileHandler{
		session,
		profiles,
	}
}

func (p *ProfileHandler) GetInfo(w http.ResponseWriter, r *http.Request) {
	AddAllowHeaders(w)
	if !auth.CheckAuthorization(r, p.Session) {
		http.Error(w, `{"error":"unauthorized"}`, http.StatusUnauthorized)
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
	w.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&Result{Body: body})
	if err == nil {
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, `{"error":"json_encoding"}`, http.StatusInternalServerError)
	}
}
