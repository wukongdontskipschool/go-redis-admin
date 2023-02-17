package admin

import (
	"net/http"
	"redisadmin/internal/accessControl"
	"redisadmin/internal/consts"
	"redisadmin/internal/databases"
	"redisadmin/internal/databases/goRedisAdmin"

	"github.com/gin-gonic/gin"
)

func RoleRule(rId string) (int, gin.H) {
	db, _ := databases.GetDb(consts.DB_RD_AD_CONF, consts.DB_RD_AD_CONF_TAG_AD)

	var role goRedisAdmin.Role
	tx := db.First(&role, rId)
	if tx.Error != nil {
		return http.StatusInternalServerError, gin.H{"msg": tx.Error.Error()}
	}

	state, rules := GetRuleList()
	if state != http.StatusOK {
		return state, rules
	}

	filteredPolicy := accessControl.GetPolicyByRole(rId)

	okRules := make([]string, len(filteredPolicy))
	for i, v := range filteredPolicy {
		okRules[i] = v[2] + "_" + v[1]
	}

	return http.StatusOK, gin.H{"status": 0, "msg": "ok", "data": rules["data"], "okRules": okRules, "name": role.Name}
}
