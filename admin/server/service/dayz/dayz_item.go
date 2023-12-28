package dayz

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dayz"
)

type DayzItemService struct {
}

//@function: CreateDayzItem
//@description: 创建物品
//@param: e model.DayzItem
//@return: err error

func (service *DayzItemService) CreateDayzItem(e dayz.DayzItem) (err error) {
	err = global.GVA_DB.Create(&e).Error
	return err
}

//@function: DeleteDayzItem
//@description: 删除物品
//@param: e model.DayzItem
//@return: err error

func (service *DayzItemService) DeleteDayzItem(e dayz.DayzItem) (err error) {
	err = global.GVA_DB.Delete(&e).Error
	return err
}

//@function: UpdateDayzItem
//@description: 更新物品
//@param: e *model.DayzItem
//@return: err error

func (service *DayzItemService) UpdateDayzItem(e *dayz.DayzItem) (err error) {
	err = global.GVA_DB.Save(e).Error
	return err
}

//@function: GetDayzItemByID
//@description: 根据ID获取物品信息
//@param: id uint
//@return: err error, customer model.ExaCustomer

func (service *DayzItemService) GetDayzItemByID(id uint) (err error, dayzItem dayz.DayzItem) {
	err = global.GVA_DB.Preload("DayzItemImg").Where("id = ?", id).First(&dayzItem).Error
	return
}

//@function: GetDayzItem
//@description: 获取物品信息
//@param: dayzItem *dayz.DayzItem
//@return: err error, DayzItemList []dayz.DayzItem

func (service *DayzItemService) GetDayzItem(dayzItem *dayz.DayzItem) (err error, DayzItemList []dayz.DayzItem) {
	db := global.GVA_DB.Model(dayzItem)
	err = db.Preload("DayzItemImg").Find(&DayzItemList).Error
	return
}

//@function: GetDayzItemInfoList
//@description: 分页获取物品列表
//@param: sysUserAuthorityID string, info request.PageInfo
//@return: err error, list interface{}, total int64

func (service *DayzItemService) GetDayzItemInfoList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&dayz.DayzItem{})
	var DayzItemList []dayz.DayzItem
	err = db.Count(&total).Error
	if err != nil {
		return err, DayzItemList, total
	} else {
		err = db.Limit(limit).Offset(offset).Preload("DayzItemImg").Find(&DayzItemList).Error
	}
	return err, DayzItemList, total
}
