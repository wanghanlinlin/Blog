package models

import "AuroraPixel/models/base"

//图片背景
type BannerModel struct {
	base.MODEL
	Path string `json:"path"`                //图片描述
	Hash string `json:"hash"`                //hash值判断图片是否唯一
	Name string `gorm:"size:38" json:"name"` //图片名称
}
