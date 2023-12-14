package dayz

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dayz"
	dayzItemsReq "github.com/flipped-aurora/gin-vue-admin/server/model/dayz/request"
)

type DayzItemsService struct {
}

// CreateDayzItems 创建DayzItems记录
func (dayzItemsService *DayzItemsService) CreateDayzItems(dayzItems dayz.DayzItems) (err error) {
	err = global.GVA_DB.Create(&dayzItems).Error
	return err
}

// DeleteDayzItems 删除DayzItems记录
func (dayzItemsService *DayzItemsService) DeleteDayzItems(dayzItems dayz.DayzItems) (err error) {
	err = global.GVA_DB.Delete(&dayzItems).Error
	return err
}

// DeleteDayzItemsByIds 批量删除DayzItems记录
func (dayzItemsService *DayzItemsService) DeleteDayzItemsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]dayz.DayzItems{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateDayzItems 更新DayzItems记录
func (dayzItemsService *DayzItemsService) UpdateDayzItems(dayzItems dayz.DayzItems) (err error) {
	err = global.GVA_DB.Save(&dayzItems).Error
	return err
}

// GetDayzItems 根据id获取DayzItems记录
func (dayzItemsService *DayzItemsService) GetDayzItems(id uint) (err error, dayzItems dayz.DayzItems) {
	err = global.GVA_DB.Where("id = ?", id).First(&dayzItems).Error
	return
}

// GetDayzItemsInfoList 分页获取DayzItems记录
func (dayzItemsService *DayzItemsService) GetDayzItemsInfoList(info dayzItemsReq.DayzItemsSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&dayz.DayzItems{})
	var dayzItemss []dayz.DayzItems

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&dayzItemss).Error
	return err, dayzItemss, total
}
