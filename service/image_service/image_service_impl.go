package imageservice

import (
	"AuroraPixel/core/plugins"
	"AuroraPixel/global"
	"AuroraPixel/models"
	"AuroraPixel/util"
	"io"
	"mime/multipart"
	"strconv"
	"strings"
)

type ImageServiceImpl struct{}

// 上传文件
func (ImageServiceImpl) Upload(files []*multipart.FileHeader) []ImagesVO {
	//返回结构
	var result = make([]ImagesVO, 0)

	for _, value := range files {
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

		//minio上传图片
		uploadImage := plugins.UplodaImages{
			BucketName:  global.Config.MinioConfig.BucketName,
			File:        value,
			ContentType: "application/octet-stream",
		}
		uploadImagesResult, uploadErro := uploadImage.UploadFile()
		if uploadErro != nil {
			result = append(result, ImagesVO{
				Path:      "",
				FileName:  filename,
				IsSuccess: false,
				Message:   uploadErro.Error(),
			})
			continue
		}

		//添加返回成功值
		result = append(result, ImagesVO{
			Path:      uploadImagesResult.Path,
			FileName:  uploadImagesResult.FileName,
			IsSuccess: true,
			Message:   "图片上传成功",
		})

		//图片入库
		global.DB.Create(&models.BannerModel{
			Path: uploadImagesResult.Path,
			Hash: md5String,
			Name: uploadImagesResult.FileName,
		})
	}
	return result
}

// 分页查询
func (ImageServiceImpl) PageList(ipage *plugins.IPage) plugins.PageResult {
	var BannerModel models.BannerModel = models.BannerModel{}
	return plugins.PageQuery(BannerModel, "created_at desc", *ipage)
}
