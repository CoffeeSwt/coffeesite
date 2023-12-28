package dayz

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type DayzRecipeOutput struct {
	global.COFFEE_MODEL
	Name        string       `json:"name" gorm:"comment:输出名称"`
	DayzItemID  uint         `json:"dayzItemID" gorm:"comment:输出物品ID"`
	DayzItem    DayzItem     `json:"dayzItem" gorm:"foreignKey:ID;references:DayzItemID;comment:输出物品"`
	Count       int          `json:"count" gorm:"comment:输出物品数量"`
	DayzRecipes []DayzRecipe `json:"dayzRecipes" gorm:"foreignKey:ID;comment:输入组"`
}
