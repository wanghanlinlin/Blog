package conf

import "strconv"

// mysql数据库配置
type MysqlConfig struct {
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	Db        string `yaml:"db"`
	User      string `yaml:"user"`
	Password  string `yaml:"password"`
	Charset   string `yaml:"charset"`
	ParseTime string `yaml:"parseTime"`
	Loc       string `yaml:"loc"`
	LogLevel  string `yaml:"log_level"`
}

//获取dsn连接地址：dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
func (m *MysqlConfig) Dsn() string {
	dsn := m.User + ":" + m.Password + "@tcp(" + m.Host + ":" + strconv.Itoa(m.Port) + ")/" + m.Db + "?charset=" + m.Charset + "&parseTime=" + m.ParseTime + "&loc=" + m.Loc
	return dsn
}
