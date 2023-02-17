package redisPool

import (
	"fmt"
	"log"
	"redisadmin/internal/services/cryptoAes"
	"sync"

	"github.com/go-redis/redis/v8"
)

type ConnectConf struct {
	Desc   string `name:"unique"`
	Host   string
	Port   string
	Auth   string
	RdType int
}

var lock sync.Mutex
var allPool map[string]*redis.Client

func init() {
	allPool = map[string]*redis.Client{}
}

func GetRedis(conf *ConnectConf, db int) *redis.Client {
	lock.Lock()
	defer lock.Unlock()

	var rdPool *redis.Client
	// var has bool

	addr := fmt.Sprintf("%s:%s", conf.Host, conf.Port)
	key := fmt.Sprintf("%s:%s", addr, conf.Auth)

	auth := conf.Auth
	if auth != "" {
		var err error
		auth, err = cryptoAes.Decrypt(conf.Auth, "")
		if err != nil {
			log.Println("[warn]", conf, "redis密码解密失败", err.Error())
			auth = conf.Auth
		}
	}

	// rdPool, has = allPool[key]
	// if !has {
	rdPool = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: auth,
		DB:       db,
		PoolSize: 1,
	})

	allPool[key] = rdPool
	// }

	return rdPool
}
