package base

import (
	"errors"
	"fmt"
	"github.com/leeprince/goinfra/utils/yamlutil"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/9/7 16:32
 * @Desc:
 */

var Config *Conf

type Conf struct {
	Env string `yaml:"env"`
}

func InitConfig(env string) error {
	var isEnvVaild bool
	for _, e := range []string{"dev", "test", "prod"} {
		if env == e {
			isEnvVaild = true
			break
		}
	}
	if !isEnvVaild {
		return errors.New("env 不符合配置项")
	}

	confFilePath := fmt.Sprintf("./conf/conf-%s.yaml", env)
	fmt.Println("配置文件路径：", confFilePath)

	Config = &Conf{}
	yamlutil.ParseYaml(confFilePath, Config)
	fmt.Println("InitConfig Config:", Config)

	return nil
}
