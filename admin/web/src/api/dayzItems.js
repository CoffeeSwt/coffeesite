import service from '@/utils/request'

// @Tags DayzItems
// @Summary 创建DayzItems
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DayzItems true "创建DayzItems"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dayzItems/createDayzItems [post]
export const createDayzItems = (data) => {
  return service({
    url: '/dayzItems/createDayzItems',
    method: 'post',
    data
  })
}

// @Tags DayzItems
// @Summary 删除DayzItems
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DayzItems true "删除DayzItems"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dayzItems/deleteDayzItems [delete]
export const deleteDayzItems = (data) => {
  return service({
    url: '/dayzItems/deleteDayzItems',
    method: 'delete',
    data
  })
}

// @Tags DayzItems
// @Summary 删除DayzItems
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DayzItems"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dayzItems/deleteDayzItems [delete]
export const deleteDayzItemsByIds = (data) => {
  return service({
    url: '/dayzItems/deleteDayzItemsByIds',
    method: 'delete',
    data
  })
}

// @Tags DayzItems
// @Summary 更新DayzItems
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DayzItems true "更新DayzItems"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dayzItems/updateDayzItems [put]
export const updateDayzItems = (data) => {
  return service({
    url: '/dayzItems/updateDayzItems',
    method: 'put',
    data
  })
}

// @Tags DayzItems
// @Summary 用id查询DayzItems
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.DayzItems true "用id查询DayzItems"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dayzItems/findDayzItems [get]
export const findDayzItems = (params) => {
  return service({
    url: '/dayzItems/findDayzItems',
    method: 'get',
    params
  })
}

// @Tags DayzItems
// @Summary 分页获取DayzItems列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取DayzItems列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dayzItems/getDayzItemsList [get]
export const getDayzItemsList = (params) => {
  return service({
    url: '/dayzItems/getDayzItemsList',
    method: 'get',
    params
  })
}
