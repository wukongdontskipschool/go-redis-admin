package configs

import (
	"fmt"
	"io/ioutil"
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
