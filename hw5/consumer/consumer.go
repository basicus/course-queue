package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"hw5/models"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	nats "github.com/nats-io/nats.go"
)

func main() {
	var (
		url     = flag.String("url", nats.DefaultURL, "NATS server URL")
		subject = flag.String("subject", "events", "Subject to subscribe to")
	)

	flag.Parse()

	nc, err := nats.Connect(*url,
		nats.ReconnectWait(5*time.Second),
		nats.MaxReconnects(5), // Unlimited reconnections
		nats.UserInfo("admin", "pwd"),
	)
	if err != nil {
		log.Fatalf("Unable to connect: %v\n", err)
	}
	defer nc.Close()
	defer func(nc *nats.Conn) {
		err := nc.Drain()
		if err != nil {
			log.Fatalf("Unable to drain connection: %v\n", err)
		}
	}(nc)

	log.Println("Consumer connected to NATS ", nc.ConnectedClusterName())

	closedChan := make(chan os.Signal, 1)
	signal.Notify(closedChan, syscall.SIGINT, syscall.SIGTERM)
	sub, _ := nc.Subscribe(*subject, func(msg *nats.Msg) {
		var m models.Message
		if err := json.Unmarshal(msg.Data, &m); err != nil {
			log.Printf("Error unmarshalling message: %v\n", err)
			return
		}
		fmt.Printf("Received message: %+v\n", msg)
		resp := models.MessageReply{
			MessageId:     m.MessageId,
			MessageStatus: 200,
			MessageReply:  "OK from " + nc.ConnectedClusterName(),
		}
		rbytes, _ := json.Marshal(resp)
		err := msg.Respond(rbytes)
		if err != nil {
			log.Printf("Error unmarshalling message: %v\n", err)
		}
	})
	defer sub.Unsubscribe()
	if err != nil {
		return
	}

	select {
	case <-closedChan:
		log.Println("Shutting down consumer...")
		return
	}
}
