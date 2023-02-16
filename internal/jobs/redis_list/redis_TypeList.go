package redis_list

import (
	"net/http"
	"redisadmin/internal/consts"
	"redisadmin/internal/databases"
	"redisadmin/internal/databases/go_redis_admin"

	"github.com/gin-gonic/gin"
)

func TypeList() (int, gin.H) {
	db, err := databases.Get_db(consts.DB_RD_AD_CONF, consts.DB_RD_AD_CONF_TAG_AD)
	if err != nil {
		return http.StatusInternalServerError, gin.H{"msg": err}
	}

	var type_list []go_redis_admin.Menu
	if res := db.Where("pid = ?", consts.MENU_RD_LIST_ID).Find(&type_list); res.Error != nil {
		return http.StatusInternalServerError, gin.H{"msg": res.Error}
	}
	// log.Printf("%#v", res)
	// log.Printf("%#v", list)
	// ok, err := json.Marshal(list)
	// log.Printf(string(ok))
	return http.StatusOK, gin.H{"msg": "ok", "data": type_list}
}
