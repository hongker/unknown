package main

import (
	"github.com/ebar-go/ego"
	"unknown/internal/handler"
)

type Options struct {
	Addr string
}

func NewOptions() *Options {
	return &Options{
		Addr: ":9010",
	}
}

func main() {
	options := NewOptions()

	aggregator := ego.New()
	aggregator.Aggregate(ego.NewHTTPServer(options.Addr).
		EnableAvailableHealthCheck().
		RegisterRouteLoader(handler.RouteLoader))

	aggregator.Run()

}
