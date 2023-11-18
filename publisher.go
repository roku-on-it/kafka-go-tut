package main

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log"
	"os"
)

func main() {
	message := os.Args[1]
	ctx := context.Background()

	partition := 0

	// Partition 0 A-M, Partition 1 N-Z
	if message[0] > 'm' {
		partition = 1
	}

	conn, err := kafka.DialLeader(
		ctx,
		"tcp",
		"localhost:29092",
		"my-topic",
		partition,
	)

	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	_, err = conn.WriteMessages(kafka.Message{
		Value: []byte(message),
	})

	if err != nil {
		log.Fatal("failed to write messages:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close writer:", err)
	}
}
