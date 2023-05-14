package routers

import (
	"AuroraPixel/api"
	"AuroraPixel/global"

	"github.com/gin-gonic/gin"
)

func (r Routers) EchoRouters(c *gin.RouterGroup) {
	EchoApi := api.GroupApi.EchoApi

	//创建连接
	c.GET("ws/echoDemo", EchoApi.Connect)
	c.GET("ws/echo", func(ctx *gin.Context) { EchoApi.EchoService(global.Sub, ctx) })
}
