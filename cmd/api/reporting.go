package main

import (
	"time"

	"github.com/Sinea/arch-async/pkg/async"
)

type UserStateChanged struct {
	User string `json:"user"`
}

type Reporting interface {
	// Expensive operation
	ComputeStats(user string) error
}

type lazyReportingService struct {
}

// This 'uses' precious CPU and Memory
func (s *lazyReportingService) ComputeStats(user string) error {
	time.Sleep(50 * time.Millisecond)
	return nil
}

type fastReportingService struct {
	pipe async.Writer
}

// This notifies your workers that the user stats needs to be computed
func (s *fastReportingService) ComputeStats(user string) error {
	return s.pipe.Write("stats_changed", UserStateChanged{user})
}
