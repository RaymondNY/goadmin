import service from '@/utils/request'

// @Tags XadCustomer
// @Summary 创建XadCustomer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.XadCustomer true "创建XadCustomer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /xadCustomer/createXadCustomer [post]
export const createXadCustomer = (data) => {
  return service({
    url: '/xadCustomer/createXadCustomer',
    method: 'post',
    data
  })
}

// @Tags XadCustomer
// @Summary 删除XadCustomer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.XadCustomer true "删除XadCustomer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /xadCustomer/deleteXadCustomer [delete]
export const deleteXadCustomer = (data) => {
  return service({
    url: '/xadCustomer/deleteXadCustomer',
    method: 'delete',
    data
  })
}

// @Tags XadCustomer
// @Summary 删除XadCustomer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除XadCustomer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /xadCustomer/deleteXadCustomer [delete]
export const deleteXadCustomerByIds = (data) => {
  return service({
    url: '/xadCustomer/deleteXadCustomerByIds',
    method: 'delete',
    data
  })
}

// @Tags XadCustomer
// @Summary 更新XadCustomer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.XadCustomer true "更新XadCustomer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /xadCustomer/updateXadCustomer [put]
export const updateXadCustomer = (data) => {
  return service({
    url: '/xadCustomer/updateXadCustomer',
    method: 'put',
    data
  })
}

// @Tags XadCustomer
// @Summary 用id查询XadCustomer
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.XadCustomer true "用id查询XadCustomer"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /xadCustomer/findXadCustomer [get]
export const findXadCustomer = (params) => {
  return service({
    url: '/xadCustomer/findXadCustomer',
    method: 'get',
    params
  })
}

// @Tags XadCustomer
// @Summary 分页获取XadCustomer列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取XadCustomer列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /xadCustomer/getXadCustomerList [get]
export const getXadCustomerList = (params) => {
  return service({
    url: '/xadCustomer/getXadCustomerList',
    method: 'get',
    params
  })
}
