package redisList

import (
	"net/http"
	"redisadmin/internal/consts"
	"redisadmin/internal/databases"
	"redisadmin/internal/databases/goRedisAdmin"

	"github.com/gin-gonic/gin"
)

func ItemDelete(id string) (int, gin.H) {
	db, err := databases.GetDb(consts.DB_RD_AD_CONF, consts.DB_RD_AD_CONF_TAG_AD)
	if err != nil {
		return http.StatusInternalServerError, gin.H{"msg": err.Error()}
	}

	var menu goRedisAdmin.RedisList
	if err := db.First(&menu, id).Error; err != nil {
		return http.StatusInternalServerError, gin.H{"msg": "非法redis"}
	}

	err = db.Unscoped().Delete(menu).Error
	if err != nil {
		return http.StatusInternalServerError, gin.H{"msg": err.Error()}
	}

	return http.StatusOK, gin.H{"data": menu.ID}
}
