package dayz

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	DayzItemApi
}

var (
	dayzItemService = service.ServiceGroupApp.DayzServiceGroup.DayzItemService
)
