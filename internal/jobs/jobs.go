package jobs

import (
	"redisadmin/internal/consts"
	"redisadmin/internal/databases"
	"redisadmin/internal/databases/go_redis_admin"
	"redisadmin/internal/redisPool"
	"sync"
)

// 数据库中redis的地址
// 没有监听数据库变动
var redis_db_conf map[int]*redisPool.ConnectConf

func init() {
	redis_db_conf = make(map[int]*redisPool.ConnectConf)
}

var lock sync.Mutex

// 从数据库获取redis链接配置
// @param id RedisList表的id
func Get_redis_connect_conf_from_db(id int) (*redisPool.ConnectConf, error) {
	lock.Lock()
	defer lock.Unlock()

	var conf *redisPool.ConnectConf
	var has bool

	conf, has = redis_db_conf[id]
	if !has {
		db, err := databases.Get_db(consts.DB_RD_AD_CONF, consts.DB_RD_AD_CONF_TAG_AD)
		if err != nil {
			return conf, err
		}

		var list go_redis_admin.RedisList
		if res := db.First(&list, id); res.Error != nil {
			return conf, res.Error
		}

		conf = &redisPool.ConnectConf{
			Desc:   list.Desc,
			Host:   list.Host,
			Port:   list.Port,
			Auth:   list.Auth,
			RdType: int(list.MenuId),
		}

		redis_db_conf[id] = conf
	}

	return conf, nil
}
