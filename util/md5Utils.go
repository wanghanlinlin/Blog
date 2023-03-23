package util

import (
	"crypto/md5"
	"fmt"
)

// MD5计算
func MD5(data []byte) string {
	sum := md5.Sum(data)
	return fmt.Sprintf("%x\n", sum)
}
