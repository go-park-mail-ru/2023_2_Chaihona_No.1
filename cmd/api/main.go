package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
	"google.golang.org/grpc"

	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/configs"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/db/postgresql"
	_ "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/docs"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/handlers"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/analytics"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/attaches"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/comments"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/devices"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/likes"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/multistorage"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/payments"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/posts"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/questions"
	sessrep "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/sessions"
	levels "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/subscribe_levels"
	subscriptionlevels "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/subscription_levels"
	subs "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/subscriptions"
	usrep "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/users"
	anal "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/usecases/analytics"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/usecases/payment"
	_ "github.com/go-swagger/go-swagger"
	"github.com/gorilla/mux"
)

func SetupFirebase() (*firebase.App, context.Context, *messaging.Client) {

	ctx := context.Background()

	// fileDir, _ := os.Getwd()
	serviceAccountKeyFilePath, err := filepath.Abs("./serviceAccountKey.json")
	if err != nil {
		log.Println("Unable to load serviceAccountKeys.json file")
	}

	opt := option.WithCredentialsFile(serviceAccountKeyFilePath)
	// fmt.Println(opt)

	config := &firebase.Config{ProjectID: "kopilka-71492"}
	//Firebase admin SDK initialization
	// app, err := firebase.NewApp(context.Background(), nil, opt)
	app, err := firebase.NewApp(context.Background(), config, opt)
	if err != nil {
		log.Println("Firebase load error")
	}

	//Messaging client
	client, _ := app.Messaging(ctx)

	return app, ctx, client
}

