package redis

import (
    "fmt"
    "github.com/gomodule/redigo/redis"
    "goinfra/config"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/2/27 上午12:22
 * @Desc:
 */

func InitRedisConn(confs config.RedisConfs) (conns map[string]redis.Conn, err error) {
    conns = make(map[string]redis.Conn, len(confs))
    for name, conf := range confs {
        conn, connErr := redis.Dial(conf.Network, conf.Addr)
        if connErr != nil {
            err = fmt.Errorf("[InitRedisClient] name:%s-err:%v", name, connErr)
            return
        }
        conns[name] = conn
    }
    return
}