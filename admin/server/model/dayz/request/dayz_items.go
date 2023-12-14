package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dayz"
)

type DayzItemsSearch struct {
	dayz.DayzItems
	request.PageInfo
}
