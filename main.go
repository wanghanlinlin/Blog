package main

import (
	"AuroraPixel/core"
	"AuroraPixel/routers"
)

func main() {
	//加载配置文件
	core.InitConf()
	//初始化日志
	core.InitLogger()
	//加载数据库
	core.InitDb()
	//运行服务
	routers.Run()
}
