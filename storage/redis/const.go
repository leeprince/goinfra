package redis

import "time"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/3/23 下午10:24
 * @Desc:
 */

var luaRedisGetAndDelScript = `
    if redis.call("GET", KEYS[1]) ~= ARGV[1] then
        return nil
    end
    return redis.call("DEL", KEYS[1])
`

var luaIncrExpireScript = `
	local i = redis.call("INCR", KEYS[1])
	redis.call("EXPIRE", KEYS[1], ARGV[1])
	return i
`

const (
	RedisClientDefautlPoolSize = 10 // 连接池数量
	RedisClientMinDB           = 0  // 库:0~15
	RedisClientMaxDB           = 15 // 库:0~15
)

const (
	redigoDialConnectTimeout = time.Second * 1
	redigoDialReadTimeout    = time.Second * 10
	redigoDialWriteTimeout   = time.Second * 10
	redigoMaxIdle            = 1000
	redigoMaxActive          = 100
	redigoIdleTimeout        = time.Second * 60
	redigoMaxConnLifetime    = time.Second * 60
)

const (
	redigoStringOk = "OK"
)
