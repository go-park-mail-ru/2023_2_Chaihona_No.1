package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/profiles"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/sessions"
	auth "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/usecases/authorization"
)

type BodyProfile struct {
	Profile model.Profile `json:"profile"`
}

type ProfileHandler struct {
	Session  sessions.SessionRepository
	Profiles profiles.ProfileRepository
}

func CreateProfileHandlerViaRepos(
	session sessions.SessionRepository,
	profiles profiles.ProfileRepository,
) *ProfileHandler {
	return &ProfileHandler{
		session,
		profiles,
	}
}

// swagger:route OPTIONS /api/v1/profile/{id} Profile GetInfoOptions
// Handle OPTIONS request
// Responses:
//
//	200: result

// swagger:route GET /api/v1/profile/{id} Profile GetInfo
// Get profile info
//
// 
// Parameters:
//   + name: id
//     in: path
//     description: ID of user
//     required: true
//     type: integer
//     format: int
//
// Responses:
//
//	200: result
//	400: result
//	401: result
//	500: result
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

	result := Result{Body: BodyProfile{Profile: *profile}}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(&result)
	if err != nil {
		log.Println(err)
	}
}