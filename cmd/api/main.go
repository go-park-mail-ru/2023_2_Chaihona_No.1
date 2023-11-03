package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/configs"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/db/postgresql"
	_ "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/docs"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/handlers"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/likes"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/payments"
	postsrep "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/posts"
	sessrep "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/sessions"
	levels "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/subscribe_levels"
	subscriptionlevels "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/subscription_levels"
	subs "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/subscriptions"
	usrep "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/users"
	_ "github.com/go-swagger/go-swagger"
	"github.com/gomodule/redigo/redis"
	"github.com/gorilla/mux"
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

	sessionStorage := sessrep.CreateSessionStorage()
	userStoarge := usrep.CreateUserStorage(db.GetDB())
	levelStorage := levels.CreateSubscribeLevelStorage(db.GetDB())
	subsStorage := subs.CreateSubscriptionsStorage(db.GetDB())
	postStorage := posts.CreatePostStorage(db.GetDB())
	likeStorage := likes.CreateLikeStorage(db.GetDB())
	paymentStorage := payments.CreatePaymentStorage(db.GetDB())
	subscriptionsStorage := subscriptions.CreateSubscriptionsStorage(db.GetDB())
	subscriptionLevelsStorage := subscriptionlevels.CreateSubscribeLevelStorage(db.GetDB())

	rep := handlers.CreateRepoHandler(sessionStorage, userStoarge)
	profileHandler := handlers.CreateProfileHandlerViaRepos(sessionStorage, userStoarge, levelStorage, subsStorage)
	postHandler := handlers.CreatePostHandlerViaRepos(sessionStorage, postStorage, likeStorage)
	paymentHandler := handlers.CreatePaymentHandlerViaRepos(sessionStorage, paymentStorage, subscriptionsStorage, subscriptionLevelsStorage)

	r := mux.NewRouter()

	r.Methods("OPTIONS").HandlerFunc(handlers.OptionsHandler)
	r.HandleFunc("/api/v1/login", rep.Login).Methods("POST")
	r.HandleFunc("/api/v1/logout", rep.Logout).Methods("POST")
	r.HandleFunc("/api/v1/registration", rep.Signup).Methods("POST")
	r.HandleFunc("/api/v1/is_authorized", rep.IsAuthorized).Methods("GET")
	r.HandleFunc("/api/v1/profile/{id:[0-9]+}", profileHandler.GetInfo).Methods("GET")
	r.HandleFunc("/api/v1/profile/{id:[0-9]+}", profileHandler.ChangeUser).Methods("POST")
	r.HandleFunc("/api/v1/profile/{id:[0-9]+}", profileHandler.DeleteUser).Methods(http.MethodDelete)
	r.HandleFunc("/api/v1/profile/{id:[0-9]+}/post", postHandler.GetAllUserPosts).Methods("GET")
	r.HandleFunc("/api/v1/post/{id:[0-9]+}", postHandler.ChangePost).Methods("POST")
	r.HandleFunc("/api/v1/post", postHandler.CreateNewPost).Methods("POST")
	r.HandleFunc("/api/v1/post/{id:[0-9]+}", postHandler.DeletePost).Methods("DELETE")
	r.HandleFunc("/api/v1/feed", postHandler.GetFeed).Methods("GET")
	r.HandleFunc("/api/v1/post/like", postHandler.LikePost).Methods("POST")
	r.HandleFunc("/api/v1/post/unlike", postHandler.UnlikePost).Methods("DELETE")
	r.HandleFunc("/api/v1/profile/{id:[0-9]+}/donates", paymentHandler.GetAuthorDonates).Methods("GET")

	fmt.Println("Server started")
	err = http.ListenAndServe(configs.BackendServerPort, r)
	fmt.Println(err)
}
