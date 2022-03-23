package redis

import (
    "context"
    "fmt"
    "github.com/go-redis/redis/v8"
    "github.com/leeprince/goinfra/config"
    "github.com/leeprince/goinfra/consts"
    "time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/2/26 下午10:31
 * @Desc:   redis
 */

// 全局变量
var gorediss map[string]*Goredis

// Redigo 实现 RedisClient 接口
var _ RedisClient = (*Goredis)(nil)

// redis 客户端结构体
type Goredis struct {
    ctx context.Context
    *redis.Client
}

// 初始化
func InitGoredis(confs config.RedisConfs) error {
    ctx := context.Background()
    clients := make(map[string]*Goredis, len(confs))
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
            return fmt.Errorf("[InitGoredis] name:%s-pong:%s-err:%v", name, pong, pingErr)
        }
        
        clients[name] = &Goredis{
            ctx:    ctx,
            Client: client,
        }
    }
    
    gorediss = clients
    
    return nil
}

func GetGoredis(name string) *Goredis {
    client, ok := gorediss[name]
    if !ok {
        return nil
    }
    return client
}

func (c *Goredis) WithContext(ctx context.Context) RedisClient {
    c.ctx = ctx
    return c
}

func (c *Goredis) SelectDB(index int) error {
    return c.Conn(c.ctx).Select(c.ctx, index).Err()
}

func (c *Goredis) SetKey(key string, value interface{}, expiration time.Duration) error {
    return c.Set(c.ctx, key, value, expiration).Err()
}

// TODO: 是否成功，而不是是否有错误。不成功，有可能无错误 - prince@todo 2022/3/24 上午1:55
func (c *Goredis) SetNx(key string, value interface{}, expiration time.Duration) error {
    return c.SetNX(c.ctx, key, value, expiration).Err()
}

func (c *Goredis) GetAndDel(key string, value interface{}) error {
    val, err := c.Eval(c.ctx, luaRedisScript, []string{key}, value).Result()
    if val == nil || err != nil {
        return fmt.Errorf("[unLock] Fail.val:%v;err:%v", val, err)
    }
    return nil
}

func (c *Goredis) GetString(key string) string {
    return c.Get(c.ctx, key).String()
}

func (c *Goredis) GetBytes(key string) ([]byte, error) {
    return c.Get(c.ctx, key).Bytes()
}

func (c *Goredis) GetScanStruct(key string, value interface{}) error {
    return c.Get(c.ctx, key).Scan(value)
}
