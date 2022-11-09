package main

import (
	"flag"
	"fmt"
	"gohub/bootstrap"
	btsconfig "gohub/config"
	"gohub/pkg/config"

	"github.com/gin-gonic/gin"
)

func init() {
	btsconfig.Initialize()
}

func main() {

	// 配置初始化，依赖命令行 --env 参数
	var env string
	flag.StringVar(&env, "env", "", "加载 .env 文件，如 --env=testing 加载的是 .env.testing 文件")
	flag.Parse()
	config.InitConfig(env)

	// 初始化 Logger
	bootstrap.SetupLogger()

	router := gin.New()

	bootstrap.SetupDB()
	bootstrap.SetupRoute(router)

	err := router.Run(":" + config.Get("app.port"))
	if err != nil {
		fmt.Println(err.Error())
	}
}
