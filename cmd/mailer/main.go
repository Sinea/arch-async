package main

import (
	"arch-async/pkg/async"
	"encoding/json"
	"fmt"
)

type UserStateChanged struct {
	User string `json:"user"`
}

func main() {
	fmt.Println("Starting reporting service")

	pipe, _ := async.New(async.RabbitConfig{Url: "amqp://guest:guest@localhost:5672/"})

	for m := range pipe.Read() {
		switch m.Kind {
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
