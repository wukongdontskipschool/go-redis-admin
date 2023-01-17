package redis_list

import (
	"fmt"
	"log"
	"net/http"
	"redisadmin/internal/consts"
	"redisadmin/internal/databases"
	"redisadmin/internal/databases/go_redis_admin"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index() (int, gin.H) {
	db, err := databases.Get_db("databases/redis_admin", "admin")
	if err != nil {
		return http.StatusInternalServerError, gin.H{"msg": err}
	}

	var list []go_redis_admin.RedisList
	if res := db.Find(&list); res.Error != nil {
		return http.StatusInternalServerError, gin.H{"msg": res.Error}
	}

	var type_list []go_redis_admin.RedisListTypes
	if res := db.Find(&type_list); res.Error != nil {
		return http.StatusInternalServerError, gin.H{"msg": res.Error}
	}
	// log.Printf("%#v", res)
	// log.Printf("%#v", list)
	// ok, err := json.Marshal(list)
	// log.Printf(string(ok))
	return http.StatusOK, gin.H{"l": list, "tl": type_list}
}

func Migrate() {
	db, _ := databases.Get_db("databases/redis_admin", "admin")
	db.AutoMigrate(&go_redis_admin.User{})
	db.AutoMigrate(&go_redis_admin.Role{})
	db.AutoMigrate(&go_redis_admin.Rule{})
	db.AutoMigrate(&go_redis_admin.RedisList{})
	db.AutoMigrate(&go_redis_admin.RedisListTypes{})
	db.AutoMigrate(&go_redis_admin.Menu{})

	InitMenu(db)
	InitRole(db)
	InitRule(db)
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

	log.Println("tbna", stmt.Schema.Table)

	if tx.Error != nil {
		panic(tx.Error)
	}

	tx = db.Create(role)
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
