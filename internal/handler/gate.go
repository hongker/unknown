package handler

import (
	"github.com/ebar-go/znet"
	"unknown/api"
	"unknown/internal/domain/service"
)

type GateHandler struct {
	gateService *service.GateService
}

func NewGateHandler(gateService *service.GateService) *GateHandler {
	return &GateHandler{
		gateService: gateService,
	}
}

func (handler *GateHandler) Login(ctx *znet.Context, req *api.GateLoginRequest) (resp *api.GateLoginResponse, err error) {
	return handler.gateService.Login(ctx, req)
}

func (handler *GateHandler) Heartbeat(ctx *znet.Context, req *api.GateHeartbeatRequest) (resp *api.GateHeartbeatResponse, err error) {
	return
}

func (handler *GateHandler) Open(connection *znet.Connection) {

}
func (handler *GateHandler) Close(connection *znet.Connection) {

}
