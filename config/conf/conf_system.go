package conf

import "strconv"

// 系统配置
type SystemConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Env  string `yaml:"env"`
}

func (s SystemConfig) ServerAddress() string {
	return s.Host + ":" + strconv.Itoa(s.Port)
}
