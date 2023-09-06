package initconfig

import (
	"errors"
	"fmt"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/9/6 18:50
 * @Desc:
 */

func Init(env string) error {
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

	return nil
}
