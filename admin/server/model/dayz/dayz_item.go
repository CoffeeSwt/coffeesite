package dayz

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type DayzItem struct {
	global.COFFEE_MODEL
	Name        string        `json:"name" form:"name" gorm:"comment:物品名称"`               // 物品名称
	MaxQuantity int           `json:"maxQuantity" form:"maxQuantity" gorm:"comment:堆叠上限"` //堆叠上限
	Info        string        `json:"info" form:"info" gorm:"comment:物品信息"`               //物品信息
	Imgs        []DayzItemImg `json:"imgs" form:"imgs" gorm:"foreignKey:ItemID"`          //物品图片
}