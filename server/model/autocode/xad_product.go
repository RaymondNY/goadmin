// 自动生成模板XadProduct
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// XadProduct 结构体
// 如果含有time.Time 请自行import time包
type XadProduct struct {
	global.GVA_MODEL
	Name       string `json:"name" form:"name" gorm:"column:name;comment:产品名称;type:varchar(20);"`
	HardInfo   string `json:"hardInfo" form:"hardInfo" gorm:"column:hard_info;comment:硬件信息;type:text;"`
	Profile    string `json:"profile" form:"profile" gorm:"column:profile;comment:简介;type:varchar(255);"`
	Status     *int   `json:"status" form:"status" gorm:"column:status;comment:状态;type:int"`
	CreateUser string `json:"createUser" form:"createUser" gorm:"column:create_user;comment:创建人;type:varchar(20);"`
}

// TableName XadProduct 表名
func (XadProduct) TableName() string {
	return "xad_product"
}
