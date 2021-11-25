// 自动生成模板XadLicense
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// XadLicense 结构体
// 如果含有time.Time 请自行import time包
type XadLicense struct {
	global.GVA_MODEL
	CustomerId  *int   `json:"customerId" form:"customerId" gorm:"column:customer_id;comment:客户编号;type:int"`
	ProductId   *int   `json:"productId" form:"productId" gorm:"column:product_id;comment:产品编号;type:int"`
	CreateUser  string `json:"createUser" form:"createUser" gorm:"column:create_user;comment:;type:varchar(191);"`
	LicenseUrl  string `json:"licenseUrl" form:"licenseUrl" gorm:"column:license_url;comment:license地址;type:varchar(255);"`
	PublicKey   string `json:"publicKey" form:"publicKey" gorm:"column:public_key;comment:公钥地址;type:varchar(255);"`
	PrivateKey  string `json:"privateKey" form:"privateKey" gorm:"column:private_key;comment:私钥地址;type:varchar(255);"`
	MachineCode string `json:"machineCode" form:"machineCode" gorm:"column:machine_code;comment:机器码;type:text;"`
}

// TableName XadLicense 表名
func (XadLicense) TableName() string {
	return "xad_license"
}
