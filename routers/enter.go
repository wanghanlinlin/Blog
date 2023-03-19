package routers

import (
	"AuroraPixel/global"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Run() {
	logrus.Infof("%v-服务启动", global.Config.SystemConfig.Port)
	initGin().Run(global.Config.ServerAddress())
}

// 初始化路由
func initGin() *gin.Engine {
	gin.SetMode(global.Config.Env)
	router := gin.Default()
	router.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "欢迎来到Aurora!",
		})
	})
	return router
}