func main() {
	// conn, err := redis.DialURL(fmt.Sprintf("redis://@%s:%s", configs.RedisServerIP, configs.RedisServerPort))
	//pool := sessrep.NewPool(fmt.Sprintf("%s:%s", configs.RedisServerIP, configs.RedisServerPort))
	//sessManager := sessrep.CreateRedisManager(pool)
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

	grcpConnSessions, err := grpc.Dial(
		"212.233.89.163:8081",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("cant connect to grpc")
	}
	defer grcpConnSessions.Close()

	grcpConnPayments, err := grpc.Dial(
		"212.233.89.163:8082",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("cant connect to grpc")
	}
	defer grcpConnPayments.Close()

	grcpConnPosts, err := grpc.Dial(
		"212.233.89.163:8083",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("cant connect to grpc")
	}
	defer grcpConnPosts.Close()

	sessManager := &sessrep.RedisManager{sessrep.NewAuthCheckerClient(grcpConnSessions)}
	payManager := &payments.PaymentManager{payments.NewPaymentsServiceClient(grcpConnPayments)}
	postManager := &posts.PostManager{posts.NewPostsServiceClient(grcpConnPosts)}
	commentManager := &comments.CommentManager{comments.NewCommentServiceClient(grcpConnPosts)}
	// err = db.MigrateUp()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	//

	// sessionStorage := sessrep.CreateRedisSessionStorage(pool)
	userStoarge := usrep.CreateUserStorage(db.GetDB())
	levelStorage := levels.CreateSubscribeLevelStorage(db.GetDB())
	subsStorage := subs.CreateSubscriptionsStorage(db.GetDB())
	// postStorage := posts.CreatePostStorage(db.GetDB())
	likeStorage := likes.CreateLikeStorage(db.GetDB())
	// paymentStorage := payments.CreatePaymentStorage(db.GetDB())
	questionsStorage := questions.CreateQuestionStorage(db.GetDB())
	// answersStorage := answers.CreateAnswerStorage(db.GetDB())
	subscriptionLevelsStorage := subscriptionlevels.CreateSubscribeLevelStorage(db.GetDB())
	attachStorage := attaches.CreateAttachStorage(db.GetDB())
	analyticsStorage := analytics.CreateAnalyticsStorage(db.GetDB())
	deviceStorage := devices.CreateDevicesStorage(db.GetDB())
	app, _, _ := SetupFirebase()

	rep := handlers.CreateRepoHandler(sessManager, userStoarge, levelStorage, analyticsStorage)
	profileHandler := handlers.CreateProfileHandlerViaRepos(sessManager, userStoarge, 
		levelStorage, subsStorage, 
		payManager, analyticsStorage, 
		deviceStorage, app)
	postHandler := handlers.CreatePostHandlerViaRepos(sessManager, postManager, likeStorage, attachStorage, commentManager, deviceStorage, app)
	paymentHandler := handlers.CreatePaymentHandlerViaRepos(sessManager, payManager, subsStorage, subscriptionLevelsStorage)
	fileHandler := handlers.CreateFileHandler(sessManager, userStoarge, attachStorage)
	csatHandler := handlers.CreateCSATHandler(sessManager,questionsStorage)
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
	r.HandleFunc("/api/v1/profiles/{limit:[0-9]+}", handlers.NewWrapper(profileHandler.GetTopUsersStratagy).ServeHTTP).Methods("GET")
	r.HandleFunc("/api/v1/post", handlers.NewWrapper(postHandler.CreateNewPostStrategy).ServeHTTP).Methods("POST")
	r.HandleFunc("/api/v1/post/{id:[0-9]+}", handlers.NewWrapper(postHandler.ChangePostStrategy).ServeHTTP).Methods("POST")
	r.HandleFunc("/api/v1/post/{id:[0-9]+}", handlers.NewWrapper(postHandler.DeletePostStrategy).ServeHTTP).Methods("DELETE")
	r.HandleFunc("/api/v1/feed", handlers.NewWrapper(postHandler.GetFeedStrategy).ServeHTTP).Methods("GET")
	r.HandleFunc("/api/v1/tag", handlers.NewWrapper(postHandler.GetPostByTagStrategy).ServeHTTP).Methods("GET")

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

	r.HandleFunc("/api/v1/questions", handlers.NewWrapper(csatHandler.GetQuestions).ServeHTTP).Methods("GET")
	r.HandleFunc("/api/v1/rate/{id:[0-9]+}", handlers.NewWrapper(csatHandler.Rate).ServeHTTP).Methods("POST")
	r.HandleFunc("/api/v1/statistic", handlers.NewWrapper(csatHandler.GetStatistic).ServeHTTP).Methods("GET")
	
	r.HandleFunc("/api/v1/post/{id:[0-9]+}/attaches", handlers.NewWrapper(fileHandler.LoadAttachesStratagy).ServeHTTP).Methods("GET")

	r.HandleFunc("/api/v1/search/{nickname:.*}", handlers.NewWrapper(profileHandler.Search).ServeHTTP).Methods("GET")

	r.HandleFunc("/api/v1/comment", handlers.NewWrapper(postHandler.AddCommentStratagy).ServeHTTP).Methods("POST")
	r.HandleFunc("/api/v1/comment/{id:[0-9]+}", handlers.NewWrapper(postHandler.DeleteCommentStrategy).ServeHTTP).Methods("DELETE")
	r.HandleFunc("/api/v1/comment/{id:[0-9]+}", handlers.NewWrapper(postHandler.ChangeCommentStrategy).ServeHTTP).Methods("POST")
	
	r.HandleFunc("/api/v1/analytics", handlers.NewWrapper(profileHandler.Analitycs).ServeHTTP).Methods("GET")

	r.HandleFunc("/api/v1/add_device", handlers.NewWrapper(profileHandler.GetDevice).ServeHTTP).Methods("POST")

	mltStorage := &multistorage.MultiStorage{
		Attaches: attachStorage,
		Comments: commentManager,
		Likes: likeStorage,
		Payments: payManager,
		Posts: postManager,
		SubscriptionLevels: subscriptionLevelsStorage,
		Users: userStoarge,
		Analytics: analyticsStorage,
	}
	anal.MakeCronAnalytics(mltStorage)
	err = payment.MakeCronCheckSubscriptions(payManager, subsStorage)
	log.Println(err)


	fmt.Println("Server started")
	err = http.ListenAndServe(configs.BackendServerPort, r)
	//err = http.ListenAndServeTLS(configs.BackendServerPort, "/etc/letsencrypt/live/my-kopilka.ru/fullchain.pem",
	//	"/etc/letsencrypt/live/my-kopilka.ru/privkey.pem", nil)
	//err = http.ListenAndServeTLS(configs.BackendServerPort, "cert.pem", "key.pem", nil)
	fmt.Println(err)

}