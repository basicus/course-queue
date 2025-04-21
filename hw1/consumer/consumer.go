package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	nats "github.com/nats-io/nats.go"
)

type Message struct {
	Timestamp   time.Time `json:"timestamp"`
	MessageId   string    `json:"message_id"`
	MessageText string    `json:"message_text"`
}

func main() {
	var (
		url     = flag.String("url", nats.DefaultURL, "NATS server URL")
		subject = flag.String("subject", "test.subject", "Subject to subscribe to")
	)

	flag.Parse()

	nc, err := nats.Connect(*url,
		nats.ReconnectWait(5*time.Second),
		nats.MaxReconnects(5), // Unlimited reconnections
	)
	if err != nil {
		log.Fatalf("Unable to connect: %v\n", err)
	}
	defer nc.Close()

	log.Println("Consumer connected to NATS")

	closedChan := make(chan os.Signal, 1)
	signal.Notify(closedChan, syscall.SIGINT, syscall.SIGTERM)
	nc.Subscribe(*subject, func(m *nats.Msg) {
		var msg Message
		if err := json.Unmarshal(m.Data, &msg); err != nil {
			log.Printf("Error unmarshalling message: %v\n", err)
			return
		}
		fmt.Printf("Received message: %+v\n", msg)
	})
	if err != nil {
		return
	}

	select {
	case <-closedChan:
		log.Println("Shutting down consumer...")
		return
	}
}
