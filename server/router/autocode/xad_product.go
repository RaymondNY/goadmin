package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type XadProductRouter struct {
}

// InitXadProductRouter 初始化 XadProduct 路由信息
func (s *XadProductRouter) InitXadProductRouter(Router *gin.RouterGroup) {
	xad_productRouter := Router.Group("xad_product").Use(middleware.OperationRecord())
	xad_productRouterWithoutRecord := Router.Group("xad_product")
	var xad_productApi = v1.ApiGroupApp.AutoCodeApiGroup.XadProductApi
	{
		xad_productRouter.POST("createXadProduct", xad_productApi.CreateXadProduct)             // 新建XadProduct
		xad_productRouter.DELETE("deleteXadProduct", xad_productApi.DeleteXadProduct)           // 删除XadProduct
		xad_productRouter.DELETE("deleteXadProductByIds", xad_productApi.DeleteXadProductByIds) // 批量删除XadProduct
		xad_productRouter.PUT("updateXadProduct", xad_productApi.UpdateXadProduct)              // 更新XadProduct
	}
	{
		xad_productRouterWithoutRecord.GET("findXadProduct", xad_productApi.FindXadProduct)       // 根据ID获取XadProduct
		xad_productRouterWithoutRecord.GET("getXadProductList", xad_productApi.GetXadProductList) // 获取XadProduct列表
	}
}
