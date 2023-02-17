package admin

import (
	"crypto/md5"
	"fmt"
	"redisadmin/internal/configs"
	"redisadmin/internal/consts"
)

// md5 加盐后密码
func getMd5SaltPass(pass string) string {
	salt := getPassSalt()
	return fmt.Sprintf("%x", md5.Sum([]byte(salt+pass)))
}

// 密码盐
func getPassSalt() (salt string) {
	var confMap map[string]string

	err := configs.GetConfig(consts.ENV_CONF, &confMap)
	if err != nil {
		return
	}

	var has bool
	salt, has = confMap[consts.ENV_CONF_PSALT]
	if !has {
		salt = ""
	}

	return
}
