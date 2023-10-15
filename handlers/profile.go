package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	auth "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/authorization"
	reg "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/registration"
)

type BodyProfile struct {
	Profile model.Profile `json:"profile"`
}

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

	result := Result{Body: BodyProfile{Profile: profile}}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&result)
	if err == nil {
		log.Fatal(err)
	}
}
