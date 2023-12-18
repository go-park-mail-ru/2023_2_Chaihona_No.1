package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/configs"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/public/notifications"
	"github.com/segmentio/kafka-go"
)

func main() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{configs.KafkaNotificationsBroker1Address},
		Topic:   configs.KafkaNotificationsTopic,
		// GroupID: "my-group",
	})

	event, err := notifications.ConsumeEvent(context.Background(), r)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(event.GetEventType())

	// data, err := event.GetMarshalled()
	// if err != nil {
	// 	log.Println(err)
	// }
	// var ev notifications.PostEvent
	// json.Unmarshal(data, &ev)
	// fmt.Println(ev)
}
