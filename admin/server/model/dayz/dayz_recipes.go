package dayz

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type DayzRecipes struct {
	global.GVA_MODEL
	Name               string            `json:"name"` //合成表名称
	RecipesInputGroup  []*DayzItemsGroup `gorm:"many2many:RecipesInputGroup;"`
	RecipesOutputGroup []*DayzItemsGroup `gorm:"many2many:RecipesOutputGroup;"`
}
