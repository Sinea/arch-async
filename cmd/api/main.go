package main

import (
	"fmt"
	"github.com/Sinea/arch-async/pkg/async"
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	fmt.Println("Starting API")

	pipe, err := async.New(async.RabbitConfig{Url: "amqp://guest:guest@broker:5672/"})

	if err != nil {
		fmt.Printf("Error creating pipe %s", err)
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
	e.Logger.Fatal(e.Start(":1323"))
}
