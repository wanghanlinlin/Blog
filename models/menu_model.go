package models

import (
	"AuroraPixel/models/base"
	"AuroraPixel/models/ctype"
)

// 菜单列表
type MenuModel struct {
	base.MODEL
	MenuTitle    string        `json:"menuTitle"`                                //菜单标题
	MenuTitleEn  string        `json:"menuTitleEn"`                              //菜单标题国际化
	Slogan       string        `json:"slogan"`                                   //标语
	Abstract     ctype.Array   `gorm:"type:string" json:"abstract"`              //简介
	AbstractTime int           `json:"abstractTime"`                             //简介切换时间
	MenuBanners  []BannerModel `gorm:"many2many:menu_banners" json:"menuImages"` //菜单图片列表
	BannerTime   int           `json:"BannerTime"`                               //菜单图片切换时间
	Sort         int           `json:"sort"`                                     //菜单的顺序
}
