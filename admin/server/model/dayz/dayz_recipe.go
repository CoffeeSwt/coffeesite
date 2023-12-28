package dayz

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type DayzRecipe struct {
	global.COFFEE_MODEL
	Name string `json:"name" gorm:"comment:输入组名称"`
	Info string `json:"info" gorm:"comment:输入组信息"`

	InputItemID1  uint     `json:"inputItemID1" gorm:"comment:输入物品1的ID"`
	DayzItem1     DayzItem `json:"dayzItem1" gorm:"foreignKey:ID;references:InputItemID1;comment:输入物品1"`
	InputCount1   int      `json:"inputCount1" gorm:"comment:输入物品1的数量"`
	InputDestroy1 bool     `json:"inputDestroy1" gorm:"输入物品1合成后是否销毁"`
	InputHealthy1 int      `json:"inputHealthy1" gorm:"输入物品1合成后健康值损耗"`

	InputItemID2  uint     `json:"inputItemID2" gorm:"comment:输入物品2的ID"`
	DayzItem2     DayzItem `json:"dayzItem2" gorm:"foreignKey:ID;references:InputItemID2;comment:输入物品2"`
	InputCount2   int      `json:"inputCount2" gorm:"comment:输入物品2的数量"`
	InputDestroy2 bool     `json:"inputDestroy2" gorm:"输入物品2合成后是否销毁"`
	InputHealthy2 int      `json:"inputHealthy2" gorm:"输入物品2合成后健康值损耗"`
}
