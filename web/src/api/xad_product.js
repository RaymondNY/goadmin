import service from '@/utils/request'

// @Tags XadProduct
// @Summary 创建XadProduct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.XadProduct true "创建XadProduct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /xad_product/createXadProduct [post]
export const createXadProduct = (data) => {
  return service({
    url: '/xad_product/createXadProduct',
    method: 'post',
    data
  })
}

// @Tags XadProduct
// @Summary 删除XadProduct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.XadProduct true "删除XadProduct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /xad_product/deleteXadProduct [delete]
export const deleteXadProduct = (data) => {
  return service({
    url: '/xad_product/deleteXadProduct',
    method: 'delete',
    data
  })
}

// @Tags XadProduct
// @Summary 删除XadProduct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除XadProduct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /xad_product/deleteXadProduct [delete]
export const deleteXadProductByIds = (data) => {
  return service({
    url: '/xad_product/deleteXadProductByIds',
    method: 'delete',
    data
  })
}

// @Tags XadProduct
// @Summary 更新XadProduct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.XadProduct true "更新XadProduct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /xad_product/updateXadProduct [put]
export const updateXadProduct = (data) => {
  return service({
    url: '/xad_product/updateXadProduct',
    method: 'put',
    data
  })
}

// @Tags XadProduct
// @Summary 用id查询XadProduct
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.XadProduct true "用id查询XadProduct"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /xad_product/findXadProduct [get]
export const findXadProduct = (params) => {
  return service({
    url: '/xad_product/findXadProduct',
    method: 'get',
    params
  })
}

// @Tags XadProduct
// @Summary 分页获取XadProduct列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取XadProduct列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /xad_product/getXadProductList [get]
export const getXadProductList = (params) => {
  return service({
    url: '/xad_product/getXadProductList',
    method: 'get',
    params
  })
}
