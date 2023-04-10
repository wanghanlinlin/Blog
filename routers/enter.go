package routers

import (
	_ "AuroraPixel/docs"
	"AuroraPixel/global"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	//swager
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//路由组
	routerGroup := router.Group("api")

	//初始化路由体
	routers := Routers{}

	//系统api
	routers.SystemRouter(routerGroup)
	//图片api
	routers.imagesRouter(routerGroup)

	return router
}
