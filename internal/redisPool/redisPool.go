package redisPool

import (
	"fmt"
	"sync"

	"github.com/go-redis/redis/v8"
)

type ConnectConf struct {
	Desc string `name:"unique"`
	Host string
	Port string
	Auth string
}

var lock sync.Mutex
var all_pool map[string]*redis.Client

func init() {
	all_pool = map[string]*redis.Client{}
}

func Get_redis(conf *ConnectConf) *redis.Client {
	lock.Lock()
	defer lock.Unlock()

	var rd_pool *redis.Client
	var has bool

	addr := fmt.Sprintf("%s:%s", conf.Host, conf.Port)
	key := fmt.Sprintf("%s:%s", addr, conf.Auth)

	rd_pool, has = all_pool[key]
	if !has {
		rd_pool = redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: conf.Auth,
			DB:       0,
			PoolSize: 5,
		})

		all_pool[key] = rd_pool
	}

	return rd_pool
}
