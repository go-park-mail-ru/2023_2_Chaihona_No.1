package main

import (
	"fmt"
	"net/http"

	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/configs"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/db/postgresql"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/handlers"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/likes"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/posts"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/profiles"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/sessions"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/users"
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

	sessionStorage := sessions.CreateSessionStorage()
	userStoarge := users.CreateUserStorage(db.GetDB())
	profileStorage := profiles.CreateProfileStorage()
	postStorage := posts.CreatePostStorage()
	likeStorage := likes.CreateLikeStorage(db.GetDB())

	rep := handlers.CreateRepoHandler(sessionStorage, userStoarge, profileStorage)
	profileHandler := handlers.CreateProfileHandlerViaRepos(sessionStorage, profileStorage)
	postHandler := handlers.CreatePostHandlerViaRepos(sessionStorage, postStorage, profileStorage, likeStorage)
	r := mux.NewRouter()

	r.Methods("OPTIONS").HandlerFunc(handlers.OptionsHandler)
	r.HandleFunc("/api/v1/login", rep.Login).Methods("POST")
	r.HandleFunc("/api/v1/logout", rep.Logout).Methods("POST")
	r.HandleFunc("/api/v1/registration", rep.Signup).Methods("POST")
	r.HandleFunc("/api/v1/is_authorized", rep.IsAuthorized).Methods("GET")
	r.HandleFunc("/api/v1/profile/{id:[0-9]+}", profileHandler.GetInfo).Methods("GET")
	r.HandleFunc("/api/v1/profile/{id:[0-9]+}/post", postHandler.GetAllUserPosts).Methods("GET")
	r.HandleFunc("/api/v1/post/like", postHandler.LikePost).Methods("POST")
	r.HandleFunc("/api/v1/post/unlike", postHandler.UnlikePost).Methods("DELETE")

	fmt.Println("Server started")
	err = http.ListenAndServe(configs.BackendServerPort, r)
	fmt.Println(err)
}
