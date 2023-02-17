package redisAdmin

import (
	baseApi "redisadmin/internal/api/v1"
	"redisadmin/internal/jobs/redisList"

	"github.com/gin-gonic/gin"
)

var ApiRedisList = &apiRedisList{}

type apiRedisList struct {
	baseApi.BaseApi
}

func (r *apiRedisList) WebIndex(ctx *gin.Context) (int, gin.H, string) {
	return 200, gin.H{}, "index.html"
}

func (r *apiRedisList) WebLogin(ctx *gin.Context) (int, gin.H, string) {
	return 200, gin.H{}, "login.html"
}

func (r *apiRedisList) Login(ctx *gin.Context) (int, gin.H, string) {
	return 200, gin.H{}, "login.html"
}

func (b *apiRedisList) GetAuthRule(ctx *gin.Context) string {
	return "/v1/redisList/item"
}

func (r *apiRedisList) TypeList(ctx *gin.Context) (int, gin.H, string) {
	httpState, gin_H := redisList.TypeList()

	return httpState, gin_H, ""
}

func (r *apiRedisList) TypeStore(ctx *gin.Context) (int, gin.H, string) {
	name := ctx.PostForm("name")
	httpState, gin_H := redisList.TypeStore(name)
	return httpState, gin_H, ""
}

func (r *apiRedisList) TypeUpdate(ctx *gin.Context) (int, gin.H, string) {
	name := ctx.PostForm("name")
	id := ctx.PostForm("id")
	httpState, gin_H := redisList.TypeUpdate(id, name)
	return httpState, gin_H, ""
}

func (r *apiRedisList) TypeDelete(ctx *gin.Context) (int, gin.H, string) {
	id := ctx.Query("id")
	httpState, gin_H := redisList.TypeDelete(id)
	return httpState, gin_H, ""
}

func (r *apiRedisList) ItemList(ctx *gin.Context) (int, gin.H, string) {
	httpState, gin_H := redisList.ItemList()

	return httpState, gin_H, ""
}

func (r *apiRedisList) ItemStore(ctx *gin.Context) (int, gin.H, string) {
	desc := ctx.PostForm("desc")
	host := ctx.PostForm("host")
	port := ctx.PostForm("port")
	auth := ctx.PostForm("auth")
	menuId := ctx.PostForm("menuId")

	httpState, gin_H := redisList.ItemStore(desc, host, port, auth, menuId)
	return httpState, gin_H, ""
}

func (r *apiRedisList) ItemUpdate(ctx *gin.Context) (int, gin.H, string) {
	desc := ctx.PostForm("desc")
	host := ctx.PostForm("host")
	port := ctx.PostForm("port")
	auth := ctx.PostForm("auth")
	menuId := ctx.PostForm("menuId")
	hasAuth := ctx.PostForm("hasAuth")
	id := ctx.PostForm("id")
	httpState, gin_H := redisList.ItemUpdate(id, desc, host, port, auth, hasAuth, menuId)
	return httpState, gin_H, ""
}

func (r *apiRedisList) ItemDelete(ctx *gin.Context) (int, gin.H, string) {
	id := ctx.Query("id")
	httpState, gin_H := redisList.ItemDelete(id)
	return httpState, gin_H, ""
}

func (r *apiRedisList) ItemInfo(ctx *gin.Context) (int, gin.H, string) {
	id := ctx.Query("id")
	httpState, gin_H := redisList.ItemInfo(id)
	return httpState, gin_H, ""
}
