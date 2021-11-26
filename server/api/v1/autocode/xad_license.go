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

type XadLicenseApi struct {
}

var xadLicenseService = service.ServiceGroupApp.AutoCodeServiceGroup.XadLicenseService

// CreateXadLicense 创建XadLicense
// @Tags XadLicense
// @Summary 创建XadLicense
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.XadLicense true "创建XadLicense"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /xadLicense/createXadLicense [post]
func (xadLicenseApi *XadLicenseApi) CreateXadLicense(c *gin.Context) {
	var xadLicense autocode.XadLicense
	_ = c.ShouldBindJSON(&xadLicense)
	if err := xadLicenseService.CreateXadLicense(xadLicense); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteXadLicense 删除XadLicense
// @Tags XadLicense
// @Summary 删除XadLicense
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.XadLicense true "删除XadLicense"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /xadLicense/deleteXadLicense [delete]
func (xadLicenseApi *XadLicenseApi) DeleteXadLicense(c *gin.Context) {
	var xadLicense autocode.XadLicense
	_ = c.ShouldBindJSON(&xadLicense)
	if err := xadLicenseService.DeleteXadLicense(xadLicense); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteXadLicenseByIds 批量删除XadLicense
// @Tags XadLicense
// @Summary 批量删除XadLicense
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除XadLicense"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /xadLicense/deleteXadLicenseByIds [delete]
func (xadLicenseApi *XadLicenseApi) DeleteXadLicenseByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := xadLicenseService.DeleteXadLicenseByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateXadLicense 更新XadLicense
// @Tags XadLicense
// @Summary 更新XadLicense
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.XadLicense true "更新XadLicense"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /xadLicense/updateXadLicense [put]
func (xadLicenseApi *XadLicenseApi) UpdateXadLicense(c *gin.Context) {
	var xadLicense autocode.XadLicense
	_ = c.ShouldBindJSON(&xadLicense)
	if err := xadLicenseService.UpdateXadLicense(xadLicense); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindXadLicense 用id查询XadLicense
// @Tags XadLicense
// @Summary 用id查询XadLicense
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.XadLicense true "用id查询XadLicense"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /xadLicense/findXadLicense [get]
func (xadLicenseApi *XadLicenseApi) FindXadLicense(c *gin.Context) {
	var xadLicense autocode.XadLicense
	_ = c.ShouldBindQuery(&xadLicense)
	if err, rexadLicense := xadLicenseService.GetXadLicense(xadLicense.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rexadLicense": rexadLicense}, c)
	}
}

// GetXadLicenseList 分页获取XadLicense列表
// @Tags XadLicense
// @Summary 分页获取XadLicense列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.XadLicenseSearch true "分页获取XadLicense列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /xadLicense/getXadLicenseList [get]
func (xadLicenseApi *XadLicenseApi) GetXadLicenseList(c *gin.Context) {
	var pageInfo autocodeReq.XadLicenseSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := xadLicenseService.GetXadLicenseInfoList(pageInfo); err != nil {
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
func (xadLicenseApi *XadLicenseApi) GetMachineCode(c *gin.Context) {
	response.OkWithMessage("oooooooo", c)
}
