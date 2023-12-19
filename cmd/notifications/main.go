package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/configs"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/db/postgresql"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/posts"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/public/notifications"
	"github.com/segmentio/kafka-go"
)

func main() {
	// var db postgresql.Database
	// err := db.Connect()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer db.Close()
	// postStorage := posts.CreatePostStorage(db)
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{configs.KafkaNotificationsBroker1Address},
		Topic:   configs.KafkaNotificationsTopic,
		GroupID: "my-group",
	})

	for {
		event, err := notifications.ConsumeEvent(context.Background(), r)
		if err != nil {
			log.Println(err)
		}
		// data, ok := event.Body["id"]
		// if ok {
		// 	ids, err := postStorage.GetDevicesID()
		// 	fmt.Println("ids: ", ids, err)
		// }
		fmt.Println(event.Body["id"])
	}
}

// data, err := event.GetMarshalled()
// if err != nil {
// 	log.Println(err)
// }
// var ev notifications.PostEvent
// json.Unmarshal(data, &ev)
// fmt.Println(ev)
