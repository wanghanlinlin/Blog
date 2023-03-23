package base

import (
	"AuroraPixel/util"
	"time"

	"gorm.io/gorm"
)

type Time time.Time

const (
	timeFormart = "2006-01-02 15:04:05"
)

// time 时间格式json解析
func (t *Time) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormart+`"`, string(data), time.Local)
	*t = Time(now)
	return
}

// time 时间格式json压缩
func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormart)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormart)
	b = append(b, '"')
	return b, nil
}

// time 时间格式字符串
func (t Time) String() string {
	return time.Time(t).Format(timeFormart)
}

// 基础结构体
type MODEL struct {
	ID        int64     `gorm:"primarykey" json:"id"` //主键id
	CreatedAt time.Time `json:"created_at"`           //创建时间
	UpdatedAt time.Time `json:"update_at"`            //更新时间
}

// 雪花id
func (u *MODEL) BeforeCreate(tx *gorm.DB) (err error) {
	w := util.NewWorker(11, 11)
	u.ID = w.GetID()
	return
}
