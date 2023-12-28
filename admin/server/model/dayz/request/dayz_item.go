package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dayz"
)

type DayzItemReq struct {
	request.PageInfo
	dayz.DayzItem
}
