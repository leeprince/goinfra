package lock

import (
    "context"
    "fmt"
    "time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/3/21 下午11:11
 * @Desc:
 */

type TryLock struct {
    ctx        context.Context
    client     LockClientInterface //
    tickerTime time.Duration       // 重新获取锁的间隔时间
    timeOut    time.Duration       // 获取锁总的超时时间. timeOut <= tickerTime 时只尝试获取一次
    debug      bool                // 是否打印获取锁过程的记录
}

type LockClientInterface interface {
    Lock(ctx context.Context, key string, value interface{}, expirtime time.Duration) error
    UnLock(ctx context.Context, key string, value interface{}) error
}

type TryLockOption struct {
    f func(*TryLock)
}

func NewTryLock(ctx context.Context, client LockClientInterface, opts ...TryLockOption) (tryLock *TryLock, err error) {
    tryLock = &TryLock{
        ctx:        ctx,
        client:     client,
        tickerTime: DefaultTickerTime,
        timeOut:    DefaultTimeOut,
    }
    for _, opt := range opts {
        opt.f(tryLock)
    }
    if tryLock.timeOut < tryLock.tickerTime {
        tryLock.timeOut = tryLock.tickerTime
    }
    return
}

func WithTickerTime(t time.Duration) TryLockOption {
    return TryLockOption{
        f: func(lock *TryLock) {
            lock.tickerTime = t
        },
    }
}

func WithTimeOut(t time.Duration) TryLockOption {
    return TryLockOption{
        f: func(lock *TryLock) {
            lock.timeOut = t
        },
    }
}

func WithDebug(debug bool) TryLockOption {
    return TryLockOption{
        f: func(lock *TryLock) {
            lock.debug = debug
        },
    }
}

func (l *TryLock) Lock(key string, value interface{}, expirtime time.Duration) bool {
    debug := l.debug
    if expirtime == 0 {
        expirtime = DefaultLockExpireTime
    }
    if err := l.client.Lock(l.ctx, key, value, expirtime); err == nil {
        if debug {
            fmt.Println("[TryLock@Lock] setLock Suucessfuly")
        }
        return true
    }
    
    if l.timeOut <= l.tickerTime {
        return false
    }
    
    ticker := time.NewTicker(l.tickerTime)
    timeoutAfter := time.After(l.timeOut)
    for {
        select {
        case <-timeoutAfter:
            if debug {
                fmt.Println("[TryLock@Lock] Fail <-time.After(l.timeOut):", l.timeOut)
            }
            return false
        case <-ticker.C:
            if err := l.client.Lock(l.ctx, key, value, expirtime); err != nil {
                if debug {
                    fmt.Println("[TryLock@Lock] <-ticker.C continue")
                }
                continue
            }
            if debug {
                fmt.Println("[TryLock@Lock] <-ticker.C Suucessfuly ")
            }
            return true
        }
    }
}

func (l *TryLock) UnLock(key string, value interface{}) bool {
    err := l.client.UnLock(l.ctx, key, value)
    if err != nil {
        return false
    }
    return true
}
