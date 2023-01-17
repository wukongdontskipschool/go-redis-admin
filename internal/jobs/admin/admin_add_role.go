package admin

import (
	"net/http"
	"redisadmin/internal/accessControl"
	"redisadmin/internal/consts"
	"redisadmin/internal/databases"
	"redisadmin/internal/databases/go_redis_admin"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Add_role(name string, ruleIds map[string]string) (int, gin.H) {
	user := go_redis_admin.Role{Name: name}

	db, _ := databases.Get_db(consts.DB_RD_AD_CONF, consts.DB_RD_AD_CONF_TAG_AD)
	tx := db.Create(&user)
	if tx.Error != nil {
		return http.StatusInternalServerError, gin.H{"msg": tx.Error.Error()}
	}

	var rIds []uint
	for _, val := range ruleIds {
		if in, err := strconv.Atoi(val); err == nil {
			rIds = append(rIds, uint(in))
		}
	}

	if len(rIds) == 0 {
		return http.StatusOK, gin.H{"status": 0, "msg": "ok", "data": user.ID}
	}

	var rules []go_redis_admin.Rule
	db.Where(rIds).Find(&rules)

	for _, rule := range rules {
		id := strconv.Itoa(int(user.ID))
		accessControl.AddPolicy(consts.ROLE_PRE+id, rule.Rule, rule.Act)
	}

	return http.StatusOK, gin.H{"status": 0, "msg": "ok", "data": user.ID}
}
