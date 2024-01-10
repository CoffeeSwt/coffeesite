package dayz

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type DayzItemImg struct {
	global.GVA_MODEL
	Name     string `json:"name" form:"name" gorm:"comment:图片名称"`         //图片名称
	Path     string `json:"path" form:"path" gorm:"comment:图片路径"`         //图片路径
	Preview  bool   `json:"preview" form:"preview" gorm:"comment:是否为预览图"` //是否为预览图
	ItemName string `json:"itemName" form:"itemName" gorm:"comment:物品名称"`
	Key      string `json:"key" gorm:"comment:编号"` // 编号
}
