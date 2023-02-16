package redis_list

import (
	"net/http"
	"redisadmin/internal/consts"
	"redisadmin/internal/databases"
	"redisadmin/internal/databases/go_redis_admin"
	"redisadmin/internal/services/cryptoAes"

	"github.com/gin-gonic/gin"
)

func ItemUpdate(id string, desc string, host string, port string, auth string, hasAuth string, menuId string) (int, gin.H) {
	db, err := databases.Get_db(consts.DB_RD_AD_CONF, consts.DB_RD_AD_CONF_TAG_AD)
	if err != nil {
		return http.StatusInternalServerError, gin.H{"msg": err.Error()}
	}

	var menu go_redis_admin.Menu
	if err := db.First(&menu, menuId).Error; err != nil || menu.Pid != consts.MENU_RD_LIST_ID {
		return http.StatusInternalServerError, gin.H{"msg": "非法分类"}
	}

	redisList := go_redis_admin.RedisList{}

	tx := db.First(&redisList, id)
	if tx.Error != nil {
		return http.StatusInternalServerError, gin.H{"msg": tx.Error.Error()}
	}

	filed1 := make(map[string]interface{})

	filed1["desc"] = desc
	filed1["host"] = host
	filed1["port"] = port
	filed1["menu_id"] = menuId

	if hasAuth == "on" {
		if auth != "" {
			auth, err = cryptoAes.Encrypt(auth, "")
			if err != nil {
				return http.StatusInternalServerError, gin.H{"msg": "密码加密错误:" + err.Error()}
			}
		}
		filed1["auth"] = auth
	}

	tx = db.Model(&redisList).Updates(filed1)
	if tx.Error != nil {
		return http.StatusInternalServerError, gin.H{"msg": tx.Error.Error()}
	}

	return http.StatusOK, gin.H{"data": redisList.ID}
}
