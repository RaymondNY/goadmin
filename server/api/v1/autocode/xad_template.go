package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autocodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type XadTemplateApi struct {
}

var xad_templateService = service.ServiceGroupApp.AutoCodeServiceGroup.XadTemplateService

// CreateXadTemplate 创建XadTemplate
// @Tags XadTemplate
// @Summary 创建XadTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.XadTemplate true "创建XadTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /xad_template/createXadTemplate [post]
func (xad_templateApi *XadTemplateApi) CreateXadTemplate(c *gin.Context) {
	var xad_template autocode.XadTemplate
	_ = c.ShouldBindJSON(&xad_template)
	if err := xad_templateService.CreateXadTemplate(xad_template); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteXadTemplate 删除XadTemplate
// @Tags XadTemplate
// @Summary 删除XadTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.XadTemplate true "删除XadTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /xad_template/deleteXadTemplate [delete]
func (xad_templateApi *XadTemplateApi) DeleteXadTemplate(c *gin.Context) {
	var xad_template autocode.XadTemplate
	_ = c.ShouldBindJSON(&xad_template)
	if err := xad_templateService.DeleteXadTemplate(xad_template); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteXadTemplateByIds 批量删除XadTemplate
// @Tags XadTemplate
// @Summary 批量删除XadTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除XadTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /xad_template/deleteXadTemplateByIds [delete]
func (xad_templateApi *XadTemplateApi) DeleteXadTemplateByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := xad_templateService.DeleteXadTemplateByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateXadTemplate 更新XadTemplate
// @Tags XadTemplate
// @Summary 更新XadTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.XadTemplate true "更新XadTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /xad_template/updateXadTemplate [put]
func (xad_templateApi *XadTemplateApi) UpdateXadTemplate(c *gin.Context) {
	var xad_template autocode.XadTemplate
	_ = c.ShouldBindJSON(&xad_template)
	if err := xad_templateService.UpdateXadTemplate(xad_template); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindXadTemplate 用id查询XadTemplate
// @Tags XadTemplate
// @Summary 用id查询XadTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.XadTemplate true "用id查询XadTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /xad_template/findXadTemplate [get]
func (xad_templateApi *XadTemplateApi) FindXadTemplate(c *gin.Context) {
	var xad_template autocode.XadTemplate
	_ = c.ShouldBindQuery(&xad_template)
	if err, rexad_template := xad_templateService.GetXadTemplate(xad_template.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rexad_template": rexad_template}, c)
	}
}

// GetXadTemplateList 分页获取XadTemplate列表
// @Tags XadTemplate
// @Summary 分页获取XadTemplate列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.XadTemplateSearch true "分页获取XadTemplate列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /xad_template/getXadTemplateList [get]
func (xad_templateApi *XadTemplateApi) GetXadTemplateList(c *gin.Context) {
	var pageInfo autocodeReq.XadTemplateSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := xad_templateService.GetXadTemplateInfoList(pageInfo); err != nil {
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
