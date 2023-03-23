package conf

//上传图片配置
type ImagesConfig struct {
	Path string `yaml:"path"` //文件路径
	Size int    `yaml:"size"` //文件大小
}
