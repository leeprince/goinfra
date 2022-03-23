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
    WithContext(ctx context.Context) RedisClient
    SelectDB(index int) error
    SetKey(key string, value interface{}, expiration time.Duration) error
    SetNx(key string, value interface{}, expiration time.Duration) error
    GetAndDel(key string, value interface{}) error
    GetString(key string) string
    GetBytes(key string) ([]byte, error)
    GetScanStruct(key string, value interface{}) error
}