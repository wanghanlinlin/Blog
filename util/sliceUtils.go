package util

import (
	"errors"
)

// 判断切片内容是否存在
func InSlice(slice []string, value string) (result bool, err error) {
	//判断切边长度
	result = false
	if len(slice) < 1 {
		err = errors.New("该切片为空")
	}
	for _, v := range slice {
		if v == value {
			result = true
		}
	}
	return
}
