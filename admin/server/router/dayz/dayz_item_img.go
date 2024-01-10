package dayz

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type DayzItemImgRouter struct{}

func (e *DayzItemRouter) InitDayzItemImgRouter(Router *gin.RouterGroup) {
	dayzItemImgRouter := Router.Group("dayzItemImg")
	dayzItemApi := v1.ApiGroupApp.DayzApiGroup.DayzItemImgApi
	{
		dayzItemImgRouter.POST("uploadFile", dayzItemApi.UploadFile)
		dayzItemImgRouter.GET("getFileList", dayzItemApi.GetFileList)
		dayzItemImgRouter.DELETE("deleteFile", dayzItemApi.DeleteFile)
	}
}
