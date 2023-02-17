package redisList

import (
	"net/http"
	"redisadmin/internal/consts"
	"redisadmin/internal/databases"
	"redisadmin/internal/databases/goRedisAdmin"

	"github.com/gin-gonic/gin"
)

func TypeList() (int, gin.H) {
	db, err := databases.GetDb(consts.DB_RD_AD_CONF, consts.DB_RD_AD_CONF_TAG_AD)
	if err != nil {
		return http.StatusInternalServerError, gin.H{"msg": err}
	}

	var typeList []goRedisAdmin.Menu
	if res := db.Where("pid = ?", consts.MENU_RD_LIST_ID).Find(&typeList); res.Error != nil {
		return http.StatusInternalServerError, gin.H{"msg": res.Error}
	}
	// log.Printf("%#v", res)
	// log.Printf("%#v", list)
	// ok, err := json.Marshal(list)
	// log.Printf(string(ok))
	return http.StatusOK, gin.H{"msg": "ok", "data": typeList}
}
