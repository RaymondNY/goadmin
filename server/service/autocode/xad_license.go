package autocode

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/chanyipiaomiao/hltool"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/xad"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type XadLicenseService struct {
}

// CreateXadLicense 创建XadLicense记录
// Author [piexlmax](https://github.com/piexlmax)
func (xadLicenseService *XadLicenseService) CreateXadLicense(xadLicense *autocode.XadLicense, ip string) (err error) {
	err = global.GVA_DB.Create(xadLicense).Error
	if err == nil {
		//1.创建时候生成密钥，只在这时候生成。
		gong, si := GenerateRSA(string(strconv.Itoa(int(xadLicense.ID))))
		//2.获取公钥私钥地址，更新数据库
		xadLicense.PrivateKey = si
		xadLicense.PublicKey = gong
		global.GVA_DB.Save(xadLicense)
		//3.发送私钥
		sendfile(ip, "zqx", "root", si)
		//4.获取模板信息,把机器码更新到指定模板
		tid := xadLicense.TemplateId
		var xadTemplate autocode.XadTemplate
		xadTemplate.ID = uint(*tid)
		global.GVA_DB.Where("id = ?", tid).First(&xadTemplate)
		xadTemplate.Code = xadLicense.MachineCode
		global.GVA_DB.Save(&xadTemplate)

		//5.license数据
		l := License{}
		l.SystemParam.UserNum = xadTemplate.UserNum
		l.SystemParam.Validtime = int64(xadTemplate.Validtime)
		l.SystemParam.ConcurrentUsers = xadTemplate.ConcurrentUsers
		l.SystemParam.Model = xadTemplate.Model
		l.SystemParam.Watermark = xadTemplate.Watermark
		l.SystemParam.UserInfo = xadTemplate.UserInfo
		l.SystemParam.Code = xadTemplate.Code
		l.SystemParam.Product = xadTemplate.Product
		l.SystemParam.MinVersion = xadTemplate.MinVersion
		l.UserParam.ConcurrencyNum = xadTemplate.ConcurrencyNum
		l.UserParam.FunctionModule = xadTemplate.FunctionModule
		l.UserParam.OnceTask = xadTemplate.OnceTask
		l.UserParam.UserConcurrencyPer = xadTemplate.UserConcurrencyPer
		l.UserParam.ConcurrencyModel = xadTemplate.ConcurrencyModel

		data, _ := json.Marshal(l)
		//6.license生成
		p := GenerateLicense(data, strconv.Itoa(int(xadLicense.ID)), gong)
		//7.
		xadLicense.LicenseUrl = p
		global.GVA_DB.Save(xadLicense)
		//e:=global.GVA_DB.Model(&xadLicense).Update("license_url",p).Error
		//if e!=nil {
		//	fmt.Println(e.Error())
		//}
	}
	return err
}

// DeleteXadLicense 删除XadLicense记录
// Author [piexlmax](https://github.com/piexlmax)
func (xadLicenseService *XadLicenseService) DeleteXadLicense(xadLicense autocode.XadLicense) (err error) {
	err = global.GVA_DB.Delete(&xadLicense).Error
	return err
}

