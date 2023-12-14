package dayz

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type DayzItemsGroup struct {
	global.GVA_MODEL
	Name              string         `json:"name"`  // 物品组名称
	Count             uint           `json:"count"` //数量
	DayzItemsID       uint           `json:"dayzItemsID"`
	DayzRecipesInput  []*DayzRecipes `gorm:"many2many:RecipesInputGroup;"`
	DayzRecipesOutput []*DayzRecipes `gorm:"many2many:RecipesOutputGroup;"`
}
