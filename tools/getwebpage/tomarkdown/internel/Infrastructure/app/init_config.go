package app

import (
	"flag"
	"fmt"
	"getwebpage-tomarkdown/internel/Infrastructure/config"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024/3/16 01:58
 * @Desc:
 */

// InitConfig 初始化配置
// 	configPathList：只取第一个元素，并且配置的文件路径相对于项目所在的"工作目录working directory"
func InitConfig(configPathList ...string) {
	fmt.Println("InitConfig")
	
	var configPath string
	
	if len(configPathList) > 0 {
		configPath = configPathList[0]
	} else {
		flag.StringVar(&configPath, "conf", "/Users/leeprince/www/go/goinfra/tools/getwebpage/tomarkdown/configs/config.yaml", "config file")
		flag.Parse()
	}
	
	// 解析配置文件
	config.C = &config.Config{}
	fmt.Println("configPath:", configPath)
	content, err := os.ReadFile(configPath)
	if err != nil {
		log.Println("加载配置文件错误" + configPath + "错误原因" + err.Error())
		return
	}
	
	err = yaml.Unmarshal(content, config.C)
	if err != nil {
		fmt.Println("解析配置文件错误" + configPath + "错误原因" + err.Error())
	}
	
	fmt.Println("InitConfig config.C:%+v\n", config.C)
}
