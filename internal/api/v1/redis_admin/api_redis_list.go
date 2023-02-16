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

func (r *api_redis_list) Login(ctx *gin.Context) (int, gin.H, string) {
	return 200, gin.H{}, "login.html"
}

func (b *api_redis_list) GetAuthRule(ctx *gin.Context) string {
	return "/v1/redisList/item"
}

func (r *api_redis_list) TypeList(ctx *gin.Context) (int, gin.H, string) {
	http_state, gin_H := redis_list.TypeList()

	return http_state, gin_H, ""
}

func (r *api_redis_list) TypeStore(ctx *gin.Context) (int, gin.H, string) {
	name := ctx.PostForm("name")
	http_state, gin_H := redis_list.TypeStore(name)
	return http_state, gin_H, ""
}

func (r *api_redis_list) TypeUpdate(ctx *gin.Context) (int, gin.H, string) {
	name := ctx.PostForm("name")
	id := ctx.PostForm("id")
	http_state, gin_H := redis_list.TypeUpdate(id, name)
	return http_state, gin_H, ""
}

func (r *api_redis_list) TypeDelete(ctx *gin.Context) (int, gin.H, string) {
	id := ctx.Query("id")
	http_state, gin_H := redis_list.TypeDelete(id)
	return http_state, gin_H, ""
}

func (r *api_redis_list) ItemList(ctx *gin.Context) (int, gin.H, string) {
	http_state, gin_H := redis_list.ItemList()

	return http_state, gin_H, ""
}

func (r *api_redis_list) ItemStore(ctx *gin.Context) (int, gin.H, string) {
	desc := ctx.PostForm("desc")
	host := ctx.PostForm("host")
	port := ctx.PostForm("port")
	auth := ctx.PostForm("auth")
	menuId := ctx.PostForm("menuId")

	http_state, gin_H := redis_list.ItemStore(desc, host, port, auth, menuId)
	return http_state, gin_H, ""
}

func (r *api_redis_list) ItemUpdate(ctx *gin.Context) (int, gin.H, string) {
	desc := ctx.PostForm("desc")
	host := ctx.PostForm("host")
	port := ctx.PostForm("port")
	auth := ctx.PostForm("auth")
	menuId := ctx.PostForm("menuId")
	hasAuth := ctx.PostForm("hasAuth")
	id := ctx.PostForm("id")
	http_state, gin_H := redis_list.ItemUpdate(id, desc, host, port, auth, hasAuth, menuId)
	return http_state, gin_H, ""
}

func (r *api_redis_list) ItemDelete(ctx *gin.Context) (int, gin.H, string) {
	id := ctx.Query("id")
	http_state, gin_H := redis_list.ItemDelete(id)
	return http_state, gin_H, ""
}

func (r *api_redis_list) ItemInfo(ctx *gin.Context) (int, gin.H, string) {
	id := ctx.Query("id")
	http_state, gin_H := redis_list.ItemInfo(id)
	return http_state, gin_H, ""
}
