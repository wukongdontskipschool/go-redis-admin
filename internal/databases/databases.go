package databases

import (
	"errors"
	"fmt"
	"redisadmin/internal/configs"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ConnectConf struct {
	Db      string `yaml:"db"`
	Host    string `yaml:"host"`
	User    string `yaml:"user"`
	Pass    string `yaml:"pass"`
	Port    int    `yaml:"port"`
	Charset string `yaml:"charset"`
}

var lock sync.Mutex

// map[路径]map[库标志]connectConf
var all_configs map[string]map[string]ConnectConf

func init() {
	all_configs = map[string]map[string]ConnectConf{}
}

func Get_db(path string, name string) (db *gorm.DB, err error) {
	lock.Lock()
	defer lock.Unlock()

	var confMap map[string]ConnectConf
	var has bool
	confMap, has = all_configs[path]
	if !has {
		err = configs.Get_config(path, &confMap)

		if err != nil {
			return nil, errors.New("config path error:" + path)
		}

		all_configs[path] = confMap
	}

	conf, has := confMap[name]
	if !has {
		return nil, errors.New("config error:" + name)
	}

	dsn := conf.User + ":" + conf.Pass + "@tcp(" + conf.Host + ":" + fmt.Sprintf("%d", conf.Port) + ")/" + conf.Db + "?charset=" + conf.Charset + "&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return
}
