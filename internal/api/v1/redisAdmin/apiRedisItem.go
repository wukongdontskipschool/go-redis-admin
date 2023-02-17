package redisAdmin

import (
	"log"
	"net/http"
	baseApi "redisadmin/internal/api/v1"
	"redisadmin/internal/jobs/redisItem"
	"strconv"

	"github.com/gin-gonic/gin"
)

var ApiRedisItem = apiRedisItem{}

type apiRedisItem struct {
	baseApi.BaseApi
}

func (b *apiRedisItem) GetAuthRule(ctx *gin.Context) string {
	typeId := ctx.Param("typeId")
	return "/v1/redisItem/redisList/" + typeId
}

func (r *apiRedisItem) Keys(ctx *gin.Context) (int, gin.H, string) {
	id := ctx.Param("rdId")
	typeId := ctx.Param("typeId")
	db := ctx.Query("db")
	macth := ctx.DefaultQuery("macth", "")
	confId, _ := strconv.Atoi(id)
	rdType, _ := strconv.Atoi(typeId)
	rdDb, _ := strconv.Atoi(db)

	state, hash := redisItem.KeyList(rdType, confId, rdDb, macth)
	return state, hash, ""
}

func (r *apiRedisItem) GetVal(ctx *gin.Context) (int, gin.H, string) {
	typeId := ctx.Param("typeId")
	id := ctx.Param("rdId")
	key := ctx.Query("key")
	db := ctx.Query("db")
	confId, _ := strconv.Atoi(id)
	rdType, _ := strconv.Atoi(typeId)
	rdDb, _ := strconv.Atoi(db)

	state, hash := redisItem.GetVal(rdType, confId, rdDb, key)
	return state, hash, ""
}

func (r *apiRedisItem) GetRdisList(ctx *gin.Context) (int, gin.H, string) {
	typeId := ctx.Param("typeId")
	confId, _ := strconv.Atoi(typeId)

	state, hash := redisItem.GetRdisList(confId)
	return state, hash, ""
}

func (r *apiRedisItem) IndexHtml(ctx *gin.Context) (int, gin.H, string) {
	rdType := ctx.Param("typeId")
	// confId, _ := strconv.Atoi(rdType)
	log.Println(rdType)
	return http.StatusOK, gin.H{}, "category.html"
}
