package yamlutil

import (
	"gopkg.in/yaml.v3"
	"os"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/4/17 18:30
 * @Desc:
 */

// ParseYaml 注意：
// 	- file:文件路径相对于项目所在的"工作目录working directory"
// 		- yaml文件字段对应golang结构体中不指定yaml的标签使用的字段时默认全小写！
//  - config：需要传递地址指针的值过来
func ParseYaml(file string, config interface{}) {
	content, err := os.ReadFile(file)
	if err != nil {
		panic("加载配置文件错误" + file + "错误原因" + err.Error())
	}
	
	err = yaml.Unmarshal(content, config)
	if err != nil {
		panic("解析配置文件错误" + file + "错误原因" + err.Error())
	}
	return
}
