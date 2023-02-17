package redisList

import (
	"net/http"
	"redisadmin/internal/accessControl"
	"redisadmin/internal/consts"
	"redisadmin/internal/databases"
	"redisadmin/internal/databases/goRedisAdmin"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TypeDelete(id string) (int, gin.H) {
	db, err := databases.GetDb(consts.DB_RD_AD_CONF, consts.DB_RD_AD_CONF_TAG_AD)
	if err != nil {
		return http.StatusInternalServerError, gin.H{"msg": err.Error()}
	}

	menu := &goRedisAdmin.Menu{}
	tx := db.First(menu, id)
	if tx.Error != nil {
		return http.StatusInternalServerError, gin.H{"msg": tx.Error.Error()}
	}

	if menu.Pid != consts.MENU_RD_LIST_ID {
		return http.StatusForbidden, gin.H{"msg": "不支持该分类"}
	}

	idInt, _ := strconv.Atoi(id)
	tx = db.Where(goRedisAdmin.RedisList{MenuId: uint(idInt)}).First(goRedisAdmin.RedisList{})
	if tx.RowsAffected > 0 {
		return http.StatusForbidden, gin.H{"msg": "该分类包含reids子项"}
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		// 删规则
		if res := db.Unscoped().Where(goRedisAdmin.Rule{Rule: menu.Rule}).Delete(&goRedisAdmin.Rule{}); res.Error != nil {
			return res.Error
		}

		// 删菜单
		if res := db.Unscoped().Delete(&menu); res.Error != nil {
			return res.Error
		}

		// 删权限
		if _, err := accessControl.DeletePermissionsForRule(menu.Rule); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return http.StatusInternalServerError, gin.H{"msg": err.Error()}
	}

	return http.StatusOK, gin.H{"data": menu.ID}
}
