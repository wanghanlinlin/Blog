package config

import "AuroraPixel/config/conf"

type Config struct {
	conf.MysqlConfig  `yaml:"mysql"`
	conf.LoggerConfig `yaml:"logger"`
	conf.SystemConfig `yaml:"system"`
	conf.ImagesConfig `yaml:"images"`
}
