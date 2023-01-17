package redis_admin

import (
	base_api "redisadmin/internal/api/v1"
	"redisadmin/internal/jobs/redis_list"

	"github.com/gin-gonic/gin"
)

var Api_redis_list = &api_redis_list{}

type api_redis_list struct {
	base_api.Base_api
}

func (r *api_redis_list) Web_index(ctx *gin.Context) (int, gin.H, string) {
	return 200, gin.H{}, "index.html"
}

func (r *api_redis_list) Web_login(ctx *gin.Context) (int, gin.H, string) {
	return 200, gin.H{}, "login.html"
}

func (r *api_redis_list) Index(ctx *gin.Context) (int, gin.H, string) {
	http_state, gin_H := redis_list.Index()

	return http_state, gin_H, ""
}

func (r *api_redis_list) Store(ctx *gin.Context) (int, gin.H, string) {
	name := ctx.PostForm("name")
	http_state, gin_H := redis_list.Store(name)
	return http_state, gin_H, ""
}

func (r *api_redis_list) Migrate(ctx *gin.Context) (int, gin.H, string) {
	redis_list.Migrate()
	return 200, gin.H{"user": "a", "value": "Info"}, ""
}
