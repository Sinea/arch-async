package main

import (
	"encoding/json"
	"fmt"
	"github.com/Sinea/arch-async/pkg/async"
	"log"
	"os"
	"strings"
)

type UserStateChanged struct {
	User string `json:"user"`
}

func main() {
	fmt.Println("Starting REPORTING")

	env := os.Getenv("ENVIRONMENT")
	if len(strings.TrimSpace(env)) == 0 {
		log.Fatal("cannot start with empty ENVIRONMENT")
	} else {
		fmt.Printf("Running in envoronment '%s'\n", env)
	}

	pipe, err := async.New(async.RabbitConfig{Url: "amqp://guest:guest@broker:5672/"})

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
