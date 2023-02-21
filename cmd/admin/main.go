package main

import (
	"flag"
	"os"
	"redisadmin/internal/accessControl"
	"redisadmin/internal/configs"
	"redisadmin/internal/consts"
	initweb "redisadmin/internal/initWeb"
	"redisadmin/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	var port string
	var init bool
	flag.StringVar(&port, "port", "12345", "监听端口")
	flag.BoolVar(&init, "init", false, "是否初始化网址")
	flag.Parse()

	// 执行程序
	cmd := os.Args[0]

	if init {
		initweb.InitWeb(cmd)
		return
	}

	accessControl.NewEnforcer()

	ginMode := configs.GetEnvVal(consts.ENV_CONF_GIN_MODE)
	gin.SetMode(ginMode)

	r := router.SetupRouter()

	r.Run(":" + port)
}
