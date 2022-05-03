package callerfile0101

import (
    "github.com/leeprince/goinfra/plog"
    "runtime"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/4/15 下午2:05
 * @Desc:
 */

func Caller(i int)  {
    pc, file, line, ok := runtime.Caller(i)
    plog.Infof("Caller::::::::i=%d:%s-%s-%d-%t", i, runtime.FuncForPC(pc).Name(), file, line, ok)
}