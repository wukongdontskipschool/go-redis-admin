package redis_list

import (
	"fmt"
	"net/http"
	"redisadmin/internal/consts"
	"redisadmin/internal/databases"
	"redisadmin/internal/databases/go_redis_admin"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func TypeStore(name string) (int, gin.H) {
	db, err := databases.Get_db(consts.DB_RD_AD_CONF, consts.DB_RD_AD_CONF_TAG_AD)
	if err != nil {
		return http.StatusInternalServerError, gin.H{"msg": err.Error()}
	}

	menu := go_redis_admin.Menu{Name: name, Pid: consts.MENU_RD_LIST_ID, State: 1}
	err = db.Transaction(func(tx *gorm.DB) error {
		if res := db.Create(&menu); res.Error != nil {
			return res.Error
		}

		menu.Url = fmt.Sprintf("%s%d", "./pages/redisItem/category.html?typeId=", menu.ID)
		menu.Rule = fmt.Sprintf("%s%d", "/v1/redisItem/redisList/", menu.ID)
		tx1 := db.Save(&menu)
		if tx1.Error != nil {
			return tx.Error
		}

		rules := []go_redis_admin.Rule{
			{
				Rule: menu.Rule,
				Act:  "GET",
				Desc: menu.Name + "列表",
			},
			{
				Rule: menu.Rule,
				Act:  "POST",
				Desc: menu.Name + "新增",
			},
			{
				Rule: menu.Rule,
				Act:  "PUT",
				Desc: menu.Name + "修改",
			},
			{
				Rule: menu.Rule,
				Act:  "DELETE",
				Desc: menu.Name + "删除",
			},
		}
		if res := db.Create(&rules); res.Error != nil {
			return res.Error
		}

		return nil
	})

	if err != nil {
		return http.StatusInternalServerError, gin.H{"msg": err.Error()}
	}

	return http.StatusOK, gin.H{"data": menu.ID}
}
