package dayz

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dayz"
	dayzReq "github.com/flipped-aurora/gin-vue-admin/server/model/dayz/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DayzItemsApi struct {
}

var dayzItemsService = service.ServiceGroupApp.DayzItemsService.DayzItemsService

// CreateDayzItems 创建DayzItems
// @Tags DayzItems
// @Summary 创建DayzItems
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.DayzItems true "创建DayzItems"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dayzItems/createDayzItems [post]
func (dayzItemsApi *DayzItemsApi) CreateDayzItems(c *gin.Context) {
	var dayzItems dayz.DayzItems
	_ = c.ShouldBindJSON(&dayzItems)
	if err := dayzItemsService.CreateDayzItems(dayzItems); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteDayzItems 删除DayzItems
// @Tags DayzItems
// @Summary 删除DayzItems
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.DayzItems true "删除DayzItems"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dayzItems/deleteDayzItems [delete]
func (dayzItemsApi *DayzItemsApi) DeleteDayzItems(c *gin.Context) {
	var dayzItems dayz.DayzItems
	_ = c.ShouldBindJSON(&dayzItems)
	if err := dayzItemsService.DeleteDayzItems(dayzItems); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteDayzItemsByIds 批量删除DayzItems
// @Tags DayzItems
// @Summary 批量删除DayzItems
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DayzItems"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /dayzItems/deleteDayzItemsByIds [delete]
func (dayzItemsApi *DayzItemsApi) DeleteDayzItemsByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := dayzItemsService.DeleteDayzItemsByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateDayzItems 更新DayzItems
// @Tags DayzItems
// @Summary 更新DayzItems
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.DayzItems true "更新DayzItems"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dayzItems/updateDayzItems [put]
func (dayzItemsApi *DayzItemsApi) UpdateDayzItems(c *gin.Context) {
	var dayzItems dayz.DayzItems
	_ = c.ShouldBindJSON(&dayzItems)
	if err := dayzItemsService.UpdateDayzItems(dayzItems); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindDayzItems 用id查询DayzItems
// @Tags DayzItems
// @Summary 用id查询DayzItems
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.DayzItems true "用id查询DayzItems"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dayzItems/findDayzItems [get]
func (dayzItemsApi *DayzItemsApi) FindDayzItems(c *gin.Context) {
	var dayzItems dayz.DayzItems
	_ = c.ShouldBindQuery(&dayzItems)
	if err, redayzItems := dayzItemsService.GetDayzItems(dayzItems.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redayzItems": redayzItems}, c)
	}
}

// GetDayzItemsList 分页获取DayzItems列表
// @Tags DayzItems
// @Summary 分页获取DayzItems列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.DayzItemsSearch true "分页获取DayzItems列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dayzItems/getDayzItemsList [get]
func (dayzItemsApi *DayzItemsApi) GetDayzItemsList(c *gin.Context) {
	var pageInfo dayzReq.DayzItemsSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := dayzItemsService.GetDayzItemsInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
