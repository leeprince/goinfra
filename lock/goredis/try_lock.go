package goredis

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
    timeOut    time.Duration // 获取锁总的超时时间. timeOut <= tickerTime 时只尝试获取一次
    debug      bool          // 是否打印获取锁过程的记录
}

type TryLockOption struct {
    f func(*TryLock)
}

const (
    DefaultTickerTime            = time.Millisecond * 500
    DefaultTimeOut               = time.Second * 2 // 相当于尝试4次+1次获取锁
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
    // if tryLock.timeOut < tryLock.tickerTime {
    //     tryLock.timeOut = tryLock.tickerTime + DefaultTimeOutThanTickerTime
    // }
    return
}

// 获取分布式锁
func (l *TryLock) Lock(key string, value interface{}, expirtime time.Duration) bool {
    return l.redis.SetNX(l.ctx, key, value, expirtime).Val()
}

// 释放分布式锁
func (l *TryLock) UnLock(key string, value interface{}) error {
    script := `
    if redis.call("GET", KEYS[1]) ~= ARGV[1] then
        return nil
    end
    return redis.call("DEL", KEYS[1])
    `
    l.redis.Get().Scan()
    
    val, err := l.redis.Eval(l.ctx, script, []string{key}, value).Result()
    if val == nil || err != nil {
        return fmt.Errorf("[unLock] Fail.val:%v;err:%v", val, err)
    }
    return err
}
