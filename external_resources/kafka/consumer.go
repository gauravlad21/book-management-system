package kafka

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

// StartConsumer listens for Kafka events
func StartConsumer() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		GroupID:  "book_consumer_group",
		Topic:    KafkaTopic,
		MinBytes: 10e3,
		MaxBytes: 10e6,
	})

	log.Println("Kafka consumer started...")

	for {
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Error reading message: %v\n", err)
			continue
		}

		fmt.Printf("Received event: %s\n", string(msg.Value))
	}
}
