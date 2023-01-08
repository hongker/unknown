package handler

import (
	"context"
	"github.com/ebar-go/ego/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"unknown/api"
	"unknown/internal/domain/service"
)

func RouteLoader(router *gin.Engine) {
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})

	handler := NewWebHandler(service.Provider().Web())

	apiGroup := router.Group("api")
	{
		apiGroup.POST("login", WrapHandler(handler.Login))
	}
}

type APIResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func WrapHandler[Request, Response any](fn func(ctx context.Context, req *Request) (resp *Response, err error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := new(Request)
		if err := ctx.ShouldBindJSON(req); err != nil {
			ce := errors.InvalidParam("invalid params: %v", err)
			ctx.JSON(http.StatusOK, &APIResult{
				Code:    ce.Code(),
				Message: ce.Message(),
			})
			return
		}

		resp, err := fn(ctx, req)

		result := &APIResult{
			Code:    0,
			Message: "",
			Data:    nil,
		}
		if err != nil {
			ce := errors.Convert(err)
			result.Code = ce.Code()
			result.Message = ce.Message()
		} else {
			result.Data = resp
		}

		ctx.JSON(http.StatusOK, result)
	}
}

type WebHandler struct {
	webService *service.WebService
}

func NewWebHandler(webService *service.WebService) *WebHandler {
	return &WebHandler{
		webService: webService,
	}
}
func (handler WebHandler) Login(ctx context.Context, req *api.WebLoginRequest) (resp *api.WebLoginResponse, err error) {
	return handler.webService.Login(ctx, req)
}
