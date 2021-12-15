import service from '@/utils/request'

// @Tags XadTemplate
// @Summary 创建XadTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.XadTemplate true "创建XadTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /xad_template/createXadTemplate [post]
export const createXadTemplate = (data) => {
  return service({
    url: '/xad_template/createXadTemplate',
    method: 'post',
    data
  })
}

// @Tags XadTemplate
// @Summary 删除XadTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.XadTemplate true "删除XadTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /xad_template/deleteXadTemplate [delete]
export const deleteXadTemplate = (data) => {
  return service({
    url: '/xad_template/deleteXadTemplate',
    method: 'delete',
    data
  })
}

// @Tags XadTemplate
// @Summary 删除XadTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除XadTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /xad_template/deleteXadTemplate [delete]
export const deleteXadTemplateByIds = (data) => {
  return service({
    url: '/xad_template/deleteXadTemplateByIds',
    method: 'delete',
    data
  })
}

// @Tags XadTemplate
// @Summary 更新XadTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.XadTemplate true "更新XadTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /xad_template/updateXadTemplate [put]
export const updateXadTemplate = (data) => {
  return service({
    url: '/xad_template/updateXadTemplate',
    method: 'put',
    data
  })
}

// @Tags XadTemplate
// @Summary 用id查询XadTemplate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.XadTemplate true "用id查询XadTemplate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /xad_template/findXadTemplate [get]
export const findXadTemplate = (params) => {
  return service({
    url: '/xad_template/findXadTemplate',
    method: 'get',
    params
  })
}

// @Tags XadTemplate
// @Summary 分页获取XadTemplate列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取XadTemplate列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /xad_template/getXadTemplateList [get]
export const getXadTemplateList = (params) => {
  return service({
    url: '/xad_template/getXadTemplateList',
    method: 'get',
    params
  })
}
