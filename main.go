package main

import (
	"fmt"
	"net/http"
	"project/handlers"
	"project/model"
	"project/test_data"

	"github.com/gorilla/mux"
)

func main() {
	rep := handlers.CreateRepoHandler()
	profileHandler := handlers.CreateProfileHandlerViaRepos(&rep.Sessions, &rep.Profiles)
	postStorage := model.CreatePostStorage()
	postHandler := handlers.CreatePostHandlerViaRepos(rep.Sessions, postStorage, rep.Profiles)
	for _, test_user := range test_data.Users {
		rep.Users.RegisterNewUser(&test_user)
	}

	for _, test_profile := range test_data.Profiles {
		rep.Profiles.RegisterNewProfile(&test_profile)
	}
	for _, test_post := range test_data.Posts {
		postStorage.CreateNewPost(test_post)
	}

	r := mux.NewRouter()
	r.Methods("OPTIONS").HandlerFunc(handlers.OptionsHandler)
	r.HandleFunc("/api/v1/login", rep.Login).Methods("POST")
	r.HandleFunc("/api/v1/logout", rep.Logout).Methods("POST")
	r.HandleFunc("/api/v1/registration", rep.Signup).Methods("POST")
	r.HandleFunc("/api/v1/is_authorized", rep.IsAuthorized).Methods("GET")
	r.HandleFunc("/api/v1/profile/{id:[0-9]+}", profileHandler.GetInfo).Methods("GET")
	r.HandleFunc("/api/v1/profile/{id:[0-9]+}/post", postHandler.GetAllUserPosts).Methods("GET")

	fmt.Println("start")
	http.ListenAndServe(":8001", r)
}
