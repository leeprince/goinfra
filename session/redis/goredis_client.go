package redis

import (
    "context"
    "fmt"
    "github.com/go-redis/redis/v8"
    "github.com/leeprince/goinfra/config"
    "github.com/leeprince/goinfra/constants"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/2/26 下午10:31
 * @Desc:   redis
 */

func InitRedisClient(ctx context.Context, confs config.RedisConfs) (clients map[string]*redis.Client, err error) {
    clients = make(map[string]*redis.Client, len(confs))
    for name, conf := range confs {
        if conf.PoolSize <= 0 {
            conf.PoolSize = consts.RedisClientDefautlPoolSize
        }
        if conf.DB < consts.RedisClientMinDB || conf.DB > consts.RedisClientMaxDB {
            conf.PoolSize = consts.RedisClientMinDB
        }
        client := redis.NewClient(&redis.Options{
            Network:  conf.Network,
            Addr:     conf.Addr,
            Username: conf.Username,
            Password: conf.Password,
            DB:       conf.DB,
            PoolSize: conf.PoolSize,
        })
        pong, pingErr := client.Ping(ctx).Result()
        if pingErr != nil {
            err = fmt.Errorf("[InitRedisClient] name:%s-pong:%s-err:%v", name, pong, pingErr)
            return
        }
        clients[name] = client
    }
    return
}
