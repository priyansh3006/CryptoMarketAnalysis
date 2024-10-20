package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	nc.Subscribe("test", func(m *nats.Msg) {
		log.Printf("Received a message: %s", string(m.Data))
	})
	log.Println("Subscribed to 'test' subject.")
	select {} // Keep the program running
}
