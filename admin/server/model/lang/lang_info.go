package lang

import "github.com/flipped-aurora/gin-vue-admin/server/global"

type LangInfo struct {
	global.COFFEE_MODEL
	Name string `json:"name" form:"name" gorm:"comment:字段名称"`
	Path string `json:"path" form:"path" gorm:"comment:页面路径"`
	ZH   string `json:"zh" form:"zh" gorm:"comment:中文"`
	EN   string `json:"en" form:"en" gorm:"comment:英文"`
}
