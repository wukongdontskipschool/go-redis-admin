package admin

import (
	"log"
	"net/http"
	"redisadmin/internal/accessControl"
	"redisadmin/internal/consts"
	"redisadmin/internal/databases"
	"redisadmin/internal/databases/go_redis_admin"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RuleDelete(rId string) (int, gin.H) {
	db, _ := databases.Get_db(consts.DB_RD_AD_CONF, consts.DB_RD_AD_CONF_TAG_AD)

	role := go_redis_admin.Rule{}
	tx := db.First(&role, rId)
	if tx.Error != nil {
		return http.StatusInternalServerError, gin.H{"msg": tx.Error.Error()}
	}

	log.Println(rId)
	log.Println(role)
	err := db.Transaction(func(tx *gorm.DB) error {
		// 删规则
		if res := db.Unscoped().Delete(&role); res.Error != nil {
			return res.Error
		}

		// 删权限
		if _, err := accessControl.DeletePermissionsForActRule(role.Rule, role.Act); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return http.StatusInternalServerError, gin.H{"msg": err.Error()}
	}

	return http.StatusOK, gin.H{"status": 0, "msg": "ok"}
}
