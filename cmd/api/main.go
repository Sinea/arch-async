package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Sinea/arch-async/pkg/async"
	"github.com/Sinea/arch-async/pkg/environment"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	log.Println("Starting API")

	rabbitURL, err := environment.Get("broker_url")
	if err != nil {
		log.Fatal(err)
	}

	port, err := environment.Get("port", "80")
	if err != nil {
		log.Fatal(err)
	}

	address := fmt.Sprintf(":%s", port)

	pipe, err := async.New(async.RabbitConfig{URL: rabbitURL})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("We're in business")

	// Create the reporting service
	reportingService := &fastReportingService{pipe}

	// Run the HTTP server
	e := echo.New()
	e.Use(middleware.Recover(), middleware.Logger())

	e.POST("/", func(c echo.Context) error {
		if err := reportingService.ComputeStats("john"); err != nil {
			return err
		}
		return c.NoContent(http.StatusAccepted)
	})
	e.Logger.Fatal(e.Start(address))
}
