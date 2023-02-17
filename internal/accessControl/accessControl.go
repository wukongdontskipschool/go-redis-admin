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
	err := configs.GetConfig(consts.DB_RD_AD_CONF, &confMap)
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

// 加规则
func AddPolicy(sub string, obj string, act string) (ok bool, err error) {
	ok, err = Enforcer.AddPolicy(sub, obj, act)
	return
}

// 批量加
func AddNamedPolicies(rules [][]string) (ok bool, err error) {
	// rules = [][]string{
	// 	[]string{"jack", "data4", "read"},
	// 	[]string{"katy", "data4", "write"},
	// 	[]string{"leyo", "data4", "read"},
	// 	[]string{"ham", "data4", "write"},
	// }

	ok, err = Enforcer.AddNamedPolicies("p", rules)
	return
}

// 获取角色权限列表
func GetPolicyByRole(rId string) [][]string {
	policys := Enforcer.GetFilteredNamedPolicy("p", 0, consts.ROLE_PRE+rId)
	return policys
}

//拦截器
func Authorize(sub string, obj string, act string) (ok bool, err error) {
	ok, err = Enforcer.Enforce(sub, obj, act)
	return
}

// 删除角色权限
func DeletePermissionsForUser(rId string) (bool, error) {
	ok, err := Enforcer.DeletePermissionsForUser(consts.ROLE_PRE + rId)
	return ok, err
}

// 删除带方法权限
func DeletePermissionsForActRule(rule string, act string) (ok bool, err error) {
	return Enforcer.RemoveFilteredPolicy(1, rule, act)
}

// 删除规则权限
func DeletePermissionsForRule(rule string) (ok bool, err error) {
	return Enforcer.RemoveFilteredPolicy(1, rule)
}

func Test(rId string) bool {
	// a := Enforcer.GetPermissionsForUser(consts.ROLE_PRE + rId)
	b, _ := Enforcer.RemoveFilteredPolicy(1, rId)
	log.Println(b)
	return false
}
