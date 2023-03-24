package plugins

import (
	"math"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type IPage struct {
	PageNum  int    `json:"pageNum" form:"pageNum"`   //当前页
	PageSize int    `json:"pageSize" form:"pageSize"` //当前页容量
	Key      string `json:"key" form:"key"`           //关键字
}

type PageResult struct {
	PageNum   int `json:"pageNum"`   //当前页
	PageSize  int `json:"pageSize"`  //每一页容量
	PageCount int `json:"pageCount"` //当前页数量
	Total     int `json:"total"`     //总数量
	PageTotal int `json:"pageTotal"` //页码总数量
	Data      any `json:"data"`      //内容
}

// 分页查询
func (ipage *IPage) Query(tx *gorm.DB, data any, order string) PageResult {
	//查询条件下总数
	var count int64
	tx.Model(&data).Count(&count)
	logrus.Info("count", count)
	//当总数为0返回空
	if count == 0 {
		return PageResult{}
	}
	//初始化页容量
	if ipage.PageSize < 1 {
		ipage.PageSize = 10
	}
	//初始化当前页码
	if ipage.PageNum < 1 {
		ipage.PageNum = 1
	}
	//判断是否超出最后一页
	lastPage := int(math.Ceil(float64(count) / float64(ipage.PageSize)))
	if ipage.PageNum > lastPage {
		ipage.PageNum = lastPage
	}

	//分页查询
	offset := (ipage.PageNum - 1) * ipage.PageSize
	tx.Offset(offset).Limit(ipage.PageSize)

	//排序
	if order != "" {
		tx.Order(order)
	}

	pageCount := tx.Find(&data).RowsAffected

	return PageResult{
		PageNum:   ipage.PageNum,
		PageSize:  ipage.PageSize,
		Total:     int(count),
		PageCount: int(pageCount),
		PageTotal: lastPage,
		Data:      data,
	}
}
