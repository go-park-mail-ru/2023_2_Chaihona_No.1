package notifications

import (
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
)

func ConsumeEvent(ctx context.Context, reader *kafka.Reader) (Event, error) {
	// r := kafka.NewReader(kafka.ReaderConfig{
	// 	Brokers: []string{broker1Address, broker2Address, broker3Address},
	// 	Topic:   topic,
	// 	GroupID: "my-group",
	// })
	msg, err := reader.ReadMessage(ctx)
	if err != nil {
		return Event{}, err
	}

	var event Event
	err = kafka.Unmarshal(msg.Value, &event)
	log.Println(event)
	if err != nil {
		return Event{}, err
	}

	return event, nil
}
