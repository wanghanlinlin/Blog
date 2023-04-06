package core

import (
	"context"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/sirupsen/logrus"
)

func InitMinio() {
	ctx := context.Background()
	endpoint := "124.222.46.195:9090"
	accessKeyID := "wang_test"
	secretAccessKey := "wang12345678"
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		logrus.Fatalln(err.Error())
	}

	logrus.Infof("minio客户端初始成功:%v", minioClient)

	//创建桶
	bucketName := "mymusic"
	location := "us-east-1"

	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			logrus.Printf("数据桶 %s 已经存在\n", bucketName)
		} else {
			logrus.Fatalln(err)
		}
	} else {
		logrus.Printf("成功创建数据桶 %s\n", bucketName)
	}

	// // Upload the zip file
	// objectName := "golden-oldies.zip"
	// filePath := "/tmp/golden-oldies.zip"
	// contentType := "application/zip"

	// // Upload the zip file with FPutObject
	// info, err := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	// if err != nil {
	// 		log.Fatalln(err)
	// }

	// log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)
}
