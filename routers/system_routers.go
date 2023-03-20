package routers

import (
	"AuroraPixel/api"

	"github.com/gin-gonic/gin"
)

func (r Routers) SystemRouter(c *gin.RouterGroup) {
	systemapi := api.GroupApi.SystemApi
	c.GET("", systemapi.SystemApiVO)
}
