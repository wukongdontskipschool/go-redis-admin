package admin

import (
	"net/http"
	"redisadmin/internal/consts"
	"redisadmin/internal/databases"
	"redisadmin/internal/databases/goRedisAdmin"

	"github.com/gin-gonic/gin"
)

func GetUserList() (int, gin.H) {
	db, _ := databases.GetDb(consts.DB_RD_AD_CONF, consts.DB_RD_AD_CONF_TAG_AD)
	var roles []goRedisAdmin.User
	if res := db.Find(&roles); res.Error != nil {
		return http.StatusInternalServerError, gin.H{"msg": res.Error.Error()}
	}

	type res struct {
		ID    uint
		Name  string
		Rname string
	}
	var user []res
	db.Model(&goRedisAdmin.User{}).Select("users.id, users.name, roles.name as rname").Joins("left join roles on roles.id = users.role_id").Scan(&user)

	return http.StatusOK, gin.H{"status": 0, "msg": "ok", "data": user}
}
