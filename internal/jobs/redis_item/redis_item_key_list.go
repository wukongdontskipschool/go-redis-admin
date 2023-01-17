package redis_item

import (
	"context"
	"fmt"
	"net/http"
	"redisadmin/internal/jobs"
	"redisadmin/internal/redisPool"
	"sort"

	"github.com/gin-gonic/gin"
)

func Key_list(redisId int, match string) (int, gin.H) {
	conf, err := jobs.Get_redis_connect_conf_from_db(redisId)
	if err != nil {
		return http.StatusInternalServerError, gin.H{"msg": fmt.Sprintf("%s", err)}
	}

	rd := redisPool.Get_redis(conf)

	var ctx = context.Background()

	len1, err := rd.DBSize(ctx).Result()
	if err != nil {
		return http.StatusInternalServerError, gin.H{"msg": err}
	}

	var cursor uint64 = 0
	var totle int64 = 10
	var val []string
	keys := make([]string, 0, len1)

	for {
		val, cursor, err = rd.Scan(ctx, cursor, match, totle).Result()

		if err != nil {
			return http.StatusInternalServerError, gin.H{"msg": err}
		}

		keys = append(keys, val...)

		if cursor == 0 {
			break
		}
	}

	sort.Strings(keys)
	return http.StatusOK, gin.H{"l": keys}
}
