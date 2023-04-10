package main

import (
	"AuroraPixel/core"
	"AuroraPixel/flag"
	"AuroraPixel/routers"

	"github.com/sirupsen/logrus"
)

// @title AuroraPixel
// @version 1.0
// @description AuroraPixel系统API文档
// @termsOfService https://www.test.com

// @contact.name www.test.com
// @contact.url https://www.test.com
// @contact.email me@test.me

// @license.name AuroraPixel
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8080
// @BasePath
func main() {
	//加载配置文件
	core.InitConf()
	//初始化日志
	core.InitLogger()
	//加载数据库
	core.InitDb()
	//加载minio
	core.InitMinio()
	//数据库迁移
	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	}

	//运行服务
	erro := routers.Run()
	if erro != nil {
		logrus.Errorln(erro)
	}
}
