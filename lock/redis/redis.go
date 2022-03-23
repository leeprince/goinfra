package redis

import (
    "context"
    "github.com/leeprince/goinfra/session/redis"
    "time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/3/23 上午2:41
 * @Desc:
 */

type RedisClient struct {
    ctx context.Context
    redis.RedisClient
}

func NewRedisClient(ctx context.Context, redis redis.RedisClient) *RedisClient {
    return &RedisClient{
        ctx:         ctx,
        RedisClient: redis,
    }
}

func (r *RedisClient) Lock(key string, value interface{}, expirtime time.Duration) bool {
    return r.SetNX(r.ctx, key, value, expirtime)
}

func (r *RedisClient) UnLock(key string, value interface{}) bool {
    return r.GetAndDel(r.ctx, key, value)
}
