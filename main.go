package main

import (
	echoapi "AuroraPixel/api/echo_api"
	"AuroraPixel/core"
	"AuroraPixel/flag"
	"AuroraPixel/flag/option"
	"AuroraPixel/global"
	"AuroraPixel/routers"

	"github.com/sirupsen/logrus"
)

// @title AuroraPixel
// @version 1.0
// @description AuroraPixel系统API文档
// @termsOfService https://github.com/AuroraPixel

// @contact.name AuroraPixel
// @contact.url https://github.com/AuroraPixel
// @contact.email wanghanlinlin@126.com

// @license.name AuroraPixel
// @license.url https://github.com/AuroraPixel

// @host 127.0.0.1:8080
// @BasePath
func main() {
	//解析命令行
	global.Option = flag.Parse()

	//加载配置文件
	core.InitConf()
	//初始化日志
	core.InitLogger()
	//加载数据库
	core.InitDb()
	//加载minio
	core.InitMinio()
	//加载websocket消息总线
	global.Sub = echoapi.NewHub()
	go global.Sub.HubRun()
	//是否停止项目
	if flag.IsWebStop(*global.Option) {
		option.SwitchOption(*global.Option)
		return
	}

	//运行服务
	erro := routers.Run()
	if erro != nil {
		logrus.Errorln(erro)
	}
}
