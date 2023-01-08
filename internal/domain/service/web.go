package service

import (
	"context"
	"github.com/ebar-go/ego/errors"
	"unknown/api"
	"unknown/internal/domain/application"
)

type WebService struct {
	userApp *application.UserApp
}

func NewWebService(userApp *application.UserApp) *WebService {
	return &WebService{
		userApp: userApp,
	}
}

func (service WebService) Login(ctx context.Context, req *api.LoginRequest) (resp *api.LoginResponse, err error) {
	user, lastErr := service.userApp.Login(ctx, req.Phone)
	if lastErr != nil {
		err = errors.Convert(lastErr).WithMessage("login user")
		return
	}

	token, lastErr := service.userApp.GenToken(ctx, user)
	if lastErr != nil {
		err = errors.Convert(lastErr).WithMessage("generate token")
		return
	}

	resp = &api.LoginResponse{
		UserID: user.ID,
		Token:  token,
	}
	return
}
