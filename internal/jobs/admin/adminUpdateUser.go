package admin

import (
	"net/http"
	"redisadmin/internal/consts"
	"redisadmin/internal/databases"
	"redisadmin/internal/databases/goRedisAdmin"
	"redisadmin/internal/services/auth"

	"github.com/gin-gonic/gin"
)

func UpdateUser(uId uint, name string, pass string, roleId uint, jwt *auth.JwtClaims) (int, gin.H) {
	user := &goRedisAdmin.User{}

	db, _ := databases.GetDb(consts.DB_RD_AD_CONF, consts.DB_RD_AD_CONF_TAG_AD)
	tx := db.First(user, uId)
	if tx.Error != nil {
		return http.StatusInternalServerError, gin.H{"msg": tx.Error.Error()}
	}

	if uId == consts.SUP_ADMIN_ID && jwt.UId != uId {
		return http.StatusForbidden, gin.H{"msg": "没有权限改超级用户的"}
	}

	if uId == consts.SUP_ADMIN_ID && roleId != consts.SUP_ADMIN_RID {
		return http.StatusForbidden, gin.H{"msg": "超级用户不能修改角色"}
	}

	if user.RoleId == consts.SUP_ADMIN_RID && jwt.UId != uId && consts.SUP_ADMIN_ID != jwt.UId {
		return http.StatusForbidden, gin.H{"msg": "没有权限改其他超级管理员的"}
	}

	update := false
	filed1 := make(map[string]interface{})
	if roleId != 0 {
		tx = db.Find(&goRedisAdmin.Role{}, roleId)
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
