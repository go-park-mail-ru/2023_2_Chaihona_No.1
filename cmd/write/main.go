package main

import (
	"context"
	"fmt"

	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/configs"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/public/notifications"
	"github.com/segmentio/kafka-go"
)

func main() {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{configs.KafkaNotificationsBroker1Address, configs.KafkaNotificationsBroker3Address},
		Topic:   configs.KafkaNotificationsTopic,
	})
	if w == nil {
		fmt.Println("nil")
	}
	// fmt.Println(w)
	err := notifications.ProduceNotification(context.Background(), w, notifications.PostEvent{1})
	fmt.Println(err)
}
