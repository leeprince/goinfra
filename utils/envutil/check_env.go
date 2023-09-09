package envutil

import (
	"github.com/leeprince/goinfra/consts"
	"os"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/2/25 下午9:21
 * @Desc:   环境检测
 */

func IsProd(env string) (is bool) {
	is = false
	if env == consts.ENV_PROD {
		is = true
	}
	return
}

func IsProdOrSandbox(env string) (is bool) {
	return env == consts.ENV_PROD || env == consts.ENV_UAT || env == consts.ENV_SANDBOX
}

func IsLocal(env string) (is bool) {
	is = false
	if env == consts.ENV_LOCAL {
		is = true
	}
	return
}

func EnvIsMock() (is bool) {
	if os.Getenv("IsMock") == "True" ||
		os.Getenv("ismock") == "true" ||
		os.Getenv("IsMock") == "true" ||
		os.Getenv("IsMock") == "True" {
		return true
	}
	return false
}