// DeleteXadLicenseByIds 批量删除XadLicense记录
// Author [piexlmax](https://github.com/piexlmax)
func (xadLicenseService *XadLicenseService) DeleteXadLicenseByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.XadLicense{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateXadLicense 更新XadLicense记录
// Author [piexlmax](https://github.com/piexlmax)
func (xadLicenseService *XadLicenseService) UpdateXadLicense(xadLicense autocode.XadLicense) (err error) {
	gong := xadLicense.PublicKey
	temId := xadLicense.TemplateId

	var xadTemplate autocode.XadTemplate
	xadTemplate.ID = uint(*temId)
	global.GVA_DB.Where("id = ?", temId).First(&xadTemplate)

	//5.license数据
	l := License{}
	l.SystemParam.UserNum = xadTemplate.UserNum
	l.SystemParam.Validtime = int64(xadTemplate.Validtime)
	l.SystemParam.ConcurrentUsers = xadTemplate.ConcurrentUsers
	l.SystemParam.Model = xadTemplate.Model
	l.SystemParam.Watermark = xadTemplate.Watermark
	l.SystemParam.UserInfo = xadTemplate.UserInfo
	l.SystemParam.Code = xadTemplate.Code
	l.SystemParam.Product = xadTemplate.Product
	l.SystemParam.MinVersion = xadTemplate.MinVersion
	l.UserParam.ConcurrencyNum = xadTemplate.ConcurrencyNum
	l.UserParam.FunctionModule = xadTemplate.FunctionModule
	l.UserParam.OnceTask = xadTemplate.OnceTask
	l.UserParam.UserConcurrencyPer = xadTemplate.UserConcurrencyPer
	l.UserParam.ConcurrencyModel = xadTemplate.ConcurrencyModel
	data, _ := json.Marshal(l)

	//6.license生成
	p := GenerateLicense(data, strconv.Itoa(int(xadLicense.ID)), gong)
	xadLicense.LicenseUrl = p
	err = global.GVA_DB.Save(&xadLicense).Error
	return err
}

// GetXadLicense 根据id获取XadLicense记录
// Author [piexlmax](https://github.com/piexlmax)
func (xadLicenseService *XadLicenseService) GetXadLicense(id uint) (err error, xadLicense autocode.XadLicense) {
	err = global.GVA_DB.Where("id = ?", id).First(&xadLicense).Error
	return
}

// GetXadLicenseInfoList 分页获取XadLicense记录
// Author [piexlmax](https://github.com/piexlmax)
func (xadLicenseService *XadLicenseService) GetXadLicenseInfoList(info autoCodeReq.XadLicenseSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&autocode.XadLicense{})
	var xadLicenses []autocode.XadLicense
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.CustomerId != nil {
		db = db.Where("customer_id = ?", info.CustomerId)
	}
	if info.CreateUser != "" {
		db = db.Where("create_user LIKE ?", "%"+info.CreateUser+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&xadLicenses).Error
	return err, xadLicenses, total
}
func GenerateLicense(data []byte, licenseId string, gong string) (lpath string) {
	licensePath := "E:/data/translate/license"
	a := hltool.IsExist(strings.Join([]string{licensePath, licenseId}, "/"))
	if !a {
		os.MkdirAll(strings.Join([]string{licensePath, licenseId}, "/"), 0777)
	}

	loc, _ := time.LoadLocation("Asia/Shanghai") //设置时区
	t := time.Unix(time.Now().In(loc).Unix(), 0).Format("20060102150405")
	tem := strings.Join([]string{licensePath, licenseId, "license" + t}, "/")
	fl, el := os.Create(tem)
	if el != nil {
		fmt.Println(el.Error())
	}
	defer fl.Close()
	//公钥
	f, err := ioutil.ReadFile(gong)
	if err != nil {
		fmt.Println("read fail", err)
	}
	fming, err := base64.StdEncoding.DecodeString(string(f))
	miwen, err := xad.RsaEncryptBlock(data, fming)
	if err != nil {
		log.Println(err.Error())
	}
	jiamihou := base64.StdEncoding.EncodeToString(miwen)
	fmt.Println("加密后base64：\t", jiamihou)
	//写入license
	fl.WriteString(jiamihou)
	return tem
}

//生成密钥
func GenerateRSA(licenseId string) (gong string, si string) {
	keyPath := "E:/data/translate/key"
	a := hltool.IsExist(strings.Join([]string{keyPath, licenseId}, "/"))
	if !a {
		os.MkdirAll(strings.Join([]string{keyPath, licenseId}, "/"), 0777)
	}
	rsapub := strings.Join([]string{keyPath, licenseId, "id_rsa.pub"}, "/")
	rsa := strings.Join([]string{keyPath, licenseId, "id_rsa"}, "/")

	fkpub, efk1 := os.Create(rsapub)
	if efk1 != nil {
		fmt.Println(efk1.Error())
	}
	defer fkpub.Close()

	fk, efk2 := os.Create(rsa)
	if efk2 != nil {
		fmt.Println(efk2.Error())
	}
	defer fk.Close()

	//RSA的内容使用base64打印
	privateKey, publicKey, _ := xad.GenRSAKey(2048)
	//base64后的密钥
	log.Println("rsa私钥:\t", base64.StdEncoding.EncodeToString(privateKey))
	log.Println("rsa公钥:\t", base64.StdEncoding.EncodeToString(publicKey))

	//写入文件中
	fkpub.WriteString(base64.StdEncoding.EncodeToString(publicKey))
	fk.WriteString(base64.StdEncoding.EncodeToString(privateKey))
	gong = rsapub
	si = rsa
	return
}

func sendfile(ip string, account string, pwd string, localFilePath string) {
	client, err := CreateSftp(ip, account, pwd, 22)
	if err != nil {
		return
	}
	//  1) 打开本地文件
	srcFile, err := os.Open(localFilePath)
	if err != nil {
		panic(err)
	}
	defer srcFile.Close()
	//  绝对路径中获取文件名
	remoteFileName := path.Base(localFilePath)
	fmt.Println(remoteFileName)
	//  2) 打开服务器文件
	remoteFilePath := "/data/translate/key"
	dstFile, err := client.Create(path.Join(remoteFilePath, remoteFileName))
	if err != nil {
		panic(err)
	}
	defer dstFile.Close()
	//  3) 将本地文件内容写入服务器文件
	buf := make([]byte, 1024)
	for {
		n, _ := srcFile.Read(buf) // 将文件2进制内容写入buf字节切片中
		if n == 0 {
			break
		}
		dstFile.Write(buf)
	}
	fmt.Println("文件上传完毕")

	//target_url := "http://192.168.1.117:9995/uploadLicense"
	//filename := "C:/Program Files/MD5 Ltd/VFC5/vfc5.log"
	//postFile(filename, target_url)
}

// CreateSftp 创建sftp会话
func CreateSftp(sshHost, sshUser, sshPassword string, sshPort int) (*sftp.Client, error) {
	// 连接Linux服务器
	conn, err := pwdAuthConnect(sshHost, sshUser, sshPassword, sshPort)
	if err != nil {
		fmt.Println("连接Linux服务器失败", err)
		panic(err)
	}
	//defer conn.Close()

	// 创建sftp会话
	client, err := sftp.NewClient(conn)
	if err != nil {
		return nil, err
	}
	//defer client.Close()
	return client, nil
}

func pwdAuthConnect(sshHost, sshUser, sshPassword string, sshPort int) (*ssh.Client, error) {
	config := ssh.ClientConfig{
		Timeout:         5 * time.Second,
		User:            sshUser,
		Auth:            []ssh.AuthMethod{ssh.Password((sshPassword))},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	addr := fmt.Sprintf("%s:%d", sshHost, sshPort)
	Client, err := ssh.Dial("tcp", addr, &config)
	if err != nil {
		return nil, err
	}
	return Client, err
}

func postFile(filename string, targetUrl string) error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	// 关键的一步操作
	fileWriter, err := bodyWriter.CreateFormFile("file", filename)
	if err != nil {
		fmt.Println("error writing to buffer")
		return err
	}

	// 打开文件句柄操作
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		return err
	}
	defer fh.Close()

	// iocopy
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
}

