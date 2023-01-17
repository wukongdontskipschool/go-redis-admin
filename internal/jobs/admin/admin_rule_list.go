package admin

import (
	"net/http"
	"redisadmin/internal/consts"
	"redisadmin/internal/databases"
	"redisadmin/internal/databases/go_redis_admin"

	"github.com/gin-gonic/gin"
)

func Get_rule_list() (int, gin.H) {
	db, _ := databases.Get_db(consts.DB_RD_AD_CONF, consts.DB_RD_AD_CONF_TAG_AD)
	var roles []go_redis_admin.Rule
	if res := db.Find(&roles); res.Error != nil {
		return http.StatusInternalServerError, gin.H{"msg": res.Error.Error()}
	}

	return http.StatusOK, gin.H{"status": 0, "msg": "ok", "data": roles}
}
