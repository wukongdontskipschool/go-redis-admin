package admin

import (
	"net/http"
	"redisadmin/internal/consts"
	"redisadmin/internal/databases"
	"redisadmin/internal/databases/go_redis_admin"

	"github.com/gin-gonic/gin"
)

func UserInfo(rId string) (int, gin.H) {
	db, _ := databases.Get_db(consts.DB_RD_AD_CONF, consts.DB_RD_AD_CONF_TAG_AD)

	var role go_redis_admin.User
	tx := db.First(&role, rId)
	if tx.Error != nil {
		return http.StatusInternalServerError, gin.H{"msg": tx.Error.Error()}
	}

	return http.StatusOK, gin.H{"status": 0, "msg": "ok", "data": role}
}
