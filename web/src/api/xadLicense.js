import service from '@/utils/request'

// @Tags XadLicense
// @Summary 创建XadLicense
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.XadLicense true "创建XadLicense"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /xadLicense/createXadLicense [post]
export const createXadLicense = (data) => {
  return service({
    url: '/xadLicense/createXadLicense',
    method: 'post',
    data,
  })
}

// @Tags XadLicense
// @Summary 删除XadLicense
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.XadLicense true "删除XadLicense"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /xadLicense/deleteXadLicense [delete]
export const deleteXadLicense = (data) => {
  return service({
    url: '/xadLicense/deleteXadLicense',
    method: 'delete',
    data,
  })
}

// @Tags XadLicense
// @Summary 删除XadLicense
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除XadLicense"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /xadLicense/deleteXadLicense [delete]
export const deleteXadLicenseByIds = (data) => {
  return service({
    url: '/xadLicense/deleteXadLicenseByIds',
    method: 'delete',
    data,
  })
}

// @Tags XadLicense
// @Summary 更新XadLicense
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.XadLicense true "更新XadLicense"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /xadLicense/updateXadLicense [put]
export const updateXadLicense = (data) => {
  return service({
    url: '/xadLicense/updateXadLicense',
    method: 'put',
    data,
  })
}

// @Tags XadLicense
// @Summary 用id查询XadLicense
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.XadLicense true "用id查询XadLicense"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /xadLicense/findXadLicense [get]
export const findXadLicense = (params) => {
  return service({
    url: '/xadLicense/findXadLicense',
    method: 'get',
    params,
  })
}

// @Tags XadLicense
// @Summary 分页获取XadLicense列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取XadLicense列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /xadLicense/getXadLicenseList [get]
export const getXadLicenseList = (params) => {
  return service({
    url: '/xadLicense/getXadLicenseList',
    method: 'get',
    params,
  })
}

export const getMachineCode = (params) => {
  return service({
    url: '/xadLicense/getMachineCode',
    method: 'get',
    params,
  })
}
