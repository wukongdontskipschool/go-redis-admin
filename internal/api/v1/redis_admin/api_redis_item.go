package redis_admin

import (
	"log"
	"net/http"
	base_api "redisadmin/internal/api/v1"
	"redisadmin/internal/jobs/redis_item"
	"strconv"

	"github.com/gin-gonic/gin"
)

var Api_redis_item = api_redis_item{}

type api_redis_item struct {
	base_api.Base_api
}

func (r *api_redis_item) Keys(ctx *gin.Context) (int, gin.H, string) {
	id := ctx.Param("id")
	macth := ctx.DefaultQuery("macth", "")
	confId, _ := strconv.Atoi(id)

	state, hash := redis_item.Key_list(confId, macth)
	return state, hash, ""
}

func (r *api_redis_item) Get_val(ctx *gin.Context) (int, gin.H, string) {
	id := ctx.Param("id")
	key := ctx.Param("key")
	confId, _ := strconv.Atoi(id)

	state, hash := redis_item.Get_val(confId, key)
	return state, hash, ""
}

func (r *api_redis_item) Get_Rdis_list(ctx *gin.Context) (int, gin.H, string) {
	typeId := ctx.Param("typeId")
	confId, _ := strconv.Atoi(typeId)

	state, hash := redis_item.Get_Rdis_list(confId)
	return state, hash, ""
}

func (r *api_redis_item) IndexHtml(ctx *gin.Context) (int, gin.H, string) {
	rdType := ctx.Param("typeId")
	// confId, _ := strconv.Atoi(rdType)
	log.Println(rdType)
	return http.StatusOK, gin.H{}, "category.html"
}
