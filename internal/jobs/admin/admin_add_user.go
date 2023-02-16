package admin

import (
	"net/http"
	"redisadmin/internal/consts"
	"redisadmin/internal/databases"
	"redisadmin/internal/databases/go_redis_admin"

	"github.com/gin-gonic/gin"
)

func Add_user(name string, pass string, roleId uint) (int, gin.H) {
	pass = getMd5SaltPass(pass)
	user := go_redis_admin.User{Name: name, Pass: pass, RoleId: roleId}

	db, _ := databases.Get_db(consts.DB_RD_AD_CONF, consts.DB_RD_AD_CONF_TAG_AD)
	tx := db.Create(&user)
	if tx.Error != nil {
		return http.StatusInternalServerError, gin.H{"msg": tx.Error.Error()}
	}
	return http.StatusOK, gin.H{"status": 0, "msg": "ok", "data": user.ID}
}
