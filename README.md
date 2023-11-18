# Kafka - Partitioned Producer-Consumer
This project demonstrates a simple implementation of a Kafka producer and two consumers in Go. The producer publishes names to a Kafka topic, and the consumers are assigned partitions based on the initial letter of the names (I assume Kafka does that automatically). Names starting with A to N go to partition 0, and names starting from M to Z go to partition 1.

`publisher.go` contains the implementation of the Kafka producer.

```go
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

```
`consumer.go` file contains the implementation of the Kafka consumers.

```go
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

```

### How to Run
Ensure you have a running Kafka broker on `localhost:29092`.

Compile and run the `publisher.go` file to push names to the Kafka topic.

```bash
go run publisher.go <name>
```

Compile and run the `consumer.go` file to read messages from the Kafka topic.

``` bash
go run consumer.go
```

### Dependencies

[segmentio/kafka-go](https://github.com/segmentio/kafka-go): A Kafka library in Go.
