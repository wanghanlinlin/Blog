package imagesapi

import (
	"AuroraPixel/core/res"

	"github.com/gin-gonic/gin"
)

// 图片上传
func (i ImagesApi) Upload(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		res.ErrorWithCodeData(err.Error(), res.ArgumentError, c)
		return
	}
	fh := form.File["images"]
	if len(fh) < 1 {
		res.ErrorWithMessage("images参数中的图片为空", c)
		return
	}
	res.Ok(fh, "图片上传成功!", c)
}
