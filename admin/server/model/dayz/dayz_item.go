package dayz

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type DayzItem struct {
	global.GVA_MODEL
	Name        string        `json:"name" form:"name" gorm:"comment:物品名称"`
	MaxQuantity int           `json:"maxQuantity" form:"maxQuantity" gorm:"comment:堆叠上限"`
	Category    int           `json:"category" form:"category" gorm:"comment:物品类别"`
	Info        string        `json:"info" form:"info" gorm:"comment:物品信息"`
	Imgs        []DayzItemImg `json:"imgs" form:"imgs" gorm:"foreignKey:ItemName"`
}
