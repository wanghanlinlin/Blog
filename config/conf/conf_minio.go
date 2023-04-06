package conf

//minio配置
type MinioConfig struct {
	Endpoint        string `yaml:"endpoint"`        //站点
	AccessKeyID     string `yaml:"accessKeyID"`     //连接密钥
	SecretAccessKey string `yaml:"secretAccessKey"` //连接密钥密码
	UseSSL          bool   `yaml:"useSSL"`          //是否https连接
}
