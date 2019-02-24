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

	rabbitUrl, err := environment.Get("broker_url")
	if err != nil {
		log.Fatal(err)
	}

	pipe, err := async.New(async.RabbitConfig{Url: rabbitUrl})
	if err != nil {
		log.Fatalf("error connecting pipes %s\n", err)
	}

	fmt.Println("We're in business")

	for m := range pipe.Read() {
		switch m.Tag {
		case "stats_changed":
			event := UserStateChanged{}
			if json.Unmarshal(m.Payload, &event) == nil {
				fmt.Printf("Computing for '%s'\n", event.User)
			}
		default:
			fmt.Println("unknown message type")
		}
	}
}
