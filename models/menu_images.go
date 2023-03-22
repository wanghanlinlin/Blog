package models

import "time"

type MenuImages struct {
	MenuID      uint        `gorm:"primaryKey"`
	MenuModel   MenuModel   `gorm:"foreignkey:MenuID"`
	BannerID    uint        `gorm:"primaryKey"`
	BannerModel BannerModel `gorm:"foreignkey:BannerID"`
	CreatedAt   time.Time
}
