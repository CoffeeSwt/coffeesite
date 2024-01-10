package dayz

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dayz"
	dayzItemReq "github.com/flipped-aurora/gin-vue-admin/server/model/dayz/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DayzItemApi struct {
}

// @Tags CreateDayzItem
// @Summary 创建物品
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dayz.DayzItem
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /dayz/createDayzItem [post]
func (dayzItemApi *DayzItemApi) CreateDayzItem(c *gin.Context) {
	var dayzItem dayz.DayzItem
	_ = c.ShouldBindJSON(&dayzItem)
	if err := dayzItemService.CreateDayzItem(dayzItem); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)

	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteDayzItemByIds 批量删除DayzItem
// @Summary 批量删除DayzItem
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DayzItem"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /dayz/deleteDayzItemByIds [delete]
func (dayzItemApi *DayzItemApi) DeleteDayzItemByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := dayzItemService.DeleteDayzItemByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateDayzItem 更新DayzItem
// @Tags DayzItem
// @Summary 更新客户信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body dayz.dayzItem
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dayz/updateDayzItem [put]
func (dayzItemApi *DayzItemApi) UpdateDayzItem(c *gin.Context) {
	var dayzItem dayz.DayzItem
	_ = c.ShouldBindJSON(&dayzItem)
	if err := utils.Verify(dayzItem.GVA_MODEL, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := dayzItemService.UpdateDayzItem(dayzItem); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// GetDayzItemByID 根据ID获取物品
// @Tags DayzItem
// @Summary 根据ID获取物品
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dayz.DayzItem
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dayz/getDayzItemByID [get]
func (dayzItemApi *DayzItemApi) GetDayzItemByID(c *gin.Context) {
	var dayzItem dayz.DayzItem
	_ = c.ShouldBindQuery(&dayzItem)
	if err := utils.Verify(dayzItem.GVA_MODEL, utils.IdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err, data := dayzItemService.GetDayzItemByID(dayzItem.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(gin.H{"items": data}, c)
	}
}

// GetDayzItem
// @Tags DayzItem
// @Summary 根据ID获取物品
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dayz.DayzItem
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dayz/getDayzItemByID [get]
func (dayzItemApi *DayzItemApi) GetDayzItem(c *gin.Context) {
	var dayzItem dayzItemReq.DayzItemQuery
	_ = c.ShouldBindJSON(&dayzItem)
	err, data := dayzItemService.GetDayzItemQuery(dayzItem)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(gin.H{"items": data}, c)
	}
}

// @Tags GetDayzItemList
// @Summary 分页获取物品
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query dayzItemReq.DayzItemReq
// @Router /dayz/getDayzItemList [get]
func (dayzItemApi *DayzItemApi) GetDayzItemList(c *gin.Context) {
	var pageInfo dayzItemReq.DayzItemPageQuery
	_ = c.ShouldBindJSON(&pageInfo)
	if err, list, total := dayzItemService.GetDayzItemList(pageInfo); err != nil {
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
