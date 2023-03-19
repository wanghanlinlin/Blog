package conf

// logger日志配置
type LoggerConfig struct {
	Level        string `yaml:"level"`
	Prefix       string `yaml:"prefix"`
	Director     string `yaml:"director"`
	ShowLine     bool   `yaml:"show-line"`
	LogInConsole bool   `yaml:"log-in-console"`
}
