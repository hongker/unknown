package handler

import "github.com/ebar-go/znet"

type LoginRequest struct{}
type LoginResponse struct{}

func LoginHandler(ctx *znet.Context, req *LoginRequest) (resp *LoginResponse, err error) {
	return
}

type HeartbeatRequest struct{}
type HeartbeatResponse struct{}

func HeartbeatHandler(ctx *znet.Context, req *HeartbeatRequest) (resp *HeartbeatResponse, err error) {
	return
}

func OpenHandler(connection *znet.Connection) {

}
func CloseHandler(connection *znet.Connection) {

}