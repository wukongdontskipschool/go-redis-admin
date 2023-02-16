package redis_list

import (
	"net/http"
	"redisadmin/internal/consts"
	"redisadmin/internal/databases"
	"redisadmin/internal/databases/go_redis_admin"

	"github.com/gin-gonic/gin"
)

func ItemList() (int, gin.H) {
	db, err := databases.Get_db(consts.DB_RD_AD_CONF, consts.DB_RD_AD_CONF_TAG_AD)
	if err != nil {
		return http.StatusInternalServerError, gin.H{"msg": err}
	}

	type res struct {
		go_redis_admin.RedisList
		Tname string
	}
	var user []res
	db.Model(&go_redis_admin.RedisList{}).Select("redis_lists.*, '' as auth, menus.name as tname").Joins("left join menus on menus.id = redis_lists.menu_id").Scan(&user)

	return http.StatusOK, gin.H{"msg": "ok", "data": user}
}
