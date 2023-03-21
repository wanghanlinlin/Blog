package models

import "time"

//文章标签关联表
type AriticleTag struct {
	AriticleID   uint         `gorm:"primaryKey"`
	ArticleModel ArticleModel `gorm:"foreignkey:AriticleID"`
	TagID        uint         `gorm:"primaryKey"`
	TagModel     TagModel     `gorm:"foreignkey:TagID"`
	CreatedAt    time.Time
}
