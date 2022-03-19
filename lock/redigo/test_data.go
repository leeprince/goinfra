package redigo

import (
    "github.com/leeprince/goinfra/config"
    "time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/3/3 下午9:05
 * @Desc:   测试数据
 */

var (
    RedisName = "local"
    RedisConfs = config.RedisConfs{
        RedisName: config.RedisConf{
            Network:  "tcp",
            Addr:     "127.0.0.1:6379",
            Username: "",
            Password: "",
            DB:       0,
            PoolSize: 1,
        },
    }
    LockKey = "princeLockKey01"
    LockValue = "princeLockValue01"
    LockExpire = time.Second * 2
)
