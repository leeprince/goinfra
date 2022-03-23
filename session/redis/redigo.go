package redis

import (
    "context"
    "fmt"
    "github.com/gomodule/redigo/redis"
    "github.com/leeprince/goinfra/config"
    "time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/2/27 上午12:22
 * @Desc:
 */

type Redigo struct {
    redis.Conn
}

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

func (c *Redigo) SetNX(ctx context.Context, key string, value interface{}, expiration time.Duration) bool {
    return true
}

func (c *Redigo) GetAndDel(ctx context.Context, key string, value interface{}) bool {
    return true
}

func (c *Redigo) GetString(ctx context.Context, key string) string {
    return ""
}

func (c *Redigo) GetBytes(ctx context.Context, key string) []byte {
    return nil
}

func (c *Redigo) GetScan(ctx context.Context, key string, value interface{}) error {
    return nil
}