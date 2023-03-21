package models

import (
	"AuroraPixel/models/base"
	"AuroraPixel/models/ctype"
)

// 用户表
type UserModel struct {
	base.MODEL
	NickName       string           `gorm:"size:36" json:"nickName"`                                                      //昵称
	UserName       string           `gorm:"size:42" json:"userName"`                                                      //用户名
	Password       string           `gorm:"size:64" json:"password"`                                                      //密码
	Avatar         string           `gorm:"size:256" json:"avater"`                                                       //头像地址
	Email          string           `gorm:"size:128" json:"email"`                                                        //邮箱
	Tel            string           `gorm:"size:18" json:"tel"`                                                           //手机号
	Addr           string           `gorm:"size:128" json:"addr"`                                                         //地址
	OtherToken     string           `gorm:"size:128" json:"otherToken"`                                                   //其他平台的token
	IP             string           `gorm:"size:12" json:"ip"`                                                            //ip地址
	Role           ctype.Role       `gorm:"size:4;default:3" json:"role"`                                                 //角色
	SignSource     ctype.SignSource `gorm:"type=smallint(6)" json:"signSource"`                                           //注册来源
	ArticleModels  []ArticleModel   `gorm:"foreignkey:UserID" json:"-"`                                                   //发布文章列表
	CollectsModels []ArticleModel   `gorm:"many2many:auth2_collects;foreignkey:UserID;joinReferences:ArticleID" json:"-"` //收集文章列表
}
