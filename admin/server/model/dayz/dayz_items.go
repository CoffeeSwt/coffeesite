package dayz

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type DayzItems struct {
	global.GVA_MODEL
	Name           string `json:"name"`        // 名称
	MaxQuantity    uint   `json:"maxQuantity"` //最大数量
	DayzItemsImgs  []DayzItemsImgs
	DayzItemsGroup []DayzItemsGroup
}
