package main

import (
	"fmt"
	"net/http"
	"project/model"
	"project/handlers"

	"github.com/gorilla/mux"
)

// заглушка main
func main() {
	rep := handlers.CreateRepoHandler()
	profileHandler := handlers.CreateProfileHandlerViaRepos(rep.Sessions, rep.Profiles)
	postStorage := model.CreatePostStorage()
	postHandler := handlers.CreatePostHandlerViaRepos(rep.Sessions, postStorage, rep.Profiles)
	

	r := mux.NewRouter()
	r.Methods("OPTIONS").HandlerFunc(handlers.OptionsHandler)
	r.HandleFunc("/api/v1/login", rep.Login).Methods("POST")
	r.HandleFunc("/api/v1/logout", rep.Logout).Methods("POST")
	r.HandleFunc("/api/v1/registration", rep.Signup).Methods("POST")
	r.HandleFunc("/api/v1/is_authorized", rep.IsAuthorized).Methods("GET")
	r.HandleFunc("/api/v1/profile/{id:[0-9]+}", profileHandler.GetInfo).Methods("GET")
	r.HandleFunc("/api/v1/profile/{id:[0-9]+}/post", postHandler.GetAllUserPosts).Methods("GET")

	fmt.Println("Server started")
	http.ListenAndServe(":8001", r)
}
