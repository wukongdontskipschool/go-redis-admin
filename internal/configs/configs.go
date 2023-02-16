package configs

import (
	"fmt"
	"io/ioutil"
	"redisadmin/internal/consts"
	"sync"

	"gopkg.in/yaml.v2"
)

var all_configs map[string][]byte
var m sync.RWMutex

func init() {
	all_configs = map[string][]byte{}
}

func Get_config(name string, out interface{}) error {
	m.Lock()
	defer m.Unlock()

	var err error

	bytes, has := all_configs[name]
	if !has {
		bytes, err = ioutil.ReadFile(fmt.Sprintf("internal/configs/%s.yaml", name))

		if err != nil {
			return err
		}

		all_configs[name] = bytes
	}

	err = yaml.Unmarshal(bytes, out)

	return err
}

func GetEnvVal(key string) (val string) {
	var confMap map[string]string

	err := Get_config(consts.ENV_CONF, &confMap)
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
