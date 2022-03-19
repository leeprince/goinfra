package env

import "github.com/leeprince/goinfra/constants"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/2/25 下午9:21
 * @Desc:   环境检测
 */

func IsProd(env string) (is bool) {
    is = false
    if env == constants.EnvProd {
        is = true
    }
    return
}

func IsProdOrUat(env string) (is bool) {
    is = false
    if env == constants.EnvProd || env == constants.EnvUat  {
        is = true
    }
    return
}

func IsLocal(env string) (is bool) {
    is = false
    if env == constants.EnvLocal {
        is = true
    }
    return
}