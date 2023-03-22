package systemapi

import (
	"AuroraPixel/core/res"

	"github.com/gin-gonic/gin"
)

func (s *SystemApi) SystemApiVO(ctx *gin.Context) {
	res.OkWithMessage("欢迎来到AURORA!", ctx)
}
