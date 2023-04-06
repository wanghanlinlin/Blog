package main

import (
	"AuroraPixel/core"
	"AuroraPixel/flag"
	"AuroraPixel/routers"

	"github.com/sirupsen/logrus"
)

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
