package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type XadCustomerService struct {
}

// CreateXadCustomer 创建XadCustomer记录
// Author [piexlmax](https://github.com/piexlmax)
func (xadCustomerService *XadCustomerService) CreateXadCustomer(xadCustomer autocode.XadCustomer) (err error) {
	err = global.GVA_DB.Create(&xadCustomer).Error
	return err
}

// DeleteXadCustomer 删除XadCustomer记录
// Author [piexlmax](https://github.com/piexlmax)
func (xadCustomerService *XadCustomerService) DeleteXadCustomer(xadCustomer autocode.XadCustomer) (err error) {
	err = global.GVA_DB.Delete(&xadCustomer).Error
	return err
}

// DeleteXadCustomerByIds 批量删除XadCustomer记录
// Author [piexlmax](https://github.com/piexlmax)
func (xadCustomerService *XadCustomerService) DeleteXadCustomerByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.XadCustomer{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateXadCustomer 更新XadCustomer记录
// Author [piexlmax](https://github.com/piexlmax)
func (xadCustomerService *XadCustomerService) UpdateXadCustomer(xadCustomer autocode.XadCustomer) (err error) {
	err = global.GVA_DB.Save(&xadCustomer).Error
	return err
}

// GetXadCustomer 根据id获取XadCustomer记录
// Author [piexlmax](https://github.com/piexlmax)
func (xadCustomerService *XadCustomerService) GetXadCustomer(id uint) (err error, xadCustomer autocode.XadCustomer) {
	err = global.GVA_DB.Where("id = ?", id).First(&xadCustomer).Error
	return
}

// GetXadCustomerInfoList 分页获取XadCustomer记录
// Author [piexlmax](https://github.com/piexlmax)
func (xadCustomerService *XadCustomerService) GetXadCustomerInfoList(info autoCodeReq.XadCustomerSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.XadCustomer{})
	var xadCustomers []autocode.XadCustomer
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&xadCustomers).Error
	return err, xadCustomers, total
}
