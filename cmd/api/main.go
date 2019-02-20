package main

import (
	"fmt"
	"github.com/Sinea/arch-async/pkg/async"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	fmt.Println("Starting API")

	env := os.Getenv("ENVIRONMENT")
	if len(strings.TrimSpace(env)) == 0 {
		log.Fatal("cannot start with empty ENVIRONMENT")
	} else {
		fmt.Printf("Running in envoronment '%s'\n", env)
	}

	pipe, err := async.New(async.RabbitConfig{Url: "amqp://guest:guest@broker:5672/"})

	if err != nil {
		fmt.Printf("Error creating pipe %s\n", err)
		return
	}

	fmt.Println("We're in business")

	reportingService := &fastReportingService{pipe}

	//reportingService := &lazyReportingService{}

	e := echo.New()
	e.POST("/", func(c echo.Context) error {
		reportingService.ComputeStats("john")
		return c.String(http.StatusOK, "")
	})
	e.Logger.Fatal(e.Start(":80"))
}
