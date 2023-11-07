package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

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
	subs "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/subscriptions"
	usrep "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/users"
	_ "github.com/go-swagger/go-swagger"
	"github.com/gorilla/csrf"
	h "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// conn, err := redis.DialURL(fmt.Sprintf("redis://@%s:%s", configs.RedisServerIP, configs.RedisServerPort))
	pool := sessrep.NewPool(fmt.Sprintf("%s:%s", configs.RedisServerIP, configs.RedisServerPort))
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

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
	//

	sessionStorage := sessrep.CreateRedisSessionStorage(pool)
	userStoarge := usrep.CreateUserStorage(db.GetDB())
	levelStorage := levels.CreateSubscribeLevelStorage(db.GetDB())
	subsStorage := subs.CreateSubscriptionsStorage(db.GetDB())
	postStorage := posts.CreatePostStorage(db.GetDB())
	likeStorage := likes.CreateLikeStorage(db.GetDB())
	paymentStorage := payments.CreatePaymentStorage(db.GetDB())
	subscriptionLevelsStorage := subscriptionlevels.CreateSubscribeLevelStorage(db.GetDB())

	rep := handlers.CreateRepoHandler(sessionStorage, userStoarge, levelStorage)
	profileHandler := handlers.CreateProfileHandlerViaRepos(sessionStorage, userStoarge, levelStorage, subsStorage, paymentStorage)
	postHandler := handlers.CreatePostHandlerViaRepos(sessionStorage, postStorage, likeStorage)
	paymentHandler := handlers.CreatePaymentHandlerViaRepos(sessionStorage, paymentStorage, subsStorage, subscriptionLevelsStorage)
	fileHandler := handlers.CreateFileHandler(sessionStorage, userStoarge)
	r := mux.NewRouter()

	r.Methods("OPTIONS").HandlerFunc(handlers.OptionsHandler)
	r.HandleFunc("/api/v1/login", handlers.NewWrapper(rep.LoginStrategy).ServeHTTP).Methods("POST")
	r.HandleFunc("/api/v1/logout", handlers.NewWrapper(rep.LogoutStrategy).ServeHTTP).Methods("POST")
	r.HandleFunc("/api/v1/registration", handlers.NewWrapper(rep.SignupStrategy).ServeHTTP).Methods("POST")
	r.HandleFunc("/api/v1/is_authorized", handlers.NewWrapper(rep.IsAuthorizedStrategy).ServeHTTP).Methods("GET")

	r.HandleFunc("/api/v1/profile/{id:[0-9]+}", handlers.NewWrapper(profileHandler.GetInfoStrategy).ServeHTTP).Methods("GET")
	// r.HandleFunc("/api/v1/profile/{id:[0-9]+}", handlers.NewWrapper(profileHandler.ChangeUserStratagy).ServeHTTP).Methods("POST")
	r.HandleFunc("/api/v1/profile/{id:[0-9]+}", handlers.NewFileWrapper(profileHandler.ChangeUserStratagy).ServeHTTP).Methods("POST")
	r.HandleFunc("/api/v1/profile/{id:[0-9]+}/status", handlers.NewWrapper(profileHandler.ChangeUserStatusStratagy).ServeHTTP).Methods("POST")
	r.HandleFunc("/api/v1/profile/{id:[0-9]+}/description", handlers.NewWrapper(profileHandler.ChangeUserDescriptionStratagy).ServeHTTP).Methods("POST")
	r.HandleFunc("/api/v1/profile/{id:[0-9]+}", handlers.NewWrapper(profileHandler.DeleteUserStratagy).ServeHTTP).Methods(http.MethodDelete)

	r.HandleFunc("/api/v1/profile/{id:[0-9]+}/post", handlers.NewWrapper(postHandler.GetAllUserPostsStrategy).ServeHTTP).Methods("GET")
	r.HandleFunc("/api/v1/post", handlers.NewWrapper(postHandler.CreateNewPostStrategy).ServeHTTP).Methods("POST")
	r.HandleFunc("/api/v1/post/{id:[0-9]+}", handlers.NewWrapper(postHandler.ChangePostStrategy).ServeHTTP).Methods("POST")
	r.HandleFunc("/api/v1/post/{id:[0-9]+}", handlers.NewWrapper(postHandler.DeletePostStrategy).ServeHTTP).Methods("DELETE")
	r.HandleFunc("/api/v1/feed", handlers.NewWrapper(postHandler.GetFeedStrategy).ServeHTTP).Methods("GET")

	r.HandleFunc("/api/v1/post/{id:[0-9]+}/like", handlers.NewWrapper(postHandler.LikePostStrategy).ServeHTTP).Methods("POST")
	r.HandleFunc("/api/v1/post/{id:[0-9]+}/unlike", handlers.NewWrapper(postHandler.UnlikePostStrategy).ServeHTTP).Methods("DELETE")

	r.HandleFunc("/api/v1/profile/{id:[0-9]+}/donates/creator", handlers.NewWrapper(paymentHandler.GetAuthorDonatesStratagy).ServeHTTP).Methods("GET")
	r.HandleFunc("/api/v1/profile/{id:[0-9]+}/donates/donater", handlers.NewWrapper(paymentHandler.GetUsersDonatesStratagy).ServeHTTP).Methods("GET")
	r.HandleFunc("/api/v1/donate", handlers.NewWrapper(paymentHandler.DonateStratagy).ServeHTTP).Methods("POST")

	r.HandleFunc("/api/v1/upload", handlers.NewFileWrapper(fileHandler.UploadFileStratagy).ServeHTTP).Methods("POST")
	r.HandleFunc("/api/v1/profile/{id:[0-9]+}/avatar", fileHandler.LoadFileStratagy).Methods("GET")
	//probably different wrapper
	r.HandleFunc("/api/v1/profile/{id:[0-9]+}/follow", handlers.NewWrapper(profileHandler.FollowStratagy).ServeHTTP).Methods("POST")
	r.HandleFunc("/api/v1/profile/{id:[0-9]+}/unfollow", handlers.NewWrapper(profileHandler.UnfollowStratagy).ServeHTTP).Methods("POST")

	CSRFMiddleware := csrf.Protect(
		[]byte("place-your-32-byte-long-key-here"),
		csrf.Secure(false),                 // false in development only!
		csrf.RequestHeader("X-CSRF-Token"), // Must be in CORS Allowed and Exposed Headers
	)

	// APIRouter := r.PathPrefix("/api").Subrouter()
	r.Use(CSRFMiddleware)

	CORSMiddleware := h.CORS(
		h.AllowCredentials(),
		h.AllowedOriginValidator(
			func(origin string) bool {
				return strings.HasPrefix(origin, "http://212.233.89.163:8001")
			},
		),
		h.AllowedMethods([]string{ "GET", "POST", "OPTIONS", "DELETE"}),
		h.AllowedHeaders([]string{"X-CSRF-Token"}),
		h.ExposedHeaders([]string{"X-CSRF-Token"}),
	)

	server := &http.Server{
		Handler:      CORSMiddleware(r),
		Addr:         "localhost:8000",
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}

	fmt.Println("Server started")
	err = server.ListenAndServe()
	// err = http.ListenAndServe(configs.BackendServerPort, r)
	fmt.Println(err)
}