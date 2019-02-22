package main

import (
	"fmt"
	"github.com/Sinea/arch-async/pkg/async"
	"github.com/Sinea/arch-async/pkg/environment"
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	fmt.Println("Starting API")

	rabbitUrl, err := environment.Get("broker_url")
	if err != nil {
		panic(err)
	}

	port, err := environment.Get("port", "80")
	if err != nil {
		panic(err)
	}

	address := fmt.Sprintf(":%s", port)

	pipe, err := async.New(async.RabbitConfig{Url: rabbitUrl})

	if err != nil {
		fmt.Printf("Error creating pipe %s\n", err)
		return
	}

	fmt.Println("We're in business")

	// Create the reporting service
	reportingService := &fastReportingService{pipe}

	// Run the HTTP server
	e := echo.New()
	e.POST("/", func(c echo.Context) error {
		reportingService.ComputeStats("john")
		return c.String(http.StatusOK, "")
	})
	e.Logger.Fatal(e.Start(address))
}
