package main

import (
	"github.com/ebar-go/ego/utils/runtime/signal"
	"github.com/ebar-go/znet"
	"unknown/api"
	"unknown/internal/domain/service"
	"unknown/internal/handler"
)

type Options struct {
	TCPAddr       string
	WebsocketAddr string
}

func NewOptions() Options {
	return Options{
		TCPAddr:       ":9000",
		WebsocketAddr: ":9001",
	}
}

func main() {
	gateHandler := handler.NewGateHandler(service.Provider().Gate())
	server := znet.New(func(options *znet.Options) {
		options.OnOpen = gateHandler.Open
		options.OnClose = gateHandler.Close
	})
	options := NewOptions()

	server.Router().Route(api.ActionLogin, znet.StandardHandler(gateHandler.Login))
	server.Router().Route(api.ActionHeartbeat, znet.StandardHandler(gateHandler.Heartbeat))

	server.ListenTCP(options.TCPAddr)
	server.ListenWebsocket(options.WebsocketAddr)

	server.Run(signal.SetupSignalHandler())
}
