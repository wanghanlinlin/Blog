package flag

import sys_flag "flag"

type Option struct {
	DB bool
}

// 解析命令行
func Parse() Option {
	db := sys_flag.Bool("db", false, "初始化数据库")
	//解析命令行
	sys_flag.Parse()
	return Option{
		DB: *db,
	}
}

// 是否停止web项目
func IsWebStop(option Option) bool {
	return option.DB
}

// 根据命令执行不同的函数
func SwitchOption(option Option) {
	if option.DB {
		Makemigrations()
	}
}
