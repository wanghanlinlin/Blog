package option

import (
	"AuroraPixel/flag"
	"AuroraPixel/global"
	"AuroraPixel/models"

	"github.com/sirupsen/logrus"
)

func Makemigrations() {
	logrus.Info("开始迁移数据库")

	global.DB.SetupJoinTable(&models.UserModel{}, "CollectsModels", &models.UserCollects{})

	global.DB.SetupJoinTable(&models.ArticleModel{}, "TagModels", &models.AriticleTag{})

	global.DB.SetupJoinTable(&models.MenuModel{}, "MenuBanners", &models.MenuImages{})

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
		&models.MenuModel{},
		&models.MenuImages{},
		&models.MessageModel{},
		&models.FadeBackModel{},
	)
	if err != nil {
		logrus.Errorln("数据库初始化失败!", err)
		return
	}

	logrus.Info("数据库初始化成功!")
}

// 根据命令执行不同的函数
func SwitchOption(option flag.Option) {
	if option.DB {
		Makemigrations()
	}
}
