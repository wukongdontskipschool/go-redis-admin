package redisItem

import (
	"net/http"
	"redisadmin/internal/consts"
	"redisadmin/internal/databases"
	"redisadmin/internal/databases/goRedisAdmin"

	"github.com/gin-gonic/gin"
)

func GetRdisList(redisId int) (int, gin.H) {
	db, err := databases.GetDb(consts.DB_RD_AD_CONF, consts.DB_RD_AD_CONF_TAG_AD)
	if err != nil {
		return http.StatusInternalServerError, gin.H{"msg": err.Error()}
	}

	var list []goRedisAdmin.RedisList
	tx := db.Select("id", "desc").Where("menu_id = ?", redisId).Find(&list)

	if tx.Error != nil {
		return http.StatusInternalServerError, gin.H{"msg": tx.Error.Error()}
	}

	return http.StatusOK, gin.H{"data": list}
}
