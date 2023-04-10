package imagesapi

import (
	"AuroraPixel/core/plugins"
	"AuroraPixel/core/res"
	"AuroraPixel/global"
	"AuroraPixel/models"
	"AuroraPixel/util"
	"io"
	"path"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
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

	//返回结构
	var result = make([]ImagesVO, 0)
	for _, value := range fh {
		filename := value.Filename
		fileSize := value.Size

		//判断文件名
		filenameSplit := strings.Split(filename, ".")
		filenameSuffix := filenameSplit[len(filenameSplit)-1]
		checkResult, _ := util.InSlice(CheckImageType, strings.ToLower(filenameSuffix))
		if !checkResult {
			result = append(result, ImagesVO{
				Path:      "",
				FileName:  filename,
				IsSuccess: false,
				Message:   "图片格式不正确!",
			})
			continue
		}

		//判断文件大小是否合适
		if (fileSize / 1024 / 1024) > (int64(global.Config.ImagesConfig.Size)) {
			result = append(result, ImagesVO{
				Path:      "",
				FileName:  filename,
				IsSuccess: false,
				Message:   filename + ":文件大小超过" + strconv.Itoa(global.Config.ImagesConfig.Size) + "MB",
			})
			continue
		}

		//计算文件内容的hash
		file, _ := value.Open()
		ioByte, _ := io.ReadAll(file)
		md5String := util.MD5(ioByte)
		defer file.Close()

		//查询图片是否存在
		var queryBanner models.BannerModel
		global.DB.Where(&models.BannerModel{Hash: md5String}).Find(&queryBanner)
		if queryBanner.Hash != "" {
			result = append(result, ImagesVO{
				Path:      queryBanner.Path,
				FileName:  queryBanner.Name,
				IsSuccess: true,
				Message:   "图片已经存在",
			})
			continue
		}

		//上传图片
		path := path.Join(global.Config.ImagesConfig.Path, filename)

		//minio上传图片
		uploadImage := plugins.UplodaImages{
			BucketName:  global.Config.MinioConfig.BucketName,
			File:        value,
			ContentType: "application/octet-stream",
		}
		uploadImagesResult := uploadImage.UploadFile(c)
		logrus.Info(uploadImagesResult)

		//添加返回成功值
		result = append(result, ImagesVO{
			Path:      path,
			FileName:  filename,
			IsSuccess: true,
			Message:   "图片上传成功",
		})

		//图片入库
		global.DB.Create(&models.BannerModel{
			Path: path,
			Hash: md5String,
			Name: filename,
		})
	}
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
	var BannerModel models.BannerModel = models.BannerModel{}
	result := plugins.PageQuery(BannerModel, "created_at desc", ipage)
	res.OkWithData(result, c)
}
