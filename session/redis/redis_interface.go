package redis

import (
    "context"
    "time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/3/23 上午2:15
 * @Desc:
 */

type RedisClient interface {
    SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) bool
    GetAndDel(ctx context.Context, key string, value interface{}) bool
    GetString(ctx context.Context, key string) string
    GetBytes(ctx context.Context, key string) []byte
    GetScan(ctx context.Context, key string, value interface{}) error
}