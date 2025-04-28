package main

import (
	"encoding/json"
	"flag"
	"hw5/models"
	"log"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/google/uuid"
	nats "github.com/nats-io/nats.go"
)

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
		subject = flag.String("subject", "events", "Subject to subscribe to")
	)

	flag.Parse()

	nc, err := nats.Connect(*url,
		nats.ReconnectWait(5*time.Second),
		nats.MaxReconnects(5),
		nats.UserInfo("admin", "pwd"),
	)
	if err != nil {
		log.Fatalf("Unable to connect: %v\n", err)
	}
	defer nc.Close()

	log.Println("Producer connected to NATS ", nc.ConnectedClusterName())

	closedChan := make(chan os.Signal, 1)
	signal.Notify(closedChan, syscall.SIGINT, syscall.SIGTERM)

	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ticker.C:
			msg := models.Message{
				Timestamp:   time.Now(),
				MessageId:   uuid.New().String(),
				MessageText: randtext(20),
			}

			jsonMsg, err := json.Marshal(msg)
			if err != nil {
				log.Printf("Error marshalling message: %v\n", err)
				continue
			}

			reply, err := nc.Request(*subject, jsonMsg, 1*time.Second)
			if err != nil {
				log.Printf("Error publishing message: %v\n", err)
			} else {
				var r models.MessageReply
				if err := json.Unmarshal(reply.Data, &r); err != nil {
					log.Printf("Error unmarshalling reply message: %v\n", err)
					return
				}
				log.Printf(">>>[%s] Published message: %+v\n", nc.ConnectedClusterName(), msg)
				log.Printf("<<<Reply message: %+v\n", r)

			}
		case <-closedChan:
			ticker.Stop()
			log.Println("Shutting down producer...")
			return
		}
	}
}
