package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type XadLicenseService struct {
}

// CreateXadLicense 创建XadLicense记录
// Author [piexlmax](https://github.com/piexlmax)
func (xadLicenseService *XadLicenseService) CreateXadLicense(xadLicense autocode.XadLicense) (err error) {
	err = global.GVA_DB.Create(&xadLicense).Error
	return err
}

// DeleteXadLicense 删除XadLicense记录
// Author [piexlmax](https://github.com/piexlmax)
func (xadLicenseService *XadLicenseService) DeleteXadLicense(xadLicense autocode.XadLicense) (err error) {
	err = global.GVA_DB.Delete(&xadLicense).Error
	return err
}

// DeleteXadLicenseByIds 批量删除XadLicense记录
// Author [piexlmax](https://github.com/piexlmax)
func (xadLicenseService *XadLicenseService) DeleteXadLicenseByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.XadLicense{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateXadLicense 更新XadLicense记录
// Author [piexlmax](https://github.com/piexlmax)
func (xadLicenseService *XadLicenseService) UpdateXadLicense(xadLicense autocode.XadLicense) (err error) {
	err = global.GVA_DB.Save(&xadLicense).Error
	return err
}

// GetXadLicense 根据id获取XadLicense记录
// Author [piexlmax](https://github.com/piexlmax)
func (xadLicenseService *XadLicenseService) GetXadLicense(id uint) (err error, xadLicense autocode.XadLicense) {
	err = global.GVA_DB.Where("id = ?", id).First(&xadLicense).Error
	return
}

// GetXadLicenseInfoList 分页获取XadLicense记录
// Author [piexlmax](https://github.com/piexlmax)
func (xadLicenseService *XadLicenseService) GetXadLicenseInfoList(info autoCodeReq.XadLicenseSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.XadLicense{})
	var xadLicenses []autocode.XadLicense
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.CustomerId != nil {
		db = db.Where("customer_id = ?", info.CustomerId)
	}
	if info.CreateUser != "" {
		db = db.Where("create_user LIKE ?", "%"+info.CreateUser+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&xadLicenses).Error
	return err, xadLicenses, total
}
