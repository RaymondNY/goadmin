package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type XadCustomerRouter struct {
}

// InitXadCustomerRouter 初始化 XadCustomer 路由信息
func (s *XadCustomerRouter) InitXadCustomerRouter(Router *gin.RouterGroup) {
	xadCustomerRouter := Router.Group("xadCustomer").Use(middleware.OperationRecord())
	xadCustomerRouterWithoutRecord := Router.Group("xadCustomer")
	var xadCustomerApi = v1.ApiGroupApp.AutoCodeApiGroup.XadCustomerApi
	{
		xadCustomerRouter.POST("createXadCustomer", xadCustomerApi.CreateXadCustomer)             // 新建XadCustomer
		xadCustomerRouter.DELETE("deleteXadCustomer", xadCustomerApi.DeleteXadCustomer)           // 删除XadCustomer
		xadCustomerRouter.DELETE("deleteXadCustomerByIds", xadCustomerApi.DeleteXadCustomerByIds) // 批量删除XadCustomer
		xadCustomerRouter.PUT("updateXadCustomer", xadCustomerApi.UpdateXadCustomer)              // 更新XadCustomer
	}
	{
		xadCustomerRouterWithoutRecord.GET("findXadCustomer", xadCustomerApi.FindXadCustomer)       // 根据ID获取XadCustomer
		xadCustomerRouterWithoutRecord.GET("getXadCustomerList", xadCustomerApi.GetXadCustomerList) // 获取XadCustomer列表
	}
}
