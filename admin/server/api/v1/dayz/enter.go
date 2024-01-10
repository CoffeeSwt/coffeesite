package dayz

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	DayzItemApi
	DayzItemImgApi
}

var (
	dayzItemService    = service.ServiceGroupApp.DayzServiceGroup.DayzItemService
	dayzItemImgService = service.ServiceGroupApp.DayzServiceGroup.DayzItemImgsService
)
