package redis_list

import (
	"database/sql"
	"log"
	"net/http"
	"redisadmin/internal/consts"
	"redisadmin/internal/databases"
	"redisadmin/internal/databases/go_redis_admin"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TypeUpdate(id string, name string) (int, gin.H) {
	db, err := databases.Get_db(consts.DB_RD_AD_CONF, consts.DB_RD_AD_CONF_TAG_AD)
	if err != nil {
		return http.StatusInternalServerError, gin.H{"msg": err.Error()}
	}

	menu := &go_redis_admin.Menu{}
	tx := db.First(menu, id)
	if tx.Error != nil {
		return http.StatusInternalServerError, gin.H{"msg": tx.Error.Error()}
	}

	if menu.Pid != consts.MENU_RD_LIST_ID {
		return http.StatusForbidden, gin.H{"msg": "不支持该分类"}
	}

	oldName := menu.Name
	if menu.Name == name {
		return http.StatusForbidden, gin.H{"msg": "没有修改"}
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		if res := db.Model(&menu).Update("name", name); res.Error != nil {
			return res.Error
		}

		// 改rule的名称
		sqlStr := "Update `rules` SET `desc` = replace(`desc`, @oldName, @name) Where `rule` = @rule"
		res := db.Exec(sqlStr, sql.Named("oldName", oldName), sql.Named("name", name), sql.Named("rule", menu.Rule))
		log.Println(res)
		if res.Error != nil {
			return res.Error
		}

		return nil
	})

	if err != nil {
		return http.StatusInternalServerError, gin.H{"msg": err.Error()}
	}

	return http.StatusOK, gin.H{"data": menu.ID}
}
