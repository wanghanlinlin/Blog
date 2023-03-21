package routers

import (
	"AuroraPixel/global"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// router运行
func Run() error {
	logrus.Infof("%v-服务启动", global.Config.SystemConfig.Port)
	return initGin().Run(global.Config.ServerAddress())
}

// router结构体
type Routers struct {
}

// 初始化路由
func initGin() *gin.Engine {
	//配置Gin开发环境
	gin.SetMode(global.Config.Env)

	//路由配置
	router := gin.Default()

	//路由组
	routerGroup := router.Group("api")

	//初始化路由体
	routers := Routers{}

	//系统api
	routers.SystemRouter(routerGroup)

	return router
}
