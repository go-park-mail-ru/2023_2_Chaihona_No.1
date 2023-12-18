package main

import (
	"context"
	"log"

	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/configs"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/public/notifications"
	"github.com/segmentio/kafka-go"
)

func main() {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{configs.KafkaNotificationsBroker1Address},
		Topic:   configs.KafkaNotificationsTopic,
	})
	// fmt.Println(w)
	err := notifications.ProduceNotification(context.Background(), w, notifications.Event{1, map[string]any{"id": 1}})
	log.Println("Error writer:", err)
}
