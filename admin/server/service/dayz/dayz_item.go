package dayz

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/dayz"
	dayzItemReq "github.com/flipped-aurora/gin-vue-admin/server/model/dayz/request"
)

type DayzItemService struct {
}

// CreateDayzItem 创建DayzItem记录
func (dayzItemService *DayzItemService) CreateDayzItem(item dayz.DayzItem) (err error) {
	err = global.GVA_DB.Create(&item).Error
	return err
}

// DeleteDayzItemByIds 批量删除DayzItem记录
func (dayzItemService *DayzItemService) DeleteDayzItemByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]dayz.DayzItem{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateDayzItem 更新DayzItem记录
func (dayzItemService *DayzItemService) UpdateDayzItem(dayzItem dayz.DayzItem) (err error) {
	err = global.GVA_DB.Save(&dayzItem).Error
	return err
}

// GetDayzItemByID 根据id获取DayzItem记录
func (dayzItemService *DayzItemService) GetDayzItemByID(id uint) (err error, item dayz.DayzItem) {
	err = global.GVA_DB.Where("id = ?", id).First(&item).Error
	return
}

// GetDayzItemQuery 查询物品
func (dayzItemService *DayzItemService) GetDayzItemQuery(item dayz.DayzItem) (err error, items []dayz.DayzItem) {
	// 创建db
	db := global.GVA_DB.Model(&dayz.DayzItem{})
	db = db.Preload("Imgs")
	// 如果有条件搜索 下方会自动创建搜索语句
	if item.Name != "" {
		db = db.Where("name = ?", item.Name)
	}

	err = db.Find(&items).Error
	return err, items
}

// GetPageInfosInfoList 分页获取DayzItem记录
func (dayzItemService *DayzItemService) GetDayzItemList(item dayzItemReq.DayzItemReq) (err error, dayzItemss []dayz.DayzItem, total int64) {
	limit := item.PageSize
	offset := item.PageSize * (item.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&dayz.DayzItem{})
	db = db.Preload("Imgs")
	// 如果有条件搜索 下方会自动创建搜索语句
	if item.Name != "" {
		db = db.Where("name = ?", item.Name)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&dayzItemss).Error
	return err, dayzItemss, total
}
