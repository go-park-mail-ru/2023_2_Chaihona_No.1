package notifications

import (
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"
)

func ProduceNotification(ctx context.Context, writer *kafka.Writer, event Event) error {
	msg, err := json.Marshal(event)
	if err != nil {
		return err
	}

	err = writer.WriteMessages(ctx, kafka.Message{
		Value: []byte(msg),
	})

	return err
	// w := kafka.NewWriter(kafka.WriterConfig{
	// 	Brokers: []string{broker1Address, broker2Address, broker3Address},
	// 	Topic:   topic,
	// })
}
