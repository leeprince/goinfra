package retry

import (
    "github.com/leeprince/goinfra/consts"
    "time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/3/18 上午9:58
 * @Desc:
 */

type retryTask struct {
    retryNum  int           // 重试次数
    retryTime time.Duration // 重试时间
}

type MessageHandleOption func(ms *retryTask)

func NewRetryTask(opts ...MessageHandleOption) *retryTask {
    mr := &retryTask{}
    
    for _, opt := range opts {
        opt(mr)
    }
    
    if mr.retryNum == 0 {
        mr.retryNum = consts.RetryDefaultNum
    }
    
    if mr.retryTime == 0 {
        mr.retryNum = consts.RetryDefaulTime
    }
    return mr
}

// --- 设置参数 ---
func WithRetryNum(n int) MessageHandleOption {
    return func(ms *retryTask) {
        ms.retryNum = n
    }
}

func WithRetryTime(t time.Duration) MessageHandleOption {
    return func(ms *retryTask) {
        ms.retryTime = t
    }
}

// --- 设置参数 -end ---

// retry 重试
func (mr *retryTask) Retry(retryData []byte, f func([]byte) error) error {
    var err error
    for i := 0; i < mr.retryNum; i++ {
        ticker := time.NewTicker(mr.retryTime)
        <-ticker.C
        err = f(retryData)
        if err == nil {
            return nil
        }
    }
    
    return err
}

// TODO: 退出策略 - prince@todo 2022/3/18 上午10:05