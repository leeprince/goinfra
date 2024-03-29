package redislock

import (
	"context"
	"fmt"
	goinfraRedis "github.com/leeprince/goinfra/storage/redis"
	"testing"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/3/24 上午1:37
 * @Desc:
 */

var (
	RedisName  = "local"
	RedisConfs = goinfraRedis.RedisConfs{
		RedisName: goinfraRedis.RedisConf{
			Network:  "tcp",
			Addr:     "127.0.0.1:6379",
			Username: "",
			Password: "",
			DB:       0,
			PoolSize: 2,
		},
	}
	LockKey    = "princeLockKey01"
	LockValue  = "princeLockValue01"
	LockExpire = time.Second * 3
)

func TestNewTryLock(t *testing.T) {
	// --- redis 客户端

	// Goredis 客户端
	err := goinfraRedis.InitGoredisList(RedisConfs)
	if err != nil {
		fmt.Printf("[goinfraRedis.InitGoredis] err:%v \n", err)
		return
	}
	redisClient := goinfraRedis.GetGoredis(RedisName)

	// Redigo 客户端
	/*err := goinfraRedis.InitRedigo(RedisConfs)
	  if err != nil {
	      fmt.Printf("[goinfraRedis.InitGoredis] err:%v \n", err)
	      return
	  }
	  redisClient := goinfraRedis.GetRedigo(RedisName)*/

	// --- redis 客户端-end

	ctx := context.Background()

	type args struct {
		ctx    context.Context
		client LockClientInterface
		opts   []TryLockOption
	}
	tests := []struct {
		name        string
		args        args
		wantTryLock *TryLock
		wantErr     bool
	}{
		{
			args: args{
				ctx:    ctx,
				client: NewRedisClient(redisClient),
				opts: []TryLockOption{
					WithDebug(true),
				},
			},
		},
		{
			args: args{
				ctx:    ctx,
				client: NewRedisClient(redisClient),
				opts: []TryLockOption{
					WithTickerTime(time.Millisecond * 500),
					WithTimeOut(time.Millisecond * 500),
					WithDebug(true),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTryLock, err := NewTryLock(tt.args.ctx, tt.args.client, tt.args.opts...)
			if err != nil {
				t.Errorf("[NewTryLock()] error = %v", err)
				return
			}

			lock := gotTryLock.Lock(LockKey, LockValue, LockExpire)
			fmt.Printf("[NewTryLock() gotTryLock.TryLock] lock:%v \n", lock)

			// ok := gotTryLock.UnLock(LockKey+"-", LockValue)
			// ok := gotTryLock.UnLock(LockKey, LockValue) // 测试未解锁情况下，获取锁
			// fmt.Printf("[NewTryLock() gotTryLock.UnLock] unLock bool:%v \n", ok)
		})
	}
}
