package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type XadLicenseRouter struct {
}

// InitXadLicenseRouter 初始化 XadLicense 路由信息
func (s *XadLicenseRouter) InitXadLicenseRouter(Router *gin.RouterGroup) {
	xadLicenseRouter := Router.Group("xadLicense").Use(middleware.OperationRecord())
	xadLicenseRouterWithoutRecord := Router.Group("xadLicense")
	var xadLicenseApi = v1.ApiGroupApp.AutoCodeApiGroup.XadLicenseApi
	{
		xadLicenseRouter.POST("createXadLicense", xadLicenseApi.CreateXadLicense)             // 新建XadLicense
		xadLicenseRouter.DELETE("deleteXadLicense", xadLicenseApi.DeleteXadLicense)           // 删除XadLicense
		xadLicenseRouter.DELETE("deleteXadLicenseByIds", xadLicenseApi.DeleteXadLicenseByIds) // 批量删除XadLicense
		xadLicenseRouter.PUT("updateXadLicense", xadLicenseApi.UpdateXadLicense)              // 更新XadLicense
	}
	{
		xadLicenseRouterWithoutRecord.GET("findXadLicense", xadLicenseApi.FindXadLicense)       // 根据ID获取XadLicense
		xadLicenseRouterWithoutRecord.GET("getXadLicenseList", xadLicenseApi.GetXadLicenseList) // 获取XadLicense列表
	}
}
