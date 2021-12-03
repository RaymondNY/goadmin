package main

import (
	"bytes"
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"time"
)

func main() {
	client, err := CreateSftp("192.168.1.117", "zqx", "root", 22)
	if err != nil {
		return
	}
	//// 3.查看服务器文件**
	//fi, err := client.Lstat("/data/translate/key/yantai/id_rsa")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("%#v\n", fi)
	// 4.上传文件**
	localFilePath := "C:/Program Files/MD5 Ltd/VFC5/vfc5.log"
	remoteFilePath := "/data/translate/key"
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

	target_url := "http://192.168.1.117:9995/uploadLicense"
	filename := "C:/Program Files/MD5 Ltd/VFC5/vfc5.log"
	postFile(filename, target_url)

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
