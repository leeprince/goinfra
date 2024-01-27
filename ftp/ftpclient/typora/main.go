package main

import (
	"flag"
	"fmt"
	"github.com/jlaffaye/ftp"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"strings"
	"sync"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024/1/27 02:17
 * @Desc:
 */

func main() {
	// 初始化配置
	InitConfig()
	
	// os.Args[0] 是程序的名字，所以我们从 os.Args[1] 开始获取参数
	paths := os.Args[1:]
	
	if len(paths) == 0 {
		fmt.Println("请输入本地绝对路径文件")
		return
	}
	
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

func uploadImage(path string) {
	// 连接FTP服务器
	client, err := ftp.Dial(fmt.Sprintf("%s:%s", CFG.Host, CFG.Port))
	log.Println("uploadImage Dial")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Quit()
	
	// 登录
	err = client.Login(CFG.Username, CFG.Password)
	log.Println("uploadImage Login")
	if err != nil {
		log.Fatal(err)
	}
	
	// 打开文件
	file, err := os.Open(path)
	log.Println("uploadImage Open")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	// 上传文件
	remoteFileDir := "/typora/images/"
	remoteFilePath := remoteFileDir + getFileNameFromPath(path)
	log.Println("uploadImage remoteFileDir:", remoteFileDir, "remoteFilePath:", remoteFilePath)
	
	// 切换到目录下
	err = client.ChangeDir(remoteFileDir)
	if err != nil {
		if !strings.Contains(err.Error(), "No such file or directory") {
			log.Println("切换目录错误：", err.Error())
		}
		// 尝试创建目录
		log.Println("目录不存在，尝试重新创建目录")
		// 确定 ftp 是否会所需目录，ftp 未创建可以通过该方法创建，已创建的需要检查报错
		err = client.MakeDir(remoteFileDir)
		if err != nil {
			if !strings.Contains(err.Error(), "File exists") {
				log.Fatal("无法创建目标目录：", err)
			}
		}
		// 重新切换到目录下
		err = client.ChangeDir(remoteFileDir)
		if err != nil {
			log.Fatal("创建目录后，重新切换到目录错误：", err.Error())
			return
		}
		
		log.Println("创建目录后，尝试重新创建目录成功")
	}
	
	log.Println("uploadImage Stor ...")
	err = client.Stor(remoteFilePath, file)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("uploadImage Stor success")
}

func getFileNameFromPath(path string) string {
	// 这是一个简单的函数，用于从路径中获取文件名
	// 在实际使用中，您可能需要使用复杂的方法来处理路径中的特殊字符
	return path[strings.LastIndex(path, "/")+1:]
}

type FTP struct {
	Host     string `yaml:"Host"`
	Port     string `yaml:"Port"`
	Username string `yaml:"Username"`
	Password string `yaml:"Password"`
}

var CFG *FTP

// InitConfig 初始化配置
// 	configPathList：只取第一个元素，并且配置的文件路径相对于项目所在的"工作目录working directory"
func InitConfig(configPathList ...string) {
	log.Println("InitConfig")
	
	var configPath string
	
	if len(configPathList) > 0 {
		configPath = configPathList[0]
	} else {
		flag.StringVar(&configPath, "conf", "./config.yaml", "config file")
		flag.Parse()
	}
	
	// 解析配置文件
	CFG = &FTP{}
	fmt.Println("configPath:", configPath)
	content, err := os.ReadFile(configPath)
	if err != nil {
		log.Println("加载配置文件错误" + configPath + "错误原因" + err.Error())
		return
	}
	
	err = yaml.Unmarshal(content, &CFG)
	if err != nil {
		log.Println("解析配置文件错误" + configPath + "错误原因" + err.Error())
	}
	
	fmt.Printf("InitConfig CFG:%+v\n", CFG)
}
