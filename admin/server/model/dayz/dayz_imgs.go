package dayz

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type DayzItemsImgs struct {
	global.GVA_MODEL
	Name        string `json:"name"` // 名称
	Path        string `json:"path"` // 路径
	DayzItemsID uint   `json:"dayzItemsID"`
}
