package main

import (
	"net/http"
	auth "project/authorization"
	handlers "project/handlers"
	model "project/model"

	"github.com/gorilla/mux"
)

func main() {
	// registrationHandler := reg.CreateRepoHandler()
	// profileHandler := handlers.CreateProfileHandler()
	// postHandler := handlers.CreatePostHandler()
	rep := auth.CreateRepoHandler()
	profileHandler := handlers.CreateProfileHandlerViaRepos(&rep.Sessions, &rep.Profiles)
	postStorage := model.CreatePostStorage()
	postHandler := handlers.CreatePostHandlerViaRepos(&rep.Sessions, &postStorage, &rep.Profiles)

	r := mux.NewRouter()
	r.Methods("OPTIONS").HandlerFunc(handlers.OptionsHandler)
	r.HandleFunc("/api/v1/login", rep.Login).Methods("POST")
	// r.HandleFunc("/api/v1/logout", rep.Logout).Methods("POST")
	r.HandleFunc("/api/v1/registration", rep.Signup).Methods("POST")
	r.HandleFunc("/api/v1/profile/{id:[0-9]+}", profileHandler.GetInfo).Methods("GET")
	r.HandleFunc("/api/v1/profile/{id:[0-9]+}/post", postHandler.GetAllUserPosts).Methods("GET")

	http.ListenAndServeTLS(":8001", "cert.pem", "key.pem", r)
}
