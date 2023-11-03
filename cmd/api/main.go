package main

import (
	"fmt"
	"log"

	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/db/postgresql"
	_ "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/docs"
	_ "github.com/go-swagger/go-swagger"
	"github.com/gomodule/redigo/redis"

	configs "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/configs"
	postsrep "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/posts"
	profsrep "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/profiles"
	sessrep "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/sessions"
	usrep "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/users"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/testdata"
)

func main() {
	conn, err := redis.DialURL(fmt.Sprintf("redis://@%s:%s", configs.RedisServerIP, configs.RedisServerPort))
	if err != nil {
		log.Println(err)
		return
	}

	sessionStorage := sessrep.CreateRedisSessionStorage(conn)
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

	// users := squirrel.Select("*").From("public.notification")
	// type answ struct {
	// 	A string
	// 	B string
	// 	C string
	// 	D string
	// 	E string
	// }
	// var answ_ins answ
	// ans := []interface{}{&answ_ins.A, &answ_ins.B, &answ_ins.C, &answ_ins.D, &answ_ins.E}
	// err = users.RunWith(db).QueryRow().Scan(ans...)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(*ans[0].(*string))

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
