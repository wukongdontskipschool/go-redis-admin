package menuList

import (
	"fmt"
	"net/http"
	"redisadmin/internal/accessControl"
	"redisadmin/internal/consts"
	"redisadmin/internal/databases"
	"redisadmin/internal/databases/goRedisAdmin"

	"github.com/gin-gonic/gin"
)

type menuItemMap map[string]interface{}

func GetList(uId uint) (int, gin.H) {
	db, err := databases.GetDb(consts.DB_RD_AD_CONF, consts.DB_RD_AD_CONF_TAG_AD)
	if err != nil {
		return http.StatusInternalServerError, gin.H{"status": 1, "msg": err.Error()}
	}

	var user goRedisAdmin.User
	res := db.First(&user, "ID = ?", uId)
	if res.Error != nil {
		return http.StatusInternalServerError, gin.H{"status": 1, "msg": res.Error.Error()}
	}

	if res.RowsAffected < 1 {
		return http.StatusInternalServerError, gin.H{"status": 1, "msg": "没有找到用户"}
	}

	isSup := user.RoleId == consts.SUP_ADMIN_RID
	var list []goRedisAdmin.Menu

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
			return http.StatusOK, gin.H{"status": 0, "msg": "ok", "data": []menuItemMap{}}
		}
	}

	if res := db.Order("pid desc").Find(&list); res.Error != nil {
		return http.StatusOK, gin.H{"status": 0, "msg": "ok", "data": []menuItemMap{}}
	}

	pid0 := uint(0)
	data := []menuItemMap{}
	dataMap := map[uint]menuItemMap{}
	listDataMap := []menuItemMap{} // 保证排序

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

		dataMap[v.ID] = menuItemMap{
			"id":       v.ID,
			"name":     v.Name,
			"icon":     v.Icon,
			"url":      v.Url,
			"pid":      v.Pid,
			"children": []menuItemMap{},
		}

		listDataMap = append(listDataMap, dataMap[v.ID])

		// 第一级菜单
		if v.Pid == pid0 {
			data = append(data, dataMap[v.ID])
		}
	}

	// 整理子菜单
	for _, v := range listDataMap {
		pid1, _ := v["pid"].(uint)
		id1, _ := v["id"].(uint)
		if pid1 != pid0 {
			c, has := dataMap[pid1]["children"].([]menuItemMap)
			if !has {
				continue
			}
			b, has := dataMap[id1]
			if !has {
				continue
			}
			dataMap[pid1]["children"] = append(c, b)
		}
	}

	return 200, gin.H{"status": 0, "msg": "ok", "data": data}
}
