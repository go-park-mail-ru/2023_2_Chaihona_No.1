package main

import (
	"fmt"
	"net/http"

	auth "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/authorization"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/handlers"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/model"
	reg "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/registration"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/test_data"

	"github.com/gorilla/mux"
)

func main() {
	sessionStorage := auth.CreateSessionStorage()
	userStoarge := reg.CreateUserStorage()
	profileStorage := reg.CreateProfileStorage()
	postStorage := model.CreatePostStorage()

	for _, test_user := range test_data.Users {
		userStoarge.RegisterNewUser(&test_user)
	}
	for _, test_profile := range test_data.Profiles {
		profileStorage.RegisterNewProfile(&test_profile)
	}
	for _, test_post := range test_data.Posts {
		postStorage.CreateNewPost(test_post)
	}

	rep := handlers.CreateRepoHandler(sessionStorage, userStoarge, profileStorage)
	profileHandler := handlers.CreateProfileHandlerViaRepos(sessionStorage, profileStorage)
	postHandler := handlers.CreatePostHandlerViaRepos(sessionStorage, postStorage, profileStorage)
	r := mux.NewRouter()

	r.Methods("OPTIONS").HandlerFunc(handlers.OptionsHandler)
	r.HandleFunc("/api/v1/login", rep.Login).Methods("POST")
	r.HandleFunc("/api/v1/logout", rep.Logout).Methods("POST")
	r.HandleFunc("/api/v1/registration", rep.Signup).Methods("POST")
	r.HandleFunc("/api/v1/is_authorized", rep.IsAuthorized).Methods("GET")
	r.HandleFunc("/api/v1/profile/{id:[0-9]+}", profileHandler.GetInfo).Methods("GET")
	r.HandleFunc("/api/v1/profile/{id:[0-9]+}/post", postHandler.GetAllUserPosts).Methods("GET")

	fmt.Println("Server started")
	err := http.ListenAndServe(":8001", r)
	fmt.Println(err)
}
