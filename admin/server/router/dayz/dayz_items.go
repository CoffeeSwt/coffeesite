package dayz

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DayzItemsRouter struct {
}

// InitDayzItemsRouter 初始化 DayzItems 路由信息
func (s *DayzItemsRouter) InitDayzItemsRouter(Router *gin.RouterGroup) {
	dayzItemsRouter := Router.Group("dayzItems").Use(middleware.OperationRecord())
	dayzItemsRouterWithoutRecord := Router.Group("dayzItems")
	var dayzItemsApi = v1.ApiGroupApp.DayzItemsApiGroup.DayzItemsApi
	{
		dayzItemsRouter.POST("createDayzItems", dayzItemsApi.CreateDayzItems)             // 新建DayzItems
		dayzItemsRouter.DELETE("deleteDayzItems", dayzItemsApi.DeleteDayzItems)           // 删除DayzItems
		dayzItemsRouter.DELETE("deleteDayzItemsByIds", dayzItemsApi.DeleteDayzItemsByIds) // 批量删除DayzItems
		dayzItemsRouter.PUT("updateDayzItems", dayzItemsApi.UpdateDayzItems)              // 更新DayzItems
	}
	{
		dayzItemsRouterWithoutRecord.GET("findDayzItems", dayzItemsApi.FindDayzItems)       // 根据ID获取DayzItems
		dayzItemsRouterWithoutRecord.GET("getDayzItemsList", dayzItemsApi.GetDayzItemsList) // 获取DayzItems列表
	}
}
