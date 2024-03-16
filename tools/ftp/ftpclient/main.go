package main

import (
	"flag"
	"fmt"
	"github.com/jlaffaye/ftp"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024/1/27 02:17
 * @Desc:
 */

const isDebug = true

const (
	getFilePartPathFromPathDelimiter = "typora"
)

func main() {
	// 初始化配置
	InitConfig()
	
	// os.Args[0] 是程序的名字，所以我们从 os.Args[1] 开始获取参数
	paths := os.Args[1:]
	
	if len(paths) == 0 {
		fmt.Println("请输入本地绝对路径文件")
		return
	}
	
	SequenceUpload(paths) // 顺序上传
	// concurrencyUpload(paths) // 并发上传
}

// 顺序上传
func SequenceUpload(paths []string) {
	for _, path := range paths {
		mylog("path:", path)
		accessUrl := uploadImage(path)
		fmt.Println(accessUrl)
	}
}

// 并发上传
func concurrencyUpload(paths []string) {
	var wg sync.WaitGroup
	for _, path := range paths {
		wg.Add(1)
		go func(path string) {
			defer wg.Done()
			uploadImage(path)
		}(path)
	}
	wg.Wait()
}

func uploadImage(path string) (remoteAccessUrl string) {
	// 连接FTP服务器
	client, err := ftp.Dial(fmt.Sprintf("%s:%s", CFG.Host, CFG.Port))
	mylog("uploadImage Dial")
	if err != nil {
		mylog("Dial err:", err)
		return
	}
	defer client.Quit()
	
	// 登录
	err = client.Login(CFG.Username, CFG.Password)
	if err != nil {
		mylog("uploadImage Login", err)
		return
	}
	
	// 打开文件
	file, err := os.Open(path)
	mylog("uploadImage Open")
	if err != nil {
		mylog(err)
		return
	}
	defer file.Close()
	
	// 上传文件
	remoteFileDir, remoteFilePath := getFilePartPathFromPath(path, getFilePartPathFromPathDelimiter)
	mylog("uploadImage remoteFileDir:", remoteFileDir, "remoteFilePath:", remoteFilePath)
	
	// 拆分路径为各级父目录
	dirs := strings.Split(remoteFileDir, "/")
	// 递归创建缺失的父级目录
	currentDir := "/"
	for _, dir := range dirs {
		if dir == "" {
			continue
		}
		currentDir = filepath.Join(currentDir, dir)
		mylog("currentDir:", currentDir)
		// 检查目录是否存在，不存在则创建
		err = client.ChangeDir(currentDir)
		if err != nil {
			if !strings.Contains(err.Error(), "No such file or directory") {
				mylog("切换目录错误：", err.Error())
				return
			}
			// 尝试创建目录
			mylog("目录不存在，尝试重新创建目录")
			// 确定 ftp 是否会所需目录，ftp 未创建可以通过该方法创建，已创建的需要检查报错
			err = client.MakeDir(currentDir)
			if err != nil {
				if !strings.Contains(err.Error(), "File exists") {
					mylog("无法创建目标目录：", err)
					return
				}
			}
			// 重新切换到目录下
			err = client.ChangeDir(currentDir)
			if err != nil {
				mylog("创建目录后，重新切换到目录错误：", err.Error())
				return
			}
			
			mylog("创建目录后，尝试重新创建目录成功")
		}
	}
	
	mylog("uploadImage Stor ...")
	err = client.Stor(remoteFilePath, file)
	if err != nil {
		mylog(err)
		return
	}
	mylog("uploadImage Stor success")
	
	// 访问地址
	remoteAccessUrl = CFG.AccessHost + remoteFilePath
	mylog("remoteAccessUrl:", remoteAccessUrl)
	
	return
}

func getFileNameFromPath(path string) string {
	// 在实际使用中，您可能需要使用复杂的方法来处理路径中的特殊字符
	// 这是一个简单的函数，用于从路径中获取文件名
	return path[strings.LastIndex(path, "/")+1:]
}

// 同意分隔符分割字符串，并去后半部分：delimiter 用于切分为 2 维数组，并获取第二部分的路径去上传
func getFilePartPathFromPath(path string, delimiter string) (remoteFileDir, remoteFilePath string) {
	// 在实际使用中，您可能需要使用复杂的方法来处理路径中的特殊字符
	pathList := strings.Split(path, delimiter)
	remoteFilePath = filepath.Join("/"+delimiter, pathList[len(pathList)-1])
	remoteFileDir = remoteFilePath[:strings.LastIndex(remoteFilePath, "/")]
	return
}

type Config struct {
	Host       string `yaml:"Host"`
	Port       string `yaml:"Port"`
	Username   string `yaml:"Username"`
	Password   string `yaml:"Password"`
	AccessHost string `yaml:"AccessHost"`
}

var CFG *Config

// InitConfig 初始化配置
// 	configPathList：只取第一个元素，并且配置的文件路径相对于项目所在的"工作目录working directory"
func InitConfig(configPathList ...string) {
	mylog("InitConfig")
	
	var configPath string
	
	if len(configPathList) > 0 {
		configPath = configPathList[0]
	} else {
		flag.StringVar(&configPath, "conf", "/Users/leeprince/www/go/goinfra/tools/ftp/ftpclient/config.yaml", "config file")
		flag.Parse()
	}
	
	// 解析配置文件
	CFG = &Config{}
	mylog("configPath:", configPath)
	content, err := os.ReadFile(configPath)
	if err != nil {
		log.Println("加载配置文件错误" + configPath + "错误原因" + err.Error())
		return
	}
	
	err = yaml.Unmarshal(content, CFG)
	if err != nil {
		mylog("解析配置文件错误" + configPath + "错误原因" + err.Error())
	}
	
	mylog("InitConfig CFG:%+v\n", CFG)
}

func mylog(v ...any) {
	if isDebug {
		log.Println(v...)
	}
}
