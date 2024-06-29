package redis

import "time"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/3/23 下午10:24
 * @Desc:
 */

var getAndDelLuaScript = `
    if redis.call("GET", KEYS[1]) ~= ARGV[1] then
        return nil
    end
    return redis.call("DEL", KEYS[1])
`

// delBatchKeyLuaScript 删除多个keys的Lua脚本
/*
	1、初始化迭代变量i为1。
	2、然后检查条件i <= #keys（其中#keys是keys表的长度），如果条件为真，则执行循环体内的代码。
	3、循环体内调用redis.call("DEL", keys[i])来删除对应索引的key。
	4、在每次迭代结束后，i的值自动加1（这是for循环默认的行为）。
	5、当i的值超过#keys时，循环结束。
*/
var delBatchKeyLuaScript = `
	local keys = KEYS
	for i=1,#keys do
		redis.call("DEL", keys[i])
	end
	return #keys -- 返回成功删除的key数量
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
