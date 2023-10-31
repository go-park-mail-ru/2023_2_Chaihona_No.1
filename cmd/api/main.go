package main

import (
	"fmt"
	"net/http"

	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/configs"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/db/postgresql"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/handlers"
	postsrep "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/posts"
	profsrep "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/profiles"
	sessrep "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/sessions"
	usrep "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/users"
	"github.com/gorilla/mux"
)

func main() {
	var db postgresql.Database
	err := db.Connect()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	// err = db.MigrateUp()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	sessionStorage := sessrep.CreateSessionStorage()
	userStoarge := usrep.CreateUserStorage(db.GetDB())
	profileStorage := profsrep.CreateProfileStorage()
	postStorage := postsrep.CreatePostStorage(db.GetDB())

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
	r.HandleFunc("/api/v1/post/{id:[0-9]+}", postHandler.ChangePost).Methods("POST")
	r.HandleFunc("/api/v1/post", postHandler.CreateNewPost).Methods("POST")
	r.HandleFunc("/api/v1/post/{id:[0-9]+}", postHandler.DeletePost).Methods("DELETE")
	r.HandleFunc("/api/v1/feed", postHandler.GetFeed).Methods("GET")

	fmt.Println("Server started")
	err = http.ListenAndServe(configs.BackendServerPort, r)
	fmt.Println(err)
}
