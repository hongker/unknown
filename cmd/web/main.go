package main

import (
	"github.com/ebar-go/ego"
	"github.com/ebar-go/ego/component"
	"github.com/ebar-go/ego/utils/runtime"
	"gorm.io/gorm"
	"unknown/internal/domain/application"
	"unknown/internal/handler"
)

type Options struct {
	Addr string
	DSN  string
	Sign []byte
}

func NewOptions() *Options {
	return &Options{
		Addr: ":9010",
		DSN:  "root:123456@tcp(127.0.0.1:3306)/game?charset=utf8mb4&parseTime=True&loc=Local",
		Sign: []byte("12345678"),
	}
}

func main() {
	options := NewOptions()

	application.SetSign(options.Sign)

	err := runtime.Call(func() error {
		return component.DB().Connect(options.DSN, &gorm.Config{})
	})

	if err != nil {
		panic(err)
	}

	aggregator := ego.New()
	aggregator.Aggregate(ego.NewHTTPServer(options.Addr).
		EnableAvailableHealthCheck().
		RegisterRouteLoader(handler.RouteLoader))

	aggregator.Run()

}
