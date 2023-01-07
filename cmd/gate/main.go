package main

import (
	"github.com/ebar-go/ego/utils/runtime/signal"
	"github.com/ebar-go/znet"
	"unknown/api"
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
	server := znet.New(func(options *znet.Options) {
		options.OnOpen = handler.OpenHandler
		options.OnClose = handler.CloseHandler
	})
	options := NewOptions()

	server.Router().Route(api.ActionLogin, znet.StandardHandler(handler.LoginHandler))
	server.Router().Route(api.ActionHeartbeat, znet.StandardHandler(handler.HeartbeatHandler))

	server.ListenTCP(options.TCPAddr)
	server.ListenWebsocket(options.WebsocketAddr)

	server.Run(signal.SetupSignalHandler())
}
