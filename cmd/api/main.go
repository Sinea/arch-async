package main

import (
	"arch-async/pkg/async"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	fmt.Println("Starting API")

	pipe, _ := async.New(async.RabbitConfig{Url: "amqp://guest:guest@localhost:5672/"})
	reportingService := &fastReportingService{pipe}

	//reportingService := &lazyReportingService{}

	e := echo.New()
	e.POST("/", func(c echo.Context) error {
		reportingService.ComputeStats("john")
		return c.String(http.StatusOK, "")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
