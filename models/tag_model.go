package models

import "AuroraPixel/models/base"

//标签
type TagModel struct {
	base.MODEL
	TagName  string         `gorm:"size:16" json:"tagName"`         //标签名称
	Articles []ArticleModel `gorm:"many2many:article_tag" json:"-"` // 关联该标签的文章列表
}
