package admin

import (
	"fmt"
	"net/http"
	"redisadmin/internal/consts"
	"redisadmin/internal/databases"
	"redisadmin/internal/databases/goRedisAdmin"

	"github.com/gin-gonic/gin"
)

func Delete(rId string) (int, gin.H) {
	db, _ := databases.GetDb(consts.DB_RD_AD_CONF, consts.DB_RD_AD_CONF_TAG_AD)

	if rId == fmt.Sprintf("%d", consts.SUP_ADMIN_ID) {
		return http.StatusForbidden, gin.H{"msg": "超级管理员不能删除"}
	}

	role := goRedisAdmin.User{}
	tx := db.First(&role, rId)
	if tx.Error != nil {
		return http.StatusInternalServerError, gin.H{"msg": tx.Error.Error()}
	}

	// 硬删除
	tx = db.Unscoped().Delete(&role)
	if tx.Error != nil {
		return http.StatusInternalServerError, gin.H{"msg": tx.Error.Error()}
	}

	return http.StatusOK, gin.H{"status": 0, "msg": "ok"}
}
