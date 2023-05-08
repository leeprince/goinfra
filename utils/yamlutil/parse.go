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

func ParseYaml(file string, config interface{}) {
	content, err := os.ReadFile(file)
	if err != nil {
		panic("加载配置文件错误" + file + "错误原因" + err.Error())
	}

	err = yaml.Unmarshal(content, config)
	if err != nil {
		panic("解析配置文件错误" + file + "错误原因" + err.Error())
	}
}
