package ftpclient

import (
	"bytes"
	"fmt"
	"github.com/jlaffaye/ftp"
	"path/filepath"
	"strings"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024/3/16 03:19
 * @Desc:
 */

func (r *FtpClient) UploadImage(imgDataBytes []byte, remoteFileDir, remoteFilePath string) (remoteAccessUrl string, err error) {
	if err = r.checkInit(); err != nil {
		fmt.Println("UploadImage checkInit err:", err)
		return
	}
	
	// 连接FTP服务器
	client, err := ftp.Dial(fmt.Sprintf("%s:%s", r.Conf.Host, r.Conf.Port))
	fmt.Println("UploadImage Dial")
	if err != nil {
		fmt.Println("Dial err:", err)
		return
	}
	defer client.Quit()
	
	// 登录
	err = client.Login(r.Conf.Username, r.Conf.Password)
	if err != nil {
		fmt.Println("UploadImage Login", err)
		return
	}
	
	// 拆分路径为各级父目录
	dirs := strings.Split(remoteFileDir, "/")
	// 递归创建缺失的父级目录
	currentDir := "/"
	for _, dir := range dirs {
		if dir == "" {
			continue
		}
		currentDir = filepath.Join(currentDir, dir)
		fmt.Println("currentDir:", currentDir)
		// 检查目录是否存在，不存在则创建
		err = client.ChangeDir(currentDir)
		if err != nil {
			if !strings.Contains(err.Error(), "No such file or directory") {
				fmt.Println("切换目录错误：", err.Error())
				return
			}
			// 尝试创建目录
			fmt.Println("目录不存在，尝试重新创建目录")
			// 确定 ftp 是否会所需目录，ftp 未创建可以通过该方法创建，已创建的需要检查报错
			err = client.MakeDir(currentDir)
			if err != nil {
				if !strings.Contains(err.Error(), "File exists") {
					fmt.Println("无法创建目标目录：", err)
					return
				}
			}
			// 重新切换到目录下
			err = client.ChangeDir(currentDir)
			if err != nil {
				fmt.Println("创建目录后，重新切换到目录错误：", err.Error())
				return
			}
			
			fmt.Println("创建目录后，尝试重新创建目录成功")
		}
	}
	
	fmt.Println("UploadImage Stor ...")
	imgDataReader := bytes.NewReader(imgDataBytes)
	// Code=553;Msg=Can't open that file: No such file or directory 可能的原因：项目根目录remoteFilePath不对
	if remoteFilePath[0] != filepath.Separator {
		remoteFilePath = filepath.Join(string(filepath.Separator), remoteFilePath)
	}
	err = client.Stor(remoteFilePath, imgDataReader)
	if err != nil {
		fmt.Println("UploadImage Stor err:", err.Error())
		return
	}
	fmt.Println("UploadImage Stor success")
	
	// 访问地址
	remoteAccessUrl = r.AccessHost + remoteFilePath
	fmt.Println("remoteAccessUrl:", remoteAccessUrl)
	
	return
}
