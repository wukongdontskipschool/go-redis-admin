package redis_item

import (
	"context"
	"net/http"
	"redisadmin/internal/jobs"
	"redisadmin/internal/redisPool"
	"sort"

	"github.com/gin-gonic/gin"
)

func Key_list(rdType int, redisId int, rdDb int, match string) (int, gin.H) {
	conf, err := jobs.Get_redis_connect_conf_from_db(redisId)
	if err != nil {
		return http.StatusInternalServerError, gin.H{"msg": err.Error()}
	}

	if conf.RdType != rdType {
		return http.StatusForbidden, gin.H{"msg": "参数错误"}
	}

	var ctx = context.Background()
	rd := redisPool.Get_redis(conf, rdDb)
	defer rd.Close()

	len1, err := rd.DBSize(ctx).Result()
	if err != nil {
		return http.StatusInternalServerError, gin.H{"msg": err.Error()}
	}

	var cursor uint64 = 0
	var totle int64 = 10
	var val []string
	keys := make([]string, 0, len1)

	for {
		val, cursor, err = rd.Scan(ctx, cursor, match, totle).Result()

		if err != nil {
			return http.StatusInternalServerError, gin.H{"msg": err.Error()}
		}

		keys = append(keys, val...)

		if cursor == 0 {
			break
		}
	}

	sort.Strings(keys)
	return http.StatusOK, gin.H{"l": keys}
}
