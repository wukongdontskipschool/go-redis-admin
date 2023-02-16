package admin

import (
	"net/http"
	"redisadmin/internal/consts"
	"redisadmin/internal/databases"
	"redisadmin/internal/databases/go_redis_admin"

	"github.com/gin-gonic/gin"
)

func Get_user_list() (int, gin.H) {
	db, _ := databases.Get_db(consts.DB_RD_AD_CONF, consts.DB_RD_AD_CONF_TAG_AD)
	var roles []go_redis_admin.User
	if res := db.Find(&roles); res.Error != nil {
		return http.StatusInternalServerError, gin.H{"msg": res.Error.Error()}
	}

	type res struct {
		ID    uint
		Name  string
		Rname string
	}
	var user []res
	db.Model(&go_redis_admin.User{}).Select("users.id, users.name, roles.name as rname").Joins("left join roles on roles.id = users.role_id").Scan(&user)

	return http.StatusOK, gin.H{"status": 0, "msg": "ok", "data": user}
}
