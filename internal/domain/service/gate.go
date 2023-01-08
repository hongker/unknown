package service

import (
	"context"
	"github.com/ebar-go/ego/errors"
	"time"
	"unknown/api"
	"unknown/internal/domain/application"
)

type GateService struct {
	userApp *application.UserApp
}

func NewGateService(userApp *application.UserApp) *GateService {
	return &GateService{
		userApp: userApp,
	}
}

func (service *GateService) Login(ctx context.Context, req *api.GateLoginRequest) (resp *api.GateLoginResponse, err error) {
	_, lastErr := service.userApp.Auth(ctx, req.Token)
	if lastErr != nil {
		err = errors.Convert(lastErr).WithMessage("auth token")
		return
	}

	resp = &api.GateLoginResponse{ServerTime: time.Now().Unix()}

	return

}
