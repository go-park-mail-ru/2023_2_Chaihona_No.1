package main

import (
	"fmt"
	"net/http"
	"project/model"
	"project/handlers"
	auth "project/authorization"
	reg "project/registration"

	"github.com/gorilla/mux"
)

// заглушка main
func main() {
	sessionStorage := auth.CreateSessionStorage()
	userStoarge := reg.CreateUserStorage()
	profileStorage := reg.CreateProfileStorage()
	postStorage := model.CreatePostStorage()

	
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
	http.ListenAndServe(":8001", r)
}
