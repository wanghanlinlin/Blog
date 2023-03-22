package models

import "AuroraPixel/models/base"

//回复表
type FadeBackModel struct {
	base.MODEL
	Email        string `json:"email"`        //邮箱
	Content      string `json:"content"`      //内容
	ApplyContent string `json:"applyContent"` //回复内容
	IsApply      bool   `json:"isApply"`      //是否回复
}
