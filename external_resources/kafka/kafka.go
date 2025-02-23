package kafka

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/segmentio/kafka-go"
)

// KafkaTopic defines the topic name
const KafkaTopic = "book_events"

// KafkaWriter is the producer instance
var KafkaWriter *kafka.Writer

// BookEvent represents the event structure
type BookEvent struct {
	Action string `json:"action"`
	BookID string `json:"book_id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   int    `json:"year"`
	Time   string `json:"time"`
}

// Initialize Kafka Producer
func InitKafkaProducer() {
	KafkaWriter = &kafka.Writer{
		Addr:     kafka.TCP(os.Getenv("KAFKA_BROKER")),
		Topic:    KafkaTopic,
		Balancer: &kafka.LeastBytes{},
	}
}

// PublishEvent sends a book-related event to Kafka
func PublishEvent(action, bookID, title, author string, year int) {
	event := BookEvent{
		Action: action,
		BookID: bookID,
		Title:  title,
		Author: author,
		Year:   year,
		Time:   time.Now().Format(time.RFC3339),
	}

	eventJSON, _ := json.Marshal(event)
	err := KafkaWriter.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(bookID),
			Value: eventJSON,
		},
	)

	if err != nil {
		log.Printf("Failed to publish event: %v\n", err)
	} else {
		log.Printf("Event published: %s\n", eventJSON)
	}
}
