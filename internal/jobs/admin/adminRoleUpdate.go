package admin

import (
	"fmt"
	"net/http"
	"redisadmin/internal/accessControl"
	"redisadmin/internal/consts"
	"redisadmin/internal/databases"
	"redisadmin/internal/databases/goRedisAdmin"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RoleUpdate(rId string, name string, ruleIds map[string]string) (int, gin.H) {
	db, _ := databases.GetDb(consts.DB_RD_AD_CONF, consts.DB_RD_AD_CONF_TAG_AD)

	if rId == fmt.Sprintf("%d", consts.SUP_ADMIN_RID) {
		return http.StatusForbidden, gin.H{"msg": "超级管理员包含所有权限"}
	}

	user := goRedisAdmin.Role{Name: name}
	tx := db.First(&user, rId)
	if tx.Error != nil {
		return http.StatusInternalServerError, gin.H{"msg": tx.Error.Error()}
	}

	// 更新角色名
	if name != user.Name {
		tx = db.Model(&user).Update("name", name)
		if tx.Error != nil {
			return http.StatusInternalServerError, gin.H{"msg": tx.Error.Error()}
		}
	}

	var rIds []uint
	for _, val := range ruleIds {
		if in, err := strconv.Atoi(val); err == nil {
			rIds = append(rIds, uint(in))
		}
	}

	// 清空权限 图便捷
	if _, err := accessControl.DeletePermissionsForUser(rId); err != nil {
		return http.StatusInternalServerError, gin.H{"msg": err.Error()}
	}

	if len(rIds) > 0 {
		var rules []goRedisAdmin.Rule
		db.Where(rIds).Find(&rules)

		rules1 := make([][]string, len(rules))
		for i, rule := range rules {
			// accessControl.AddPolicy(consts.ROLE_PRE+rId, rule.Rule, rule.Act)
			rules1[i] = []string{consts.ROLE_PRE + rId, rule.Rule, rule.Act}
		}

		// 加权限
		if _, err := accessControl.AddNamedPolicies(rules1); err != nil {
			return http.StatusInternalServerError, gin.H{"msg": err.Error()}
		}
	}

	return http.StatusOK, gin.H{"status": 0, "msg": "ok", "data": user.ID}
}
