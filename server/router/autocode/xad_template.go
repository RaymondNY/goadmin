package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type XadTemplateRouter struct {
}

// InitXadTemplateRouter 初始化 XadTemplate 路由信息
func (s *XadTemplateRouter) InitXadTemplateRouter(Router *gin.RouterGroup) {
	xad_templateRouter := Router.Group("xad_template").Use(middleware.OperationRecord())
	xad_templateRouterWithoutRecord := Router.Group("xad_template")
	var xad_templateApi = v1.ApiGroupApp.AutoCodeApiGroup.XadTemplateApi
	{
		xad_templateRouter.POST("createXadTemplate", xad_templateApi.CreateXadTemplate)             // 新建XadTemplate
		xad_templateRouter.DELETE("deleteXadTemplate", xad_templateApi.DeleteXadTemplate)           // 删除XadTemplate
		xad_templateRouter.DELETE("deleteXadTemplateByIds", xad_templateApi.DeleteXadTemplateByIds) // 批量删除XadTemplate
		xad_templateRouter.PUT("updateXadTemplate", xad_templateApi.UpdateXadTemplate)              // 更新XadTemplate
	}
	{
		xad_templateRouterWithoutRecord.GET("findXadTemplate", xad_templateApi.FindXadTemplate)       // 根据ID获取XadTemplate
		xad_templateRouterWithoutRecord.GET("getXadTemplateList", xad_templateApi.GetXadTemplateList) // 获取XadTemplate列表
	}
}
