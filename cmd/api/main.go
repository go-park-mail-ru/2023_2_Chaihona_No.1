package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	configs "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/configs"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/handlers"
	postsrep "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/posts"
	profsrep "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/profiles"
	sessrep "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/sessions"
	usrep "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/users"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/testdata"
)

func main() {
	sessionStorage := sessrep.CreateSessionStorage()
	userStoarge := usrep.CreateUserStorage()
	profileStorage := profsrep.CreateProfileStorage()
	postStorage := postsrep.CreatePostStorage()

	for _, testUser := range testdata.Users {
		userStoarge.RegisterNewUser(&testUser)
	}
	for _, testProfile := range testdata.Profiles {
		profileStorage.RegisterNewProfile(&testProfile)
	}
	for _, testPost := range testdata.Posts {
		postStorage.CreateNewPost(testPost)
	}

	rep := handlers.CreateRepoHandler(sessionStorage, userStoarge, profileStorage)
	profileHandler := handlers.CreateProfileHandlerViaRepos(sessionStorage, profileStorage)
	postHandler := handlers.CreatePostHandlerViaRepos(sessionStorage, postStorage, profileStorage)
	r := mux.NewRouter()

	r.Methods("OPTIONS").HandlerFunc(handlers.OptionsHandler)
	r.HandleFunc("/api/v1/login", handlers.NewWrapper(rep.LoginStrategy).ServeHTTP).Methods("POST")
	r.HandleFunc("/api/v1/logout", handlers.NewWrapper(rep.LogoutStrategy).ServeHTTP).Methods("POST")
	r.HandleFunc("/api/v1/registration", handlers.NewWrapper(rep.SignupStrategy).ServeHTTP).Methods("POST")
	r.HandleFunc("/api/v1/is_authorized", handlers.NewWrapper(rep.IsAuthorizedStrategy).ServeHTTP).Methods("GET")
	r.HandleFunc("/api/v1/profile/{id:[0-9]+}", handlers.NewWrapper(profileHandler.GetInfoStrategy).ServeHTTP).Methods("GET")
	r.HandleFunc("/api/v1/profile/{id:[0-9]+}/post", postHandler.GetAllUserPosts).Methods("GET")

	fmt.Println("Server started")
	err := http.ListenAndServe(configs.BackendServerPort, r)
	fmt.Println(err)
}
