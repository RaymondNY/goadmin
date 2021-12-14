package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type XadTemplateService struct {
}

// CreateXadTemplate 创建XadTemplate记录
// Author [piexlmax](https://github.com/piexlmax)
func (xad_templateService *XadTemplateService) CreateXadTemplate(xad_template autocode.XadTemplate) (err error) {
	err = global.GVA_DB.Create(&xad_template).Error
	return err
}

// DeleteXadTemplate 删除XadTemplate记录
// Author [piexlmax](https://github.com/piexlmax)
func (xad_templateService *XadTemplateService) DeleteXadTemplate(xad_template autocode.XadTemplate) (err error) {
	err = global.GVA_DB.Delete(&xad_template).Error
	return err
}

// DeleteXadTemplateByIds 批量删除XadTemplate记录
// Author [piexlmax](https://github.com/piexlmax)
func (xad_templateService *XadTemplateService) DeleteXadTemplateByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.XadTemplate{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateXadTemplate 更新XadTemplate记录
// Author [piexlmax](https://github.com/piexlmax)
func (xad_templateService *XadTemplateService) UpdateXadTemplate(xad_template autocode.XadTemplate) (err error) {
	err = global.GVA_DB.Save(&xad_template).Error
	return err
}

// GetXadTemplate 根据id获取XadTemplate记录
// Author [piexlmax](https://github.com/piexlmax)
func (xad_templateService *XadTemplateService) GetXadTemplate(id uint) (err error, xad_template autocode.XadTemplate) {
	err = global.GVA_DB.Where("id = ?", id).First(&xad_template).Error
	return
}

// GetXadTemplateInfoList 分页获取XadTemplate记录+
// Author [piexlmax](https://github.com/piexlmax)
func (xad_templateService *XadTemplateService) GetXadTemplateInfoList(info autoCodeReq.XadTemplateSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.XadTemplate{})
	var xad_templates []autocode.XadTemplate
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Code != "" {
		db = db.Where("code LIKE ?", "%"+info.Code+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&xad_templates).Error
	return err, xad_templates, total
}
