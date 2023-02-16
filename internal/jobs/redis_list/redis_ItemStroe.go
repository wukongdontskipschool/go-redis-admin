package redis_list

import (
	"net/http"
	"redisadmin/internal/consts"
	"redisadmin/internal/databases"
	"redisadmin/internal/databases/go_redis_admin"
	"redisadmin/internal/services/cryptoAes"

	"github.com/gin-gonic/gin"
)

func ItemStore(desc string, host string, port string, auth string, menuId string) (int, gin.H) {
	db, err := databases.Get_db(consts.DB_RD_AD_CONF, consts.DB_RD_AD_CONF_TAG_AD)
	if err != nil {
		return http.StatusInternalServerError, gin.H{"msg": err.Error()}
	}

	var menu go_redis_admin.Menu
	if err := db.First(&menu, menuId).Error; err != nil || menu.Pid != consts.MENU_RD_LIST_ID {
		return http.StatusInternalServerError, gin.H{"msg": "非法分类"}
	}

	if auth != "" {
		auth, err = cryptoAes.Encrypt(auth, "")
		if err != nil {
			return http.StatusInternalServerError, gin.H{"msg": "密码加密错误:" + err.Error()}
		}
	}

	redisList := go_redis_admin.RedisList{
		MenuId: menu.ID,
		Desc:   desc,
		Host:   host,
		Port:   port,
		Auth:   auth,
	}

	tx := db.Create(&redisList)
	if tx.Error != nil {
		return http.StatusInternalServerError, gin.H{"msg": tx.Error.Error()}
	}

	return http.StatusOK, gin.H{"data": redisList.ID}
}
