package imageservice

import (
	"AuroraPixel/core/plugins"
	"mime/multipart"
)

// 符合场景的图片格式
var CheckImageType = []string{
	"jpg",
	"png",
	"tif",
	"gif",
	"bmp",
	"svg",
}

// 图片返回视图
type ImagesVO struct {
	Path      string `json:"path"`      //路径
	FileName  string `json:"fileName"`  //文件名
	IsSuccess bool   `json:"isSuccess"` //是否成功
	Message   string `json:"message"`   //消息内容
}

type ImageService interface {
	//文件上传
	Upload(files []*multipart.FileHeader) []ImagesVO
	//分页查询
	PageList(ipage *plugins.IPage) plugins.PageResult
}
