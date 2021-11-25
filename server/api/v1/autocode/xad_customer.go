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

type XadCustomerApi struct {
}

var xadCustomerService = service.ServiceGroupApp.AutoCodeServiceGroup.XadCustomerService

// CreateXadCustomer 创建XadCustomer
// @Tags XadCustomer
// @Summary 创建XadCustomer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.XadCustomer true "创建XadCustomer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /xadCustomer/createXadCustomer [post]
func (xadCustomerApi *XadCustomerApi) CreateXadCustomer(c *gin.Context) {
	var xadCustomer autocode.XadCustomer
	_ = c.ShouldBindJSON(&xadCustomer)
	if err := xadCustomerService.CreateXadCustomer(xadCustomer); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteXadCustomer 删除XadCustomer
// @Tags XadCustomer
// @Summary 删除XadCustomer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.XadCustomer true "删除XadCustomer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /xadCustomer/deleteXadCustomer [delete]
func (xadCustomerApi *XadCustomerApi) DeleteXadCustomer(c *gin.Context) {
	var xadCustomer autocode.XadCustomer
	_ = c.ShouldBindJSON(&xadCustomer)
	if err := xadCustomerService.DeleteXadCustomer(xadCustomer); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteXadCustomerByIds 批量删除XadCustomer
// @Tags XadCustomer
// @Summary 批量删除XadCustomer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除XadCustomer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /xadCustomer/deleteXadCustomerByIds [delete]
func (xadCustomerApi *XadCustomerApi) DeleteXadCustomerByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := xadCustomerService.DeleteXadCustomerByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateXadCustomer 更新XadCustomer
// @Tags XadCustomer
// @Summary 更新XadCustomer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.XadCustomer true "更新XadCustomer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /xadCustomer/updateXadCustomer [put]
func (xadCustomerApi *XadCustomerApi) UpdateXadCustomer(c *gin.Context) {
	var xadCustomer autocode.XadCustomer
	_ = c.ShouldBindJSON(&xadCustomer)
	if err := xadCustomerService.UpdateXadCustomer(xadCustomer); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindXadCustomer 用id查询XadCustomer
// @Tags XadCustomer
// @Summary 用id查询XadCustomer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.XadCustomer true "用id查询XadCustomer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /xadCustomer/findXadCustomer [get]
func (xadCustomerApi *XadCustomerApi) FindXadCustomer(c *gin.Context) {
	var xadCustomer autocode.XadCustomer
	_ = c.ShouldBindQuery(&xadCustomer)
	if err, rexadCustomer := xadCustomerService.GetXadCustomer(xadCustomer.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rexadCustomer": rexadCustomer}, c)
	}
}

// GetXadCustomerList 分页获取XadCustomer列表
// @Tags XadCustomer
// @Summary 分页获取XadCustomer列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.XadCustomerSearch true "分页获取XadCustomer列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /xadCustomer/getXadCustomerList [get]
func (xadCustomerApi *XadCustomerApi) GetXadCustomerList(c *gin.Context) {
	var pageInfo autocodeReq.XadCustomerSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := xadCustomerService.GetXadCustomerInfoList(pageInfo); err != nil {
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
