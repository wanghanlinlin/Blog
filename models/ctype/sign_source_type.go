package ctype

import "encoding/json"

//注册源
type SignSource int

var (
	SignGithub SignSource = 1 //github
	SignEmail  SignSource = 2 //Email
)

func (s SignSource) String() string {
	switch s {
	case SignGithub:
		return "github"
	case SignEmail:
		return "email"
	}
	return "其他"
}

func (s SignSource) MarshalJson() ([]byte, error) {
	return json.Marshal(s.String())
}
