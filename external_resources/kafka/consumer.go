package kafka

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/segmentio/kafka-go"
)

var (
	messageBuffer []string   // Store received messages
	mu            sync.Mutex // Ensure thread safety
)

// StartConsumer listens for Kafka events
func StartConsumer() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{os.Getenv("KAFKA_BROKER")},
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

		mu.Lock()
		messageBuffer = append(messageBuffer, string(msg.Value))
		if len(messageBuffer) > 100 { // Limit buffer size
			messageBuffer = messageBuffer[1:]
		}
		mu.Unlock()

		fmt.Printf("Received event: %s\n", string(msg.Value))
	}
}
func GetEvents(ctx context.Context) []string {
	mu.Lock()
	messages := make([]string, len(messageBuffer))
	copy(messages, messageBuffer)
	mu.Unlock()
	return messages

}
