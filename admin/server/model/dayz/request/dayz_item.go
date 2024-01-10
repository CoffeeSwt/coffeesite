package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type DayzItemPageQuery struct {
	request.PageInfo
	DayzItemQuery
}

type DayzItemQuery struct {
	Name     string `json:"name" form:"name" gorm:"comment:物品名称"`
	Category []int  `json:"category" form:"category" gorm:"comment:物品类别"`
}
