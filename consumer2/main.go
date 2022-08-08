package main

import (
	"context"
	"fmt"
	"log"

	kafka "github.com/segmentio/kafka-go"
)

func main() {
	topic := "example-123"
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		GroupID: "cons-grp-1",
		Topic:   topic,
	})

	for {
		m, err := r.ReadMessage(context.Background())

		if err != nil {
			break
		}

		fmt.Printf("message received: \n\tKey: %v, Val: %v, Offset: %v\n", string(m.Key), string(m.Value), m.Offset)
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader: ", err)
	}
}
