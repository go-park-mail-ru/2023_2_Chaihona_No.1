package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	conn, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close()
	var greeting string
	err = conn.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)
	// sessionStorage := sessrep.CreateSessionStorage()
	// userStoarge := usrep.CreateUserStorage()
	// profileStorage := profsrep.CreateProfileStorage()
	// postStorage := postsrep.CreatePostStorage()

	// for _, testUser := range testdata.Users {
	// 	userStoarge.RegisterNewUser(&testUser)
	// }
	// for _, testProfile := range testdata.Profiles {
	// 	profileStorage.RegisterNewProfile(&testProfile)
	// }
	// for _, testPost := range testdata.Posts {
	// 	postStorage.CreateNewPost(testPost)
	// }

	// rep := handlers.CreateRepoHandler(sessionStorage, userStoarge, profileStorage)
	// profileHandler := handlers.CreateProfileHandlerViaRepos(sessionStorage, profileStorage)
	// postHandler := handlers.CreatePostHandlerViaRepos(sessionStorage, postStorage, profileStorage)
	// r := mux.NewRouter()

	// r.Methods("OPTIONS").HandlerFunc(handlers.OptionsHandler)
	// r.HandleFunc("/api/v1/login", rep.Login).Methods("POST")
	// r.HandleFunc("/api/v1/logout", rep.Logout).Methods("POST")
	// r.HandleFunc("/api/v1/registration", rep.Signup).Methods("POST")
	// r.HandleFunc("/api/v1/is_authorized", rep.IsAuthorized).Methods("GET")
	// r.HandleFunc("/api/v1/profile/{id:[0-9]+}", profileHandler.GetInfo).Methods("GET")
	// r.HandleFunc("/api/v1/profile/{id:[0-9]+}/post", postHandler.GetAllUserPosts).Methods("GET")

	// fmt.Println("Server started")
	// err := http.ListenAndServe(configs.BackendServerPort, r)
	// fmt.Println(err)
}
