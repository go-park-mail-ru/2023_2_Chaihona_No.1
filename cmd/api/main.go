package main

import (
	"fmt"
	"net/http"

	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/configs"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/db/postgresql"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/handlers"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/payments"
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
	paymentStorage := payments.CreatePaymentStorage(db.GetDB())

	// for _, testUser := range testdata.Users {
	// 	userStoarge.RegisterNewUser(&testUser)
	// }
	// for _, testProfile := range testdata.Profiles {
	// 	profileStorage.RegisterNewProfile(&testProfile)
	// }
	// for _, testPost := range testdata.Posts {
	// 	postStorage.CreateNewPost(testPost)
	// }

	rep := handlers.CreateRepoHandler(sessionStorage, userStoarge, profileStorage)
	profileHandler := handlers.CreateProfileHandlerViaRepos(sessionStorage, profileStorage)
	postHandler := handlers.CreatePostHandlerViaRepos(sessionStorage, postStorage, profileStorage)
	paymentHandler := handlers.CreatePaymentHandlerViaRepos(sessionStorage, paymentStorage)
	r := mux.NewRouter()

	r.Methods("OPTIONS").HandlerFunc(handlers.OptionsHandler)
	r.HandleFunc("/api/v1/login", rep.Login).Methods("POST")
	r.HandleFunc("/api/v1/logout", rep.Logout).Methods("POST")
	r.HandleFunc("/api/v1/registration", rep.Signup).Methods("POST")
	r.HandleFunc("/api/v1/is_authorized", rep.IsAuthorized).Methods("GET")
	r.HandleFunc("/api/v1/profile/{id:[0-9]+}", profileHandler.GetInfo).Methods("GET")
	r.HandleFunc("/api/v1/profile/{id:[0-9]+}/post", postHandler.GetAllUserPosts).Methods("GET")
	r.HandleFunc("/api/v1/profile/{id:[0-9]+}/donates", paymentHandler.GetAuthorDonates).Methods("GET")

	fmt.Println("Server started")
	err = http.ListenAndServe(configs.BackendServerPort, r)
	fmt.Println(err)
}
