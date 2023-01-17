package redis_admin

import (
	base_api "redisadmin/internal/api/v1"
	"redisadmin/internal/jobs/menu_list"

	"github.com/gin-gonic/gin"
)

var Api_menu_list = &api_menu_list{}

type api_menu_list struct {
	base_api.Base_api
}

func (r *api_menu_list) Index(ctx *gin.Context) (int, gin.H, string) {
	var uId uint = 5
	http_state, gin_H := menu_list.Get_list(uId)

	return http_state, gin_H, ""
}
