package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/google/uuid"
	nats "github.com/nats-io/nats.go"
)

type Message struct {
	Timestamp   time.Time `json:"timestamp"`
	MessageId   string    `json:"message_id"`
	MessageText string    `json:"message_text"`
}

const (
	defaultLetters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

var (
	letterIdx uint32
)

// randtext generates a random string of length n.
func randtext(n int) string {
	b := make([]byte, n)
	for i := range b {
		idx := atomic.AddUint32(&letterIdx, 1)
		if idx >= uint32(len(defaultLetters)) {
			atomic.StoreUint32(&letterIdx, 0)
			idx = atomic.AddUint32(&letterIdx, 1)
		}
		b[i] = defaultLetters[idx%uint32(len(defaultLetters))]
	}
	return string(b)
}

func main() {
	var (
		url     = flag.String("url", nats.DefaultURL, "NATS server URL")
		subject = flag.String("subject", "test.subject", "Subject to publish messages to")
	)

	flag.Parse()

	nc, err := nats.Connect(*url,
		nats.ReconnectWait(5*time.Second),
		nats.MaxReconnects(5),
	)
	if err != nil {
		log.Fatalf("Unable to connect: %v\n", err)
	}
	defer nc.Close()

	log.Println("Producer connected to NATS")

	closedChan := make(chan os.Signal, 1)
	signal.Notify(closedChan, syscall.SIGINT, syscall.SIGTERM)

	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ticker.C:
			msg := Message{
				Timestamp:   time.Now(),
				MessageId:   uuid.New().String(),
				MessageText: randtext(20),
			}

			jsonMsg, err := json.Marshal(msg)
			if err != nil {
				log.Printf("Error marshalling message: %v\n", err)
				continue
			}

			if err = nc.Publish(*subject, jsonMsg); err != nil {
				log.Printf("Error publishing message: %v\n", err)
			} else {
				log.Printf("Published message: %s\n", msg.MessageText)
			}
		case <-closedChan:
			ticker.Stop()
			log.Println("Shutting down producer...")
			return
		}
	}
}
