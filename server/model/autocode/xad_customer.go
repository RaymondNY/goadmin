// 自动生成模板XadCustomer
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// XadCustomer 结构体
// 如果含有time.Time 请自行import time包
type XadCustomer struct {
	global.GVA_MODEL
	Name       string `json:"name" form:"name" gorm:"column:name;comment:名称;type:varchar(191);"`
	Address    string `json:"address" form:"address" gorm:"column:address;comment:地址;type:varchar(191);"`
	EmsCode    *int   `json:"emsCode" form:"emsCode" gorm:"column:ems_code;comment:邮编;type:int"`
	CreateUser string `json:"createUser" form:"createUser" gorm:"column:create_user;comment:;type:varchar(191);"`
}

// TableName XadCustomer 表名
func (XadCustomer) TableName() string {
	return "xad_customer"
}
