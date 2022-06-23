package concurrency

import "time"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/6/12 下午11:03
 * @Desc:   可选项
 */

const (
    defaultTimeout = time.Second * 30
)

type params struct {
    timeout time.Duration // 并发任务
    stopOnError bool // 发生错误则终止
}
type WillTerminateOpt func(param *params)

func WithTimeout(timeout time.Duration) WillTerminateOpt {
    return func(param *params) {
        param.timeout = timeout
    }
}
func WithStopOnError(v bool) WillTerminateOpt {
    return func(param *params) {
        param.stopOnError = v
    }
}