package redisItem

import (
	"context"
	"fmt"
	"net/http"
	"redisadmin/internal/jobs"
	"redisadmin/internal/redisPool"

	"github.com/gin-gonic/gin"
)

func GetVal(typeId int, redisId int, rdDb int, key string) (int, gin.H) {
	conf, err := jobs.GetRedisConnectConfFromDb(redisId)
	if err != nil {
		return http.StatusInternalServerError, gin.H{"msg": fmt.Sprintf("%s", err)}
	}

	if conf.RdType != typeId {
		return http.StatusForbidden, gin.H{"msg": "参数错误"}
	}

	var ctx = context.Background()
	rd := redisPool.GetRedis(conf, rdDb)
	defer rd.Close()

	keyType, err := rd.Type(ctx, key).Result()
	if err != nil {
		return http.StatusInternalServerError, gin.H{"msg": fmt.Sprintf("%s", err)}
	}

	var res interface{}
	switch keyType {
	case "string":
		res, err = rd.Get(ctx, key).Result()
	case "list":
		res, err = rd.LRange(ctx, key, 0, -1).Result()
	case "zset":
		res, err = rd.ZRangeWithScores(ctx, key, 0, -1).Result()
	case "set":
		res, err = rd.SMembers(ctx, key).Result()
	case "hash":
		res, err = rd.HGetAll(ctx, key).Result()
	case "none":
		return http.StatusBadRequest, gin.H{"msg": "not exist key:" + key}
	default:
		return http.StatusInternalServerError, gin.H{"msg": "error key type:" + keyType}
	}

	if err != nil {
		return http.StatusInternalServerError, gin.H{"msg": "error type key"}
	}

	return http.StatusOK, gin.H{"data": res, "keyType": keyType}
}
