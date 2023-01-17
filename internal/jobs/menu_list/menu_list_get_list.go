package menu_list

import (
	"fmt"
	"net/http"
	"redisadmin/internal/accessControl"
	"redisadmin/internal/consts"
	"redisadmin/internal/databases"
	"redisadmin/internal/databases/go_redis_admin"

	"github.com/gin-gonic/gin"
)

type menu_item_map map[string]interface{}

func Get_list(uId uint) (int, gin.H) {
	db, err := databases.Get_db(consts.DB_RD_AD_CONF, consts.DB_RD_AD_CONF_TAG_AD)
	if err != nil {
		return http.StatusInternalServerError, gin.H{"status": 1, "msg": err.Error()}
	}

	var user go_redis_admin.User
	res := db.First(&user, "ID = ?", uId)
	if res.Error != nil {
		return http.StatusInternalServerError, gin.H{"status": 1, "msg": err.Error()}
	}

	if res.RowsAffected < 1 {
		return http.StatusInternalServerError, gin.H{"status": 1, "msg": "没有找到用户"}
	}

	isSup := uId == consts.SUP_ADMIN_ID
	var list []go_redis_admin.Menu

	okps := make(map[string]struct{})
	if !isSup {
		// 获取权限列表
		uTag := fmt.Sprintf("%s%d", consts.ROLE_PRE, user.RoleId)
		ps := accessControl.Enforcer.GetPermissionsForUser(uTag)

		for _, p := range ps {
			if p[2] == "GET" {
				okps[p[1]] = struct{}{}
			}
		}

		if len(okps) == 0 {
			return http.StatusOK, gin.H{"status": 0, "msg": "ok", "data": []menu_item_map{}}
		}
	}

	if res := db.Order("pid desc").Find(&list); res.Error != nil {
		return http.StatusOK, gin.H{"status": 0, "msg": "ok", "data": []menu_item_map{}}
	}

	pid0 := uint(0)
	data := []menu_item_map{}
	data_map := map[uint]menu_item_map{}

	okPid := make(map[uint]struct{})
	for _, v := range list {
		_, has := okps[v.Rule]
		_, has1 := okPid[v.ID]
		if !has && !has1 && !isSup {
			continue
		}

		if v.Pid > 0 {
			okPid[v.Pid] = struct{}{}
		}

		data_map[v.ID] = menu_item_map{
			"id":       v.ID,
			"name":     v.Name,
			"icon":     v.Icon,
			"url":      v.Url,
			"pid":      v.Pid,
			"children": []menu_item_map{},
		}

		// 第一级菜单
		if v.Pid == pid0 {
			data = append(data, data_map[v.ID])
		}
	}

	// 整理子菜单
	for _, v := range data_map {
		pid1, _ := v["pid"].(uint)
		id1, _ := v["id"].(uint)
		if pid1 != pid0 {
			c, has := data_map[pid1]["children"].([]menu_item_map)
			if !has {
				continue
			}
			b, has := data_map[id1]
			if !has {
				continue
			}
			data_map[pid1]["children"] = append(c, b)
		}
	}

	return 200, gin.H{"status": 0, "msg": "ok", "data": data}
}
