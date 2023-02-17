package admin

import (
	"net/http"
	"redisadmin/internal/consts"
	"redisadmin/internal/databases"
	"redisadmin/internal/databases/goRedisAdmin"
	"redisadmin/internal/services/auth"

	"github.com/gin-gonic/gin"
)

func Login(userName string, pass string) (int, gin.H) {
	db, _ := databases.GetDb(consts.DB_RD_AD_CONF, consts.DB_RD_AD_CONF_TAG_AD)

	pass = getMd5SaltPass(pass)

	var user goRedisAdmin.User
	tx := db.Where("name = ? and pass = ?", userName, pass).First(&user)
	if tx.Error != nil {
		return http.StatusUnauthorized, gin.H{"msg": tx.Error.Error()}
	}

	// token
	token, err := auth.BuildJwtToken(user.ID, user.RoleId, user.Name)
	if err != nil {
		return http.StatusInternalServerError, gin.H{"msg": err.Error()}
	}

	return http.StatusOK, gin.H{"status": 0, "msg": "ok", "data": token}
}
