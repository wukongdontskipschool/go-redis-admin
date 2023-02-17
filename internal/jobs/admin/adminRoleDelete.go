package admin

import (
	"net/http"
	"redisadmin/internal/accessControl"
	"redisadmin/internal/consts"
	"redisadmin/internal/databases"
	"redisadmin/internal/databases/goRedisAdmin"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RoleDelete(rId string) (int, gin.H) {
	db, _ := databases.GetDb(consts.DB_RD_AD_CONF, consts.DB_RD_AD_CONF_TAG_AD)

	role := goRedisAdmin.Role{}
	tx := db.First(&role, rId)
	if tx.Error != nil {
		return http.StatusInternalServerError, gin.H{"msg": tx.Error.Error()}
	}

	rIdInt, _ := strconv.Atoi(rId)

	user := goRedisAdmin.User{}
	tx = db.Where("role_id = ?", rIdInt).Take(&user)
	if tx.RowsAffected > 0 {
		return http.StatusInternalServerError, gin.H{"msg": "存在拥有该角色的用户"}
	}

	// 硬删除
	tx = db.Unscoped().Delete(&role)
	if tx.Error != nil {
		return http.StatusInternalServerError, gin.H{"msg": tx.Error.Error()}
	}

	// 清空权限 图便捷
	if _, err := accessControl.DeletePermissionsForUser(rId); err != nil {
		return http.StatusInternalServerError, gin.H{"msg": err.Error()}
	}

	return http.StatusOK, gin.H{"status": 0, "msg": "ok", "data": user.ID}
}
