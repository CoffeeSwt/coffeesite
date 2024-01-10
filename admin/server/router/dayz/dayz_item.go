package dayz

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type DayzItemRouter struct{}

func (e *DayzItemRouter) InitDayzItemRouter(Router *gin.RouterGroup) {
	dayzItemRouter := Router.Group("dayzItem")
	dayzItemApi := v1.ApiGroupApp.DayzApiGroup.DayzItemApi
	{
		dayzItemRouter.POST("createDayzItem", dayzItemApi.CreateDayzItem)
		dayzItemRouter.PUT("updateDayzItem", dayzItemApi.UpdateDayzItem)
		dayzItemRouter.DELETE("deleteDayzItemByIds", dayzItemApi.DeleteDayzItemByIds)
	}
	{
		dayzItemRouter.POST("getDayzItem", dayzItemApi.GetDayzItem)
		dayzItemRouter.GET("getDayzItemByID", dayzItemApi.GetDayzItemByID)
		dayzItemRouter.POST("getDayzItemList", dayzItemApi.GetDayzItemList)
	}
}
