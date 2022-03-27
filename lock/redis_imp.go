package lock

import (
    "context"
    "github.com/leeprince/goinfra/storage/redis"
    "time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/3/23 上午2:41
 * @Desc:
 */


// RedisClient 实现 LockClientInterface 接口
var _ LockClientInterface = (*RedisClient)(nil)

type RedisClient struct {
    redis.RedisClient
}

func NewRedisClient(redis redis.RedisClient) *RedisClient {
    return &RedisClient{
        RedisClient: redis,
    }
}

func (r *RedisClient) Lock(ctx context.Context, key string, value interface{}, expirtime time.Duration) (bool, error) {
    return r.WithContext(ctx).SetNx(key, value, expirtime)
}

func (r *RedisClient) UnLock(ctx context.Context, key string, value interface{}) error {
    return r.WithContext(ctx).GetAndDel(key, value)
}
