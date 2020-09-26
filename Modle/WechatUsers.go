package Modle

import "ginApi/Entity"

type WechatUserDetail struct {
	Entity.WechatUsers
	Level string `json:"level"`
}

