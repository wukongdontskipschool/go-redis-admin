package admin

import (
	"net/http"
	"redisadmin/internal/consts"
	"redisadmin/internal/databases"
	"redisadmin/internal/databases/go_redis_admin"

	"github.com/gin-gonic/gin"
)

func Add_rule(rule string, act string, desc string) (int, gin.H) {
	ruler := go_redis_admin.Rule{Rule: rule, Act: act, Desc: desc}

	db, _ := databases.Get_db(consts.DB_RD_AD_CONF, consts.DB_RD_AD_CONF_TAG_AD)
	tx := db.Create(&ruler)
	if tx.Error != nil {
		return http.StatusInternalServerError, gin.H{"msg": tx.Error.Error()}
	}
	return http.StatusOK, gin.H{"status": 0, "msg": "ok", "data": ruler.ID}
}
