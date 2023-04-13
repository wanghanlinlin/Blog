package flag

import (
	sys_flag "flag"
)

type Option struct {
	DB      bool
	Setting string
}

// 解析命令行
func Parse() *Option {
	//db命令行
	db := sys_flag.Bool("db", false, "初始化数据库")
	//运行环境命令行
	setting := sys_flag.String("setting", "", "配置文件地址")
	//解析命令行
	sys_flag.Parse()
	return &Option{
		DB:      *db,
		Setting: *setting,
	}
}

// 是否停止web项目
func IsWebStop(option Option) bool {
	return option.DB
}
