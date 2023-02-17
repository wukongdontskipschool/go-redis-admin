package admin

import (
	"net/http"
	"redisadmin/internal/consts"
	"redisadmin/internal/databases"
	"redisadmin/internal/databases/goRedisAdmin"

	"github.com/gin-gonic/gin"
)

func GetRuleList() (int, gin.H) {
	db, _ := databases.GetDb(consts.DB_RD_AD_CONF, consts.DB_RD_AD_CONF_TAG_AD)
	var roles []goRedisAdmin.Rule
	if res := db.Find(&roles); res.Error != nil {
		return http.StatusInternalServerError, gin.H{"msg": res.Error.Error()}
	}

	return http.StatusOK, gin.H{"status": 0, "msg": "ok", "data": roles}
}
