package plugins

import (
	"AuroraPixel/core/res"
	"AuroraPixel/global"
	"context"
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/sirupsen/logrus"
)

type UplodaImages struct {
	BucketName  string                //桶名
	File        *multipart.FileHeader //FileHeader
	ContentType string                //内容体："例application/octet-stream"
}

type UploadImagesResult struct {
	Path     string //路径
	FileName string //文件名
	Size     int64  //大小
}

// 文件上传
func (u *UplodaImages) UploadFile(c *gin.Context) UploadImagesResult {
	filename := u.File.Filename
	fileSize := u.File.Size
	uplodaFile, openErr := u.File.Open()
	if openErr != nil {
		res.ErrorWithMessage(openErr.Error(), c)
	}
	defer uplodaFile.Close()
	uploadInfo, err := global.Minio.PutObject(context.Background(), u.BucketName, filename, uplodaFile, fileSize, minio.PutObjectOptions{ContentType: u.ContentType})
	if err != nil {
		res.ErrorWithMessage(err.Error(), c)
	}
	logrus.Infoln("mino图片上传成功", uploadInfo)

	return UploadImagesResult{
		Path:     "http://" + global.Config.MinioConfig.Endpoint + "/" + uploadInfo.Bucket + "/" + uploadInfo.Key,
		FileName: uploadInfo.Key,
		Size:     uploadInfo.Size,
	}
}
