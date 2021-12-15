// 自动生成模板XadTemplate
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// XadTemplate 结构体
// 如果含有time.Time 请自行import time包
type XadTemplate struct {
	global.GVA_MODEL
	UserNum            int    `json:"userNum" form:"userNum" gorm:"column:user_num;comment:用户数;type:int"`
	Validtime          int    `json:"validtime" form:"validtime" gorm:"column:validtime;comment:有效时间（试用、正式）;type:bigint"`
	ConcurrentUsers    int    `json:"concurrentUsers" form:"concurrentUsers" gorm:"column:concurrent_users;comment:同时在线用户数;type:int"`
	Model              int    `json:"model" form:"model" gorm:"column:model;comment:模式（1测试、2试用、3正式）;type:int"`
	Watermark          bool   `json:"watermark" form:"watermark" gorm:"column:watermark;comment:水印（开关）"`
	UserInfo           string `json:"userInfo" form:"userInfo" gorm:"column:user_info;comment:用户信息;type:varchar(255);"`
	Code               string `json:"code" form:"code" gorm:"column:code;comment:机器码;type:text;"`
	Product            int    `json:"product" form:"product" gorm:"column:product;comment:软件类别  1.翻译机;type:int"`
	MinVersion         string `json:"minVersion" form:"minVersion" gorm:"column:min_version;comment:最小启用版本（当小于最小启用版本则不生效）;type:varchar(100);"`
	ConcurrencyNum     int    `json:"concurrencyNum" form:"concurrencyNum" gorm:"column:concurrency_num;comment:任务并发数;type:int"`
	FunctionModule     string `json:"functionModule" form:"functionModule" gorm:"column:function_module;comment:功能模块管理;type:varchar(100);"`
	OnceTask           int    `json:"onceTask" form:"onceTask" gorm:"column:once_task;comment:单次最大任务数;type:int"`
	UserConcurrencyPer int    `json:"userConcurrencyPer" form:"userConcurrencyPer" gorm:"column:user_concurrency_per;comment:每用户最大并发数（0：默认值，无限制；>0 ：限制数量）;type:int"`
	ConcurrencyModel   int    `json:"concurrencyModel" form:"concurrencyModel" gorm:"column:concurrency_model;comment:并发模式（0：抢先式并发；1：平均并发）;type:int"`
}

// TableName XadTemplate 表名
func (XadTemplate) TableName() string {
	return "xad_template"
}
