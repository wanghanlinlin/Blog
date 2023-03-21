package models

import (
	"AuroraPixel/models/base"
	"AuroraPixel/models/ctype"
)

type ArticleModel struct {
	base.MODEL
	Title         string         `gorm:"size:64" json:"title"`                    //文章标题
	Abstract      string         `json:"abstract"`                                //文章简介
	Content       string         `json:"content"`                                 //文章内容
	LookCount     int            `json:"lookCount"`                               //浏览量
	CommentCount  int            `json:"commentCount"`                            //评论量
	DiggCount     int            `json:"digg_count"`                              //点赞量
	CollectsCount int            `json:"collectsCount"`                           //收藏量
	TagModels     []TagModel     `gorm:"many2many:article_tag" json:"tag_models"` //文章标签
	CommentModels []CommentModel `gorm:"foreignkey:ArticleID" json:"-"`           //文章评论列表
	UserModel     UserModel      `gorm:"foreignkey:UserID" json:"-"`              //文章作者
	UserID        uint           `json:"userID"`                                  //用户id
	Category      string         `gorm:"size:20" json:"category"`                 //文章分类
	Source        string         `json:"source"`                                  //文章来源
	Banner        BannerModel    `gorm:"foreignkey:BannerID" json:"-"`            //文章封面
	BannerID      uint           `json:"bannerID"`                                //文章封面ID
	BannerPath    string         `json:"banner_path"`                             //文章封面地址
	NickName      string         `gorm:"size:42" json:"nick_name"`                //发布文章的用户昵称
	Tags          ctype.Array    `gorm:"type:string;size:64" json:"tags"`         //文章标签
}
