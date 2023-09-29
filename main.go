package main

import (
	"net/http"
	auth "project/authorization"
	profile "project/profile"
	reg "project/registration"

	"github.com/gorilla/mux"
)

func main() {
	authHandler := auth.CreateRepoHandler()
	registrationHandler := reg.CreateRepoHandler()
	profileHandler := profile.CreateRepoHandler()

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/login", authHandler.Login).Methods("POST")
	r.HandleFunc("/api/v1/logout", authHandler.Logout).Methods("POST")
	r.HandleFunc("/api/v1/registration", registrationHandler.SignUp).Methods("POST")
	r.HandleFunc("/api/v1/profile/{id:[0-9]+}", profileHandler.GetInfo).Methods("GET")

	http.ListenAndServe(":8080", r)
}
