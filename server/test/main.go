package main

import (
	"bytes"
	"crypto/md5"
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"time"
)

//func main() {
//	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
//	//autocode.GenerateRSA("1")
//	//dsn := "root:root@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
//	//db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	//db.AutoMigrate(&User{})
//	//user := User{Name: "Jinzhu", Age: 18}
//	//db.Create(&user)
//
//
//}

func postFile(filename string, targetUrl string) error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	//加入token
	curtime := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(curtime, 10))

	//关键的一步操作 ，模拟上传文件
	fileWriter, err := bodyWriter.CreateFormFile("file", filename)
	if err != nil {
		fmt.Println("error writing to buffer")
		return err
	}

	//打开文件句柄操作
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		return err
	}
	defer fh.Close()

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}

	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()

	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	resp_body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
	return nil

	//上面的例子详细展示了客户端如何向服务器上传一个文件的例子，客户端通过multipart.Write把文件的文本流写入一个缓存中，然后调用http的Post方法把缓存传到服务器。

	//如果你还有其他普通字段例如username之类的需要同时写入，那么可以调用multipart的WriteField方法写很多其他类似的字段。
}

// sample usage
func main() {
	target_url := "http://192.168.1.117:9995/uploadLicense"
	filename := "E:\\data\\translate\\license\\25\\license20211214162611"
	postFile(filename, target_url)
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
