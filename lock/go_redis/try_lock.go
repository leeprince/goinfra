package go_redis

import (
    "context"
    "fmt"
    "github.com/go-redis/redis/v8"
    "time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/2/25 下午9:27
 * @Desc:
 */

type TryLock struct {
    ctx        context.Context
    redis      *redis.Client
    tickerTime time.Duration // 重新获取锁的间隔时间
    timeOut    time.Duration // 本次获取锁的超时时间. timeOut >= tickerTime
}

type TryLockOption struct {
    f func(*TryLock)
}

const (
    DefaultTickerTime            = time.Millisecond * 500
    DefaultTimeOut               = time.Second * 2
    DefaultTimeOutThanTickerTime = time.Millisecond * 2 // 默认的超时时间大于定时时间。只有 timeOut < tickerTime 时生效
)

const (
    DefaultLockExpireTime = time.Second * 2 // 默认的锁过期时间
)

func NewTryLock(ctx context.Context, redis *redis.Client, opts ...TryLockOption) (tryLock *TryLock, err error) {
    tryLock = &TryLock{
        ctx:        ctx,
        redis:      redis,
        tickerTime: DefaultTickerTime,
        timeOut:    DefaultTimeOut,
    }
    for _, opt := range opts {
        opt.f(tryLock)
    }
    if tryLock.timeOut < tryLock.tickerTime {
        tryLock.timeOut = tryLock.tickerTime + DefaultTimeOutThanTickerTime
    }
    fmt.Println("[NewTryLock] tryLock:", tryLock)
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

// 获取分布式锁
func (l *TryLock) Lock(key string, value interface{}, expirtime time.Duration) bool {
    if expirtime == 0 {
        expirtime = DefaultLockExpireTime
    }
    setLock := l.redis.SetNX(l.ctx, key, value, expirtime).Val()
    if setLock {
        fmt.Println("[TryLock@Lock] setLock Suucessfuly")
        return true
    }
    
    ticker := time.NewTicker(l.tickerTime)
    timeOutAfter := time.After(l.timeOut)
    for {
        select {
        case <-timeOutAfter:
            fmt.Println("[TryLock@Lock] setLock Fail <-time.After(l.timeOut):", l.timeOut)
            return false
        case <-ticker.C:
            setLock := l.redis.SetNX(l.ctx, key, value, expirtime).Val()
            if !setLock {
                fmt.Println("[TryLock@Lock] setLock continue <-ticker.C")
                continue
            }
            fmt.Println("[TryLock@Lock] setLock Suucessfuly <-ticker.C:")
            return true
        }
    }
}

// 释放分布式锁
func (l *TryLock) UnLock(key string) error {
    return l.redis.Del(l.ctx, key).Err()
}
