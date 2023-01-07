package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RouteLoader(router *gin.Engine) {
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})
}
