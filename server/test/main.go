package main

import (
	"database/sql"
	"github.com/flipped-aurora/gin-vue-admin/server/service/autocode"
	"gorm.io/gorm"
	"time"
)

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情

	autocode.GenerateRSA("1")
	//dsn := "root:root@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	//db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	//db.AutoMigrate(&User{})
	//user := User{Name: "Jinzhu", Age: 18}
	//db.Create(&user)
}

type User struct {
	ID           uint           `gorm:"primarykey"` // 主键ID
	CreatedAt    time.Time      // 创建时间
	UpdatedAt    time.Time      // 更新时间
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
}
