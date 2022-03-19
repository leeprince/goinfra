package goredis

import (
    "context"
    "fmt"
    "github.com/go-redis/redis/v8"
    goinfraRedis "github.com/leeprince/goinfra/session/redis"
    "testing"
    "time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/2/27 下午3:58
 * @Desc:
 */

func TestNewTryLock(t *testing.T) {
    // --- redis 客户端
    ctx := context.Background()
    redisConns, err := goinfraRedis.InitRedisClient(ctx, RedisConfs)
    if err != nil {
        fmt.Printf("[goinfraRedis.InitRedisClient] err:%v \n", err)
        return
    }
    redisClient, _ := redisConns[RedisName]
    // ---
    
    type args struct {
        ctx   context.Context
        redis *redis.Client
        opts  []TryLockOption
    }
    
    
    tests := []struct {
        name        string
        args        args
        wantTryLock TryLock
        wantErr     bool
    }{
        {
            args:args{
                ctx:   ctx,
                redis: redisClient,
                opts:  []TryLockOption{
                },
            },
        },
        {
            args:args{
                ctx:   ctx,
                redis: redisClient,
                opts:  []TryLockOption{
                    WithTickerTime(time.Millisecond * 500),
                    WithTimeOut(time.Second * 2),
                },
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            gotTryLock, err := NewTryLock(tt.args.ctx, tt.args.redis, tt.args.opts...)
            if err != nil {
                t.Errorf("[NewTryLock()] error = %v", tt.wantErr)
                return
            }
            
            lock := gotTryLock.Lock(LockKey, LockValue, LockExpire)
            fmt.Printf("[NewTryLock() gotTryLock.Lock] lock:%v \n", lock)
            
            err = gotTryLock.UnLock(LockKey)
            // err = gotTryLock.UnLock(LockKey) // 测试未解锁情况下，获取锁
            fmt.Printf("[NewTryLock() gotTryLock.UnLock] unLock err:%v \n", err)
        })
    }
}