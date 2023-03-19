package core

import (
	"AuroraPixel/global"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 初始化数据库连接
func InitDb() {
	global.DB = mysqlConnection()
	global.Log.Infof("%v数据库连接成功!", global.Config.MysqlConfig.Host)
}

// mysql连接器
func mysqlConnection() *gorm.DB {
	if global.Config.MysqlConfig.Host == "" {
		//使用全局logger处理器
		global.Log.Panic("数据库未进行配置!")
	}

	//数据库dsn地址
	dsn := global.Config.Dsn()
	if dsn == "" {
		//使用全局logger处理器
		global.Log.Panic("数据库dsn地址错误!")
	}

	var mysqlLogger logger.Interface
	//判断日志级别
	if global.Config.MysqlConfig.LogLevel == "debug" {
		//打印所有sql
		mysqlLogger = logger.Default.LogMode(logger.Info)
	} else {
		//只打印error
		mysqlLogger = logger.Default.LogMode(logger.Error)
	}

	//数据库连接客户端
	connection, error := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	if error != nil {
		global.Log.Errorf("%v: 数据库初始化失败:%v", global.Config.Prefix, error)
	}

	sqlDb, _ := connection.DB()
	sqlDb.SetMaxIdleConns(10)               // 最大空闲连接数
	sqlDb.SetMaxOpenConns(100)              // 最多容纳数量
	sqlDb.SetConnMaxIdleTime(time.Hour * 4) // 最大复用时间

	return connection
}
