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

type XadProductApi struct {
}

var xad_productService = service.ServiceGroupApp.AutoCodeServiceGroup.XadProductService

// CreateXadProduct 创建XadProduct
// @Tags XadProduct
// @Summary 创建XadProduct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.XadProduct true "创建XadProduct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /xad_product/createXadProduct [post]
func (xad_productApi *XadProductApi) CreateXadProduct(c *gin.Context) {
	var xad_product autocode.XadProduct
	_ = c.ShouldBindJSON(&xad_product)
	if err := xad_productService.CreateXadProduct(xad_product); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteXadProduct 删除XadProduct
// @Tags XadProduct
// @Summary 删除XadProduct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.XadProduct true "删除XadProduct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /xad_product/deleteXadProduct [delete]
func (xad_productApi *XadProductApi) DeleteXadProduct(c *gin.Context) {
	var xad_product autocode.XadProduct
	_ = c.ShouldBindJSON(&xad_product)
	if err := xad_productService.DeleteXadProduct(xad_product); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteXadProductByIds 批量删除XadProduct
// @Tags XadProduct
// @Summary 批量删除XadProduct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除XadProduct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /xad_product/deleteXadProductByIds [delete]
func (xad_productApi *XadProductApi) DeleteXadProductByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := xad_productService.DeleteXadProductByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateXadProduct 更新XadProduct
// @Tags XadProduct
// @Summary 更新XadProduct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.XadProduct true "更新XadProduct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /xad_product/updateXadProduct [put]
func (xad_productApi *XadProductApi) UpdateXadProduct(c *gin.Context) {
	var xad_product autocode.XadProduct
	_ = c.ShouldBindJSON(&xad_product)
	if err := xad_productService.UpdateXadProduct(xad_product); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindXadProduct 用id查询XadProduct
// @Tags XadProduct
// @Summary 用id查询XadProduct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.XadProduct true "用id查询XadProduct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /xad_product/findXadProduct [get]
func (xad_productApi *XadProductApi) FindXadProduct(c *gin.Context) {
	var xad_product autocode.XadProduct
	_ = c.ShouldBindQuery(&xad_product)
	if err, rexad_product := xad_productService.GetXadProduct(xad_product.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rexad_product": rexad_product}, c)
	}
}

// GetXadProductList 分页获取XadProduct列表
// @Tags XadProduct
// @Summary 分页获取XadProduct列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.XadProductSearch true "分页获取XadProduct列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /xad_product/getXadProductList [get]
func (xad_productApi *XadProductApi) GetXadProductList(c *gin.Context) {
	var pageInfo autocodeReq.XadProductSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := xad_productService.GetXadProductInfoList(pageInfo); err != nil {
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
