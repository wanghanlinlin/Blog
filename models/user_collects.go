package models

import "time"

//用户收藏中间表
type UserCollects struct {
	UserID       uint         `gorm:"primaryKey"`
	UserModel    UserModel    `gorm:"foreignkey:UserID"`
	ArticleID    uint         `gorm:"primaryKey"`
	ArticleModel ArticleModel `gorm:"foreignkey:ArticleID"`
	CreatedAt    time.Time
}
