package main

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"log"
)

func main() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:29092"},
		Topic:   "my-topic",
		GroupID: "my-group",
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Fatal("failed to read message:", err)
		}

		fmt.Println("msg:", string(m.Value), "partition:", m.Partition)
	}
}
