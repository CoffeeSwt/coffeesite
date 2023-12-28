package siteuser

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/satori/go.uuid"
)

type SiteUser struct {
	global.GVA_MODEL
	UUID        uuid.UUID `json:"uuid" gorm:"comment:用户UUID"`
	NickName    string    `json:"nickName" gorm:"comment:用户昵称"`
	UserName    string    `json:"userName" gorm:"comment:用户名"`
	PassWord    string    `json:"passWord" gorm:"comment:密码"`
	HeaderImg   string    `json:"headerImg" gorm:"comment:用户头像"`
	Intro       string    `json:"intro" gorm:"comment:个人简介"`
	Status      int       `json:"status" gorm:"状态 0待激活 1已激活 2已停用"`
	Mail        string    `json:"mail" gorm:"邮箱"`
	PhoneNumber string    `json:"phoneNumber" gorm:"comment:用户名"`
}
