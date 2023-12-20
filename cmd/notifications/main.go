package main

import (
	"context"
	"fmt"
	"log"
	"path/filepath"

	// "path/filepath"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/configs"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/db/postgresql"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/posts"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/public/notifications"
	"github.com/segmentio/kafka-go"
	"google.golang.org/api/option"
	// "google.golang.org/api/option"
	// "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/configs"
	// "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/db/postgresql"
	// "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/posts"
	// "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/public/notifications"
	// "github.com/segmentio/kafka-go"
)

func SetupFirebase() (*firebase.App, context.Context, *messaging.Client) {

	ctx := context.Background()

	serviceAccountKeyFilePath, err := filepath.Abs("./cmd/notifications/serviceAccountKey.json")
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

func sendToToken(app *firebase.App, token string) {
	ctx := context.Background()
	client, err := app.Messaging(ctx)
	if err != nil {
		log.Fatalf("error getting Messaging client: %v\n", err)
	}

	// registrationToken := "d3adVu0tMDyBUTPkgc_l-0:APA91bHjb6-wWkT1ABGSasFqxrsOR3AdfcTjLc8b7f7yukWLt32GS4UA5XdIwZ8p98oOLp-CBcyuYaCYdEPRji_f2WSXO9JKb7XPjotm_3bdkk-7hJyxJS8JuUHt82xzGGJ6Aacy0QWb"

	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: "Новый пост в копилке!",
			Body:  "Скорее бежим смотреть",
		},
		Token: token,
	}

	response, err := client.Send(ctx, message)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Successfully sent message:", response)
}

func main() {
	app, _, _ := SetupFirebase()
	// sendToToken(app, "dB_ArJ7Qm2g:APA91bGjy-EkmFbHoDf80Fn4PBG__urkqCfEmMpyhAhCWUqbWeGYracuM3dbn8vFb0xypcHuDpXARnLu5KSk380R8o5-UQMERAFjvpSNqO3w0pArtzAtslpdtzdUBEebqX7mmAWcaxcP")
	var db postgresql.Database
	err := db.Connect()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	postStorage := posts.CreatePostStorage(db.GetDB())
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{configs.KafkaNotificationsBroker1Address},
		// Brokers: []string{"localhost:2181"},
		Topic: configs.KafkaNotificationsTopic,
		// GroupID: "my-group",
	})

	for {
		event, err := notifications.ConsumeEvent(context.Background(), r)
		if err != nil {
			log.Println(err)
		}
		postId, ok := event.Body["id"].(float64)
		// fmt.Println(data)
		// if !ok {
		// 	continue
		// } else {
		// 	fmt.Println("error data types")
		// }
		// id, ok := data.(int)
		if ok {
			ids, err := postStorage.GetDevicesID(int(postId))
			for _, id := range ids {
				go func() {
					sendToToken(app, id.DeviceId)
				}()
			}
			sendToToken(app, ids[0].DeviceId)
			fmt.Println("ids: ", ids, err)
		} else {
			fmt.Println(postId)
			fmt.Println(event.Body["id"])
			fmt.Println("error convert data")
		}
	}
}

// data, err := event.GetMarshalled()
// if err != nil {
// 	log.Println(err)
// }
// var ev notifications.PostEvent
// json.Unmarshal(data, &ev)
// fmt.Println(ev)
