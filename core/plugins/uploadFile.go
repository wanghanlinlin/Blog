package plugins

import (
	"AuroraPixel/global"
	"context"
	"mime/multipart"

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
func (u *UplodaImages) UploadFile() (UploadImagesResult, error) {
	filename := u.File.Filename
	fileSize := u.File.Size

	//打开文件流
	uplodaFile, openErr := u.File.Open()
	if openErr != nil {
		return UploadImagesResult{}, openErr
	}
	defer uplodaFile.Close()

	//文件上传
	uploadInfo, err := global.Minio.PutObject(context.Background(), u.BucketName, filename, uplodaFile, fileSize, minio.PutObjectOptions{ContentType: u.ContentType})
	if err != nil {
		return UploadImagesResult{}, err
	}
	logrus.Infoln("mino图片上传成功", uploadInfo)

	return UploadImagesResult{
		Path:     "http://" + global.Config.MinioConfig.Endpoint + "/" + uploadInfo.Bucket + "/" + uploadInfo.Key,
		FileName: uploadInfo.Key,
		Size:     uploadInfo.Size,
	}, nil
}
