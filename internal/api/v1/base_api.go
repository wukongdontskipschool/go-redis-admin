package base_api

import (
	"github.com/gin-gonic/gin"
)

type Api_func func(*gin.Context) (int, gin.H, string)

type Base_api struct {
}

func Deal_request(api_func Api_func) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		deal(api_func, ctx)
	}
}

func deal(api_func Api_func, ctx *gin.Context) {
	http_state, gin_H, template := api_func(ctx)
	if template == "" {
		ctx.JSON(http_state, gin_H)
	} else {
		ctx.HTML(http_state, template, gin_H)
	}
}
