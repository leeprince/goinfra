package consts

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/3/2 下午11:50
 * @Desc:
 */

const (
	ENV_LOCAL   = "local"
	ENV_DEV     = "dev"
	ENV_TEST    = "test"
	ENV_UAT     = "uat"
	ENV_SANDBOX = "sandbox"
	ENV_PROD    = "prod"
)

var AllENV = []string{
	ENV_LOCAL,
	ENV_DEV,
	ENV_TEST,
	ENV_UAT,
	ENV_SANDBOX,
	ENV_PROD,
}
