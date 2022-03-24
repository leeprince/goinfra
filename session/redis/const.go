package redis

import "time"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/3/23 下午10:24
 * @Desc:
 */

var luaRedisScript = `
    if redis.call("GET", KEYS[1]) ~= ARGV[1] then
        return nil
    end
    return redis.call("DEL", KEYS[1])
`

const (
    redigoDialConnectTimeout = time.Second * 1
    redigoDialReadTimeout = time.Second * 1
    redigoDialWriteTimeout = time.Second * 10
    redigoMaxIdle = 1000
    redigoMaxActive = 100
    redigoIdleTimeout = time.Second * 60
    redigoMaxConnLifetime = time.Second * 60
)