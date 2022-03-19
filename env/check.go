package env

import (
    "github.com/leeprince/goinfra/consts"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/2/25 下午9:21
 * @Desc:   环境检测
 */

func IsProd(env string) (is bool) {
    is = false
    if env == consts.EnvProd {
        is = true
    }
    return
}

func IsProdOrUat(env string) (is bool) {
    is = false
    if env == consts.EnvProd || env == consts.EnvUat  {
        is = true
    }
    return
}

func IsLocal(env string) (is bool) {
    is = false
    if env == consts.EnvLocal {
        is = true
    }
    return
}