type License struct {
	SystemParam systemParam
	UserParam   userParam
}

type systemParam struct {
	UserNum         int    `json:"userNum"`         //用户数
	Validtime       int64  `json:"validtime"`       //有效时间（试用、正式）
	ConcurrentUsers int    `json:"concurrentUsers"` //同时在线用户数
	Model           int    `json:"model"`           //模式（1测试、2试用、3正式）
	Watermark       bool   `json:"watermark"`       //水印（开关）
	UserInfo        string `json:"userInfo"`        //用户信息
	Code            string `json:"code"`            //机器码
	Product         int    `json:"product"`         //软件类别  1.翻译机
	MinVersion      string `json:"inVersion"`       //最小启用版本（当小于最小启用版本则不生效）
	SystemParam1    string `json:"systemparam1"`
	SystemParam2    string `json:"systemparam2"`
	SystemParam3    string `json:"systemparam3"`
	SystemParam4    string `json:"systemparam4"`
	SystemParam5    string `json:"systemparam5"`
	SystemParam6    string `json:"systemparam6"`
	SystemParam7    string `json:"systemparam7"`
	SystemParam8    string `json:"systemparam8"`
	SystemParam9    string `json:"systemparam9"`
	SystemParam10   string `json:"systemparam10"`
}

type userParam struct {
	ConcurrencyNum     int    `json:"concurrencyNum"`     //任务并发数
	FunctionModule     string `json:"functionModule"`     //功能模块管理
	OnceTask           int    `json:"onceTask"`           //单次最大任务数
	UserConcurrencyPer int    `json:"userConcurrencyPer"` //每用户最大并发数（0：默认值，无限制；>0 ：限制数量）
	ConcurrencyModel   int    `json:"concurrencyModel"`   //并发模式（0：抢先式并发；1：平均并发）
	UserParam1         string `json:"userparam1"`
	UserParam2         string `json:"userparam2"`
	UserParam3         string `json:"userparam3"`
	UserParam4         string `json:"userparam4"`
	UserParam5         string `json:"userparam5"`
	UserParam6         string `json:"userparam6"`
	UserParam7         string `json:"userparam7"`
	UserParam8         string `json:"userparam8"`
	UserParam9         string `json:"userparam9"`
	UserParam10        string `json:"userparam10"`
}
