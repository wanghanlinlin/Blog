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
// @Tags 图片
// @Summary 图片上传
// @Description 图片上传
// @produce json
// @param images formData file true "图片"
// @success 200 {object} res.Response{Data=[]imageservice.ImagesVO} "Success Response"
// @Router /api/images/upload [post]
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
// @Tags 图片
// @Summary 分页列表
// @Description 分页查询数据
// @produce json
// @param pageNum query int true "当前页" defualt(0)
// @param pageSize query int true "页容量" defualt(10)
// @success 200 {object} res.Response{Data=plugins.PageResult{Data=[]models.BannerModel}} "Success Response"
// @Router /api/images/page [get]
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
