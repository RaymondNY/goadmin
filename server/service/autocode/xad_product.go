package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type XadProductService struct {
}

// CreateXadProduct 创建XadProduct记录
// Author [piexlmax](https://github.com/piexlmax)
func (xad_productService *XadProductService) CreateXadProduct(xad_product autocode.XadProduct) (err error) {
	err = global.GVA_DB.Create(&xad_product).Error
	return err
}

// DeleteXadProduct 删除XadProduct记录
// Author [piexlmax](https://github.com/piexlmax)
func (xad_productService *XadProductService) DeleteXadProduct(xad_product autocode.XadProduct) (err error) {
	err = global.GVA_DB.Delete(&xad_product).Error
	return err
}

// DeleteXadProductByIds 批量删除XadProduct记录
// Author [piexlmax](https://github.com/piexlmax)
func (xad_productService *XadProductService) DeleteXadProductByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.XadProduct{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateXadProduct 更新XadProduct记录
// Author [piexlmax](https://github.com/piexlmax)
func (xad_productService *XadProductService) UpdateXadProduct(xad_product autocode.XadProduct) (err error) {
	err = global.GVA_DB.Save(&xad_product).Error
	return err
}

// GetXadProduct 根据id获取XadProduct记录
// Author [piexlmax](https://github.com/piexlmax)
func (xad_productService *XadProductService) GetXadProduct(id uint) (err error, xad_product autocode.XadProduct) {
	err = global.GVA_DB.Where("id = ?", id).First(&xad_product).Error
	return
}

// GetXadProductInfoList 分页获取XadProduct记录
// Author [piexlmax](https://github.com/piexlmax)
func (xad_productService *XadProductService) GetXadProductInfoList(info autoCodeReq.XadProductSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.XadProduct{})
	var xad_products []autocode.XadProduct
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.CreateUser != "" {
		db = db.Where("create_user = ?", info.CreateUser)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&xad_products).Error
	return err, xad_products, total
}
