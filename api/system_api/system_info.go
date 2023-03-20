package systemapi

import (
	"AuroraPixel/model/res"

	"github.com/gin-gonic/gin"
)

func (s *SystemApi) SystemApiVO(ctx *gin.Context) {
	res.OkWithMessage("欢迎来到AURORA!", ctx)
}
