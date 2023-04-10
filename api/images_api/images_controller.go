package imagesapi

import (
	"AuroraPixel/core/plugins"
	"AuroraPixel/core/res"
	imageservice "AuroraPixel/service/image_service"

	"github.com/gin-gonic/gin"
)

// 初始化controller
func (i *ImagesApi) InitController() {
	i.ImageService = imageservice.ImageServiceImpl{}
}

// 图片上传
func (i ImagesApi) Upload(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		res.ErrorWithCodeData(err.Error(), res.ArgumentError, c)
	}
	fh := form.File["images"]
	if len(fh) < 1 {
		res.ErrorWithMessage("images参数中的图片为空", c)
	}

	//上传图片
	result := i.ImageService.Upload(fh)
	res.Ok(result, "图片上传操作成功", c)
}

// 分页列表
func (i ImagesApi) PageList(c *gin.Context) {
	var ipage plugins.IPage
	err := c.ShouldBindQuery(&ipage)
	if err != nil {
		res.ErrorWithCodeData(err.Error(), res.ArgumentError, c)
	}
	//分页查询
	result := i.ImageService.PageList(&ipage)
	res.OkWithData(result, c)
}
