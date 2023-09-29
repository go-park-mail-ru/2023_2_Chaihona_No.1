package main

import (
	"net/http"
	auth "project/authorization"
	handlers "project/handlers"
	reg "project/registration"

	"github.com/gorilla/mux"
)

func main() {
	authHandler := auth.CreateRepoHandler()
	registrationHandler := reg.CreateRepoHandler()
	profileHandler := handlers.CreateProfileHandler()
	postHandler := handlers.CreatePostHandler()

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/login", authHandler.Login).Methods("POST")
	r.HandleFunc("/api/v1/logout", authHandler.Logout).Methods("POST")
	r.HandleFunc("/api/v1/registration", registrationHandler.SignUp).Methods("POST")
	r.HandleFunc("/api/v1/profile/{id:[0-9]+}", profileHandler.GetInfo).Methods("GET")
	r.HandleFunc("/api/v1/profile/{id:[0-9]+}/post", postHandler.GetAllUserPosts).Methods("GET")

	http.ListenAndServe(":8080", r)
}
