package admin

import (
	"net/http"
	"redisadmin/internal/consts"
	"redisadmin/internal/databases"
	"redisadmin/internal/databases/go_redis_admin"

	"github.com/gin-gonic/gin"
)

func UpdateUser(uId uint, name string, pass string, roleId uint) (int, gin.H) {
	user := &go_redis_admin.User{}

	db, _ := databases.Get_db(consts.DB_RD_AD_CONF, consts.DB_RD_AD_CONF_TAG_AD)
	tx := db.First(user, uId)
	if tx.Error != nil {
		return http.StatusInternalServerError, gin.H{"msg": tx.Error.Error()}
	}

	update := false
	filed1 := make(map[string]interface{})
	if roleId != 0 {
		tx = db.Find(&go_redis_admin.Role{}, roleId)
		if tx.Error != nil {
			return http.StatusInternalServerError, gin.H{"msg": tx.Error.Error()}
		}
	}

	if roleId != user.RoleId {
		filed1["role_id"] = roleId
		update = true
	}

	if pass != "" {
		pass = getMd5SaltPass(pass)
		filed1["pass"] = pass
		update = true
	}

	if update {
		tx = db.Model(&user).Updates(filed1)
		if tx.Error != nil {
			return http.StatusInternalServerError, gin.H{"msg": tx.Error.Error()}
		}
	}

	return http.StatusOK, gin.H{"status": 0, "msg": "ok"}
}
