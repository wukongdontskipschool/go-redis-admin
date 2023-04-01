package initweb

import (
	"log"
	"os"
	"redisadmin/internal/accessControl"
	"redisadmin/internal/configs"
	"redisadmin/internal/jobs/admin"
)

func InitWeb(cmd string) {
	dbConf := "internal/configs/databases/redisAdmin.yaml"
	if _, err := os.Stat(dbConf); err != nil {
		if os.IsNotExist(err) {
			// 配置文件不存在 创建配置文件
			err := configs.CreateEnv()
			if err != nil {
				log.Println("创建env配置失败")
				panic(err)
			}
			err = configs.CreateDatabase()
			if err != nil {
				log.Println("创建database配置失败")
				panic(err)
			}
			err = configs.CreateCasbin()
			if err != nil {
				log.Println("创建casbin配置失败")
				panic(err)
			}

			log.Println("初始配置文件创建成功，请配置" + dbConf)
			log.Println("请再次执行" + cmd + " -init，初始化数据表")
			return
		} else {
			log.Println("判断配置是否存在出错")
			panic(err)
		}
	}

	// 权限表
	accessControl.NewEnforcer()
	// 建表
	admin.Migrate()
	log.Println("初始化完毕，超级用户账号：admin初始密码：123456")
	log.Println("执行" + cmd + " -port=12367指定监听端口")
}
