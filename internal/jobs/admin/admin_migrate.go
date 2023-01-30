package admin

import (
	"fmt"
	"redisadmin/internal/consts"
	"redisadmin/internal/databases"
	"redisadmin/internal/databases/go_redis_admin"

	"gorm.io/gorm"
)

func Migrate() {
	db, _ := databases.Get_db(consts.DB_RD_AD_CONF, consts.DB_RD_AD_CONF_TAG_AD)
	db.AutoMigrate(&go_redis_admin.User{})
	db.AutoMigrate(&go_redis_admin.Role{})
	db.AutoMigrate(&go_redis_admin.Rule{})
	db.AutoMigrate(&go_redis_admin.RedisList{})
	db.AutoMigrate(&go_redis_admin.RedisListTypes{})
	db.AutoMigrate(&go_redis_admin.Menu{})

	InitMenu(db)
	InitRole(db)
	InitRule(db)
	InitRedisList(db)
}

func InitMenu(db *gorm.DB) {
	tx := db.Exec("truncate table menus")
	if tx.Error != nil {
		panic(tx.Error)
	}

	admin := &go_redis_admin.Menu{
		Pid:   0,
		Name:  "管理员管理",
		Url:   "",
		State: 1,
	}

	tx = db.Create(admin)
	if tx.Error != nil {
		panic(tx.Error)
	}

	menu1 := []go_redis_admin.Menu{
		{
			Pid:   admin.ID,
			Name:  "管理员列表",
			Url:   "./pages/admin/list.html",
			Rule:  "/v1/admin/user",
			State: 1,
		},
		{
			Pid:   admin.ID,
			Name:  "角色列表",
			Url:   "./pages/admin/role.html",
			Rule:  "/v1/admin/role",
			State: 1,
		},
		{
			Pid:   admin.ID,
			Name:  "权限分类",
			Url:   "./pages/admin/cate.html",
			State: 1,
		},
		{
			Pid:   admin.ID,
			Name:  "权限列表",
			Rule:  "/v1/admin/rule",
			Url:   "./pages/admin/rule.html",
			State: 1,
		},
	}

	db.Create(&menu1)

	tx = db.Create(&go_redis_admin.Menu{
		Pid:   0,
		Name:  "redis管理",
		Url:   "./pages/redisList/index.html",
		Rule:  "/v1/redisList",
		State: 1,
	})
	if tx.Error != nil {
		panic(tx.Error)
	}

	tx = db.Create(&go_redis_admin.Menu{
		Model: gorm.Model{ID: consts.MENU_RD_LIST_ID},
		Pid:   0,
		Name:  "redis列表",
		Url:   "",
		State: 1,
	})
	if tx.Error != nil {
		panic(tx.Error)
	}

	tx = db.Create(&go_redis_admin.Menu{
		Pid:   consts.MENU_RD_LIST_ID,
		Name:  "默认",
		Url:   fmt.Sprintf("%s%d", "./pages/redisItem/category.html?typeId=", consts.MENU_RD_LIST_ID+1),
		Rule:  fmt.Sprintf("%s%d", "/v1/redisItem/redisList/", consts.MENU_RD_LIST_ID+1),
		State: 1,
	})
	if tx.Error != nil {
		panic(tx.Error)
	}

}

func InitRole(db *gorm.DB) {
	role := &go_redis_admin.Role{
		Name: "admin",
	}

	stmt := &gorm.Statement{DB: db}
	stmt.Parse(role)
	tx := db.Exec("truncate table " + stmt.Schema.Table)

	if tx.Error != nil {
		panic(tx.Error)
	}

	tx = db.Create(role)
	if tx.Error != nil {
		panic(tx.Error)
	}
}

func InitRedisList(db *gorm.DB) {
	redisList := &go_redis_admin.RedisList{
		MenuId: consts.MENU_RD_LIST_ID + 1,
		Desc:   "测试",
		Host:   "127.0.0.1",
		Port:   "6379",
		Auth:   "",
	}

	stmt := &gorm.Statement{DB: db}
	stmt.Parse(redisList)
	tx := db.Exec("truncate table " + stmt.Schema.Table)

	if tx.Error != nil {
		panic(tx.Error)
	}

	tx = db.Create(redisList)
	if tx.Error != nil {
		panic(tx.Error)
	}
}

func InitRule(db *gorm.DB) {
	menu1 := []go_redis_admin.Rule{
		{
			Rule: "/v1/admin/user",
			Act:  "GET",
			Desc: "管理员列表",
		},
		{
			Rule: "/v1/admin/role",
			Act:  "GET",
			Desc: "角色列表",
		},
		{
			Rule: "/v1/admin/rule",
			Act:  "GET",
			Desc: "权限列表",
		},
		{
			Rule: "/v1/redisList",
			Act:  "GET",
			Desc: "redis管理",
		},
		{
			Rule: fmt.Sprintf("%s%d", "/v1/redisItem/redisList/", consts.MENU_RD_LIST_ID+1),
			Act:  "GET",
			Desc: "redis默认列表",
		},
	}

	stmt := &gorm.Statement{DB: db}
	stmt.Parse(go_redis_admin.Rule{})
	tx := db.Exec("truncate table " + stmt.Schema.Table)
	if tx.Error != nil {
		panic(tx.Error)
	}

	tx = db.Create(&menu1)
	if tx.Error != nil {
		panic(tx.Error)
	}
}