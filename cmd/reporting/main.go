package main

import (
	"encoding/json"
	"fmt"
	"github.com/Sinea/arch-async/pkg/async"
)

type UserStateChanged struct {
	User string `json:"user"`
}

func main() {
	fmt.Println("Starting reporting service")

	pipe, err := async.New(async.RabbitConfig{Url: "amqp://guest:guest@broker:5672/"})

	if err != nil {
		fmt.Printf("error connecting pipes %s", err)
		return
	}

	fmt.Println("We're in business")

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
