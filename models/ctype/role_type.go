package ctype

import "encoding/json"

//定义Role角色
type Role int

//定义角色枚举
var (
	Admin       Role = 1 //管理员
	User        Role = 2 //普通用户
	Visitor     Role = 3 //访客
	DisableUser Role = 4 //被禁用户
)

//解析角色
func (r Role) String() string {
	switch r {
	case Admin:
		return "管理员"
	case User:
		return "普通用户"
	case Visitor:
		return "访客"
	case DisableUser:
		return "被禁用户"
	}
	return "其他"
}

//json转换
func (r Role) MarshalJson() ([]byte, error) {
	return json.Marshal(r.String())
}
