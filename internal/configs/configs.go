package configs

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"redisadmin/internal/consts"
	"sync"

	"gopkg.in/yaml.v2"
)

var allConfigs map[string][]byte
var m sync.RWMutex

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// 随机字符串
func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func init() {
	allConfigs = map[string][]byte{}
}

func GetConfig(name string, out interface{}) error {
	m.Lock()
	defer m.Unlock()

	var err error

	bytes, has := allConfigs[name]
	if !has {
		bytes, err = ioutil.ReadFile(fmt.Sprintf("internal/configs/%s.yaml", name))

		if err != nil {
			return err
		}

		allConfigs[name] = bytes
	}

	err = yaml.Unmarshal(bytes, out)

	return err
}

func GetEnvVal(key string) (val string) {
	var confMap map[string]string

	err := GetConfig(consts.ENV_CONF, &confMap)
	if err != nil {
		return
	}

	var has bool
	val, has = confMap[key]
	if !has {
		val = ""
	}

	return
}

func CreateEnv() error {
	err := os.MkdirAll("internal/configs/env/", 0700)
	if err != nil {
		log.Println("dir create fail")
		return err
	}
	file, error := os.Create("internal/configs/env/env.yaml")

	if error != nil {
		log.Println("file create fail")
		return error
	}

	defer file.Close()
	file.WriteString("pSalt: " + RandStringBytes(6) + "\n")      // 密码盐
	file.WriteString("jetKey: " + RandStringBytes(6) + "\n")     // jwt密钥
	file.WriteString("cryptoKey: " + RandStringBytes(16) + "\n") // redis密钥
	file.WriteString("ginMode: debug\n")                         // ginmode

	return nil
}

func CreateDatabase() error {
	err := os.MkdirAll("internal/configs/databases/", 0700)
	if err != nil {
		log.Println("dir create fail")
		return err
	}
	file, error := os.Create("internal/configs/databases/redisAdmin.yaml")

	if error != nil {
		log.Println("file create fail")
		return error
	}

	defer file.Close()
	str := `admin:
  db: go_redis_manager
  host: 127.0.0.1
  user: root
  pass: root
  port: 3306
  charset: utf8mb4`

	file.WriteString(str) //写入字符串

	return nil
}

func CreateCasbin() error {
	err := os.MkdirAll("internal/configs/casbin/", 0700)
	if err != nil {
		log.Println("dir create fail")
		return err
	}
	file, error := os.Create("internal/configs/casbin/rbacModels.conf")

	if error != nil {
		log.Println("file create fail")
		return error
	}

	defer file.Close()
	str := `[request_definition]
r = sub, obj, act
	
[policy_definition]
p = sub, obj, act
	
[role_definition]
g = _, _
	
[policy_effect]
e = some(where (p.eft == allow))
	
[matchers]
m = r.sub == "role_1" || g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act`

	file.WriteString(str) //写入字符串

	return nil
}
