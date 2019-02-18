package main

import (
	"github.com/Sinea/arch-async/pkg/async"
	"log"
	"time"
)

type UserStateChanged struct {
	User string `json:"user"`
}

type Reporting interface {
	// Expensive operation
	ComputeStats(user string)
}

type lazyReportingService struct {
}

// This 'uses' precious CPU and Memory
func (s *lazyReportingService) ComputeStats(user string) {
	time.Sleep(50 * time.Millisecond)
}

type fastReportingService struct {
	pipe async.Writer
}

// This notifies your workers that the user stats needs to be computed
func (s *fastReportingService) ComputeStats(user string) {
	if err := s.pipe.Write("stats_changed", UserStateChanged{user}); err != nil {
		log.Fatal(err)
	}
}
