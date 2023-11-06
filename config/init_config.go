package config

import (
	"flag"
	"fmt"
	"github.com/leeprince/goinfra/plog"
	"github.com/leeprince/goinfra/utils/yamlutil"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/10/29 14:36
 * @Desc:
 */

var C *Config

// InitConfig 初始化配置
// 	configPathList：只取第一个元素，并且配置的文件路径相对于项目所在的"工作目录working directory"
func InitConfig(configPathList ...string) {
	plog.Info("InitConfig")
	
	var configPath string
	
	if len(configPathList) > 0 {
		configPath = configPathList[0]
	} else {
		flag.StringVar(&configPath, "conf", "./config.yaml", "config file")
		flag.Parse()
	}
	
	// 解析配置文件
	C = &Config{}
	yamlutil.ParseYaml(configPath, C)
	
	fmt.Printf("InitConfig C:%+v\n", C)
	plog.WithField("configPath", configPath).WithField("C", C).Info("InitConfig")
}
