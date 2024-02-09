package dayz

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type DayzItemImg struct {
	global.GVA_MODEL
	Name     string `json:"name" form:"name" gorm:"comment:图片名称"` //图片名称
	Path     string `json:"path" form:"path" gorm:"comment:图片路径"` //图片路径
	ItemName string `json:"itemName" form:"itemName" gorm:"comment:物品名称"`
	FileName string `json:"fileName" form:"fileName" gorm:"comment:文件名称"`
}
