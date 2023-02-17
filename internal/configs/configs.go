package configs

import (
	"fmt"
	"io/ioutil"
	"redisadmin/internal/consts"
	"sync"

	"gopkg.in/yaml.v2"
)

var allConfigs map[string][]byte
var m sync.RWMutex

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
