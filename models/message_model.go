package models

import "AuroraPixel/models/base"

//消息模型
type MessageModel struct {
	base.MODEL
	SendUserID       uint      `gorm:"primarykey" json:"sendUserID"` //发送人id
	SendUserModel    UserModel `gorm:"foreignkey:SendUserID" json:"-"`
	SendUserNickName string    `json:"sendUserNickName"` //发送人昵称
	SendUserAvatar   string    `json:"sendUserAvatar"`   //发送人头像

	RevUserID       uint      `gorm:"primarykey" json:"revUserID"` //收件人id
	RevUserModel    UserModel `gorm:"foreignkey:RevUserID" json:"-"`
	RevUserNickName string    `json:"revUserNickName"` //收送人昵称
	RevUserAvatar   string    `json:"revUserAvatar"`   //收送人头像

	IsRead  bool   `json:"isRead"`  //是否已读
	Content string `json:"content"` //内容
}
