package accessControl

import (
	"log"
	"redisadmin/internal/configs"
	"redisadmin/internal/consts"
	"redisadmin/internal/databases"
	"strconv"

	"github.com/casbin/casbin/v2"
	xormadapter "github.com/casbin/xorm-adapter/v2"
	_ "github.com/go-sql-driver/mysql"
)

var Enforcer *casbin.SyncedEnforcer

func init() {
	var confMap map[string]databases.ConnectConf
	err := configs.Get_config(consts.DB_RD_AD_CONF, &confMap)
	if err != nil {
		panic("数据库报错")
	}

	conf, has := confMap[consts.DB_RD_AD_CONF_TAG_AD]
	if !has {
		panic("数据库报错")
	}

	dsn := conf.User + ":" + conf.Pass + "@tcp(" + conf.Host + ":" + strconv.Itoa(conf.Port) + ")/" + conf.Db + "?charset=" + conf.Charset
	adapter, err := xormadapter.NewAdapter("mysql", dsn, true)
	if err != nil {
		log.Printf("连接数据库错误: %v", err)
		return
	}

	Enforcer, err = casbin.NewSyncedEnforcer(consts.CASBIN_RBAC_CONF, adapter)
	if err != nil {
		log.Printf("初始化casbin错误: %v", err)
		return
	}

	//从DB加载策略
	Enforcer.LoadPolicy()
}

func AddPolicy(sub string, obj string, act string) (ok bool, err error) {
	ok, err = Enforcer.AddPolicy(sub, obj, act)
	return
}

//拦截器
func Authorize(sub string, obj string, act string) (ok bool, err error) {
	ok, err = Enforcer.Enforce(sub, obj, act)
	return
}
