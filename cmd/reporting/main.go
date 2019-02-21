package main

import (
	"encoding/json"
	"fmt"
	"github.com/Sinea/arch-async/pkg/async"
	"github.com/Sinea/arch-async/pkg/environment"
)

type UserStateChanged struct {
	User string `json:"user"`
}

func main() {
	fmt.Println("Starting REPORTING")

	rabbitUrl := environment.Get("broker_url")

	pipe, err := async.New(async.RabbitConfig{Url: rabbitUrl})

	if err != nil {
		fmt.Printf("error connecting pipes %s\n", err)
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
