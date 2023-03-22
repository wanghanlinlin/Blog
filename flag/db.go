package flag

import (
	"AuroraPixel/global"
	"AuroraPixel/models"

	"github.com/sirupsen/logrus"
)

func Makemigrations() {
	logrus.Info("开始迁移数据库")

	global.DB.SetupJoinTable(&models.UserModel{}, "CollectsModels", &models.UserCollects{})

	global.DB.SetupJoinTable(&models.ArticleModel{}, "TagModels", &models.AriticleTag{})

	//迁移数据库
	err := global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		Migrator().AutoMigrate(
		&models.UserModel{},
		&models.ArticleModel{},
		&models.AriticleTag{},
		&models.UserCollects{},
		&models.BannerModel{},
		&models.CommentModel{},
		&models.TagModel{},
	)
	if err != nil {
		logrus.Errorln("数据库初始化失败!", err)
		return
	}

	logrus.Info("数据库初始化成功!")
}
