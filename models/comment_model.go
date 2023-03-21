package models

import "AuroraPixel/models/base"

//评论
type CommentModel struct {
	base.MODEL
	SubComments        []*CommentModel `gorm:"foreignkey:ParentCommentID" json:"subComments"`   //子评论表
	ParentCommentModel *CommentModel   `gorm:"foreignkey:ParentCommentID" json:"parentComment"` //父级评论
	ParentCommentID    *uint           `json:"parentCommentID"`                                 //父级评论ID

	Content      string       `json:"content"`                             //评论内容
	LookCount    int          `json:"lookCount"`                           //浏览量
	CommentCount int          `json:"commentCount"`                        //子评论量
	DiggCount    int          `json:"digg_count"`                          //点赞量
	Article      ArticleModel `gorm:"foreignkey:ArticleID" json:"article"` //关联的文章
	ArticleID    uint         `json:"articleId"`                           //文章ID
	User         UserModel    `json:"user"`                                //关联用户
	UserID       uint         `json:"userID"`                              //用户id
}
