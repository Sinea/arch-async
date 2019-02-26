package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Sinea/arch-async/pkg/async"
	"github.com/Sinea/arch-async/pkg/environment"
)

type UserStateChanged struct {
	User string `json:"user"`
}

func main() {
	fmt.Println("Starting REPORTING")

	rabbitURL, err := environment.Get("broker_url")
	if err != nil {
		log.Fatal(err)
	}

	pipe, err := async.New(async.RabbitConfig{URL: rabbitURL})
	if err != nil {
		log.Fatalf("error connecting pipes %s\n", err)
	}

	fmt.Println("We're in business")
	messages, errors := pipe.Read()

Loop:
	for {
		select {
		case message := <-messages:
			switch message.Tag {
			case "stats_changed":
				event := UserStateChanged{}
				if json.Unmarshal(message.Payload, &event) == nil {
					fmt.Printf("Computing for '%s'\n", event.User)
				}
			default:
				fmt.Println("unknown message type")
			}
		case err := <-errors:
			log.Fatal(err)
			break Loop
		}
	}
}
