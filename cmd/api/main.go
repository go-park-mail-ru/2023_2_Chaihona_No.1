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
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/posts"
	sessrep "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/sessions"
	levels "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/subscribe_levels"
	subscriptionlevels "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/subscription_levels"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/subscriptions"
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

	var db postgresql.Database
	err = db.Connect()
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

	sessionStorage := sessrep.CreateRedisSessionStorage(conn)
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
	r.HandleFunc("/api/v1/login", handlers.NewWrapper(rep.LoginStrategy).ServeHTTP).Methods("POST")
	r.HandleFunc("/api/v1/logout", handlers.NewWrapper(rep.LogoutStrategy).ServeHTTP).Methods("POST")
	r.HandleFunc("/api/v1/registration", handlers.NewWrapper(rep.SignupStrategy).ServeHTTP).Methods("POST")
	r.HandleFunc("/api/v1/is_authorized", handlers.NewWrapper(rep.IsAuthorizedStrategy).ServeHTTP).Methods("GET")
	r.HandleFunc("/api/v1/profile/{id:[0-9]+}", handlers.NewWrapper(profileHandler.GetInfoStrategy).ServeHTTP).Methods("GET")
	r.HandleFunc("/api/v1/profile/{id:[0-9]+}", profileHandler.ChangeUser).Methods("POST")
	r.HandleFunc("/api/v1/profile/{id:[0-9]+}", profileHandler.DeleteUser).Methods(http.MethodDelete)
	r.HandleFunc("/api/v1/profile/{id:[0-9]+}/post", handlers.NewWrapper(postHandler.GetAllUserPostsStrategy).ServeHTTP).Methods("GET")
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
