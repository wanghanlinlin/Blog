package core

import (
	"AuroraPixel/global"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/sirupsen/logrus"
)

func InitMinio() {
	global.Minio = minioClient()
	logrus.Info("minio客户端初始化成功!")
}

func minioClient() *minio.Client {
	endpoint := global.Config.MinioConfig.Endpoint
	accessKeyID := global.Config.MinioConfig.AccessKeyID
	secretAccessKey := global.Config.MinioConfig.SecretAccessKey
	useSSL := global.Config.MinioConfig.UseSSL

	//初始化客户端
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		logrus.Fatalln(err.Error())
	}
	return minioClient
}
