package redisAdmin

import (
	baseApi "redisadmin/internal/api/v1"
	"redisadmin/internal/jobs/menuList"

	"github.com/gin-gonic/gin"
)

var ApiMenuList = &apiMenuList{}

type apiMenuList struct {
	baseApi.BaseApi
}

func (r *apiMenuList) Index(ctx *gin.Context) (int, gin.H, string) {
	var uId uint = 1
	httpState, gin_H := menuList.GetList(uId)

	return httpState, gin_H, ""
}
