package global

import (
	"AuroraPixel/config"
	"AuroraPixel/flag"

	"github.com/minio/minio-go/v7"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// 定义全局变量
var (
	//初始化配置结构体
	Config *config.Config
	DB     *gorm.DB
	Log    *logrus.Logger
	Minio  *minio.Client
	Option *flag.Option
)
