package redis

import (
    "context"
    "time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/3/23 上午2:15
 * @Desc:   redis 接口：具体实现依赖：github.com/go-redis/redis/v8、github.com/gomodule/redigo/redis
 */

type RedisClient interface {
    // 设置上下文
    WithContext(ctx context.Context) RedisClient
    // 选择 redis DB 库
    SelectDB(index int) error
    // 设置
    Set(key string, value interface{}, expiration time.Duration) error
    SetNx(key string, value interface{}, expiration time.Duration) (bool, error)
    GetAndDel(key string, value interface{}) error
    GetString(key string) string
    GetBytes(key string) ([]byte, error)
    GetScanStruct(key string, value interface{}) error
    Push(key string, value interface{}, isRight ...bool) error
    Pop(key string, isLeft ...bool) (data []byte, err error)
    BPop(key string, timeout time.Duration, isLeft ...bool) (data interface{}, err error)
}