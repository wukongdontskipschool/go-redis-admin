package admin

import (
	"crypto/md5"
	"fmt"
	"log"
	"net/http"
	"redisadmin/internal/configs"
	"redisadmin/internal/consts"
	"redisadmin/internal/databases"
	"redisadmin/internal/databases/go_redis_admin"

	"github.com/gin-gonic/gin"
)

func Add_user(name string, pass string, roleId uint) (int, gin.H) {
	salt, err := getPassSalt()
	if err != nil {
		return http.StatusInternalServerError, gin.H{"msg": err.Error()}
	}

	log.Println(salt)
	pass = fmt.Sprintf("%x", md5.Sum([]byte(salt+pass)))
	user := go_redis_admin.User{Name: name, Pass: pass, RoleId: roleId}

	db, _ := databases.Get_db(consts.DB_RD_AD_CONF, consts.DB_RD_AD_CONF_TAG_AD)
	tx := db.Create(&user)
	if tx.Error != nil {
		return http.StatusInternalServerError, gin.H{"msg": tx.Error.Error()}
	}
	return http.StatusOK, gin.H{"status": 0, "msg": "ok", "data": user.ID}
}

// 密码盐
func getPassSalt() (salt string, err error) {
	var confMap map[string]string

	err = configs.Get_config(consts.ENV_CONF, &confMap)
	if err != nil {
		return
	}

	var has bool
	salt, has = confMap[consts.ENV_CONF_PSALT]
	if !has {
		salt = ""
	}

	return
}
