package redis

import (
    "context"
    "fmt"
    "github.com/gomodule/redigo/redis"
    "github.com/leeprince/goinfra/config"
    "github.com/leeprince/goinfra/utils"
    "time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/2/27 上午12:22
 * @Desc:
 */

// 全局变量
var redigos map[string]*Redigo

// Redigo 实现 RedisClient 接口
var _ RedisClient = (*Redigo)(nil)

type Redigo struct {
    ctx context.Context
    redis.Pool
}

func InitRedigo(confs config.RedisConfs) error {
    ctx := context.Background()
    
    redigos = make(map[string]*Redigo, len(confs))
    for name, conf := range confs {
        dialFunc := func() (redis.Conn, error) {
            conn, connErr := redis.Dial(
                conf.Network,
                conf.Addr,
                redis.DialConnectTimeout(redigoDialConnectTimeout),
                redis.DialReadTimeout(redigoDialReadTimeout),
                redis.DialWriteTimeout(redigoDialWriteTimeout),
            )
            if connErr != nil {
                return nil, fmt.Errorf("[InitGoredis] name:%s-err:%v", name, connErr)
            }
            return conn, nil
        }
        
        pool := redis.Pool{
            Dial:            dialFunc,
            DialContext:     nil,
            TestOnBorrow:    nil,
            MaxIdle:         redigoMaxIdle,
            MaxActive:       redigoMaxActive,
            IdleTimeout:     redigoIdleTimeout,
            Wait:            false,
            MaxConnLifetime: redigoMaxConnLifetime,
        }
        
        redigos[name] = &Redigo{
            ctx:  ctx,
            Pool: pool,
        }
    }
    
    return nil
}

func GetRedigo(name string) *Redigo {
    redigo, ok := redigos[name]
    if !ok {
        return nil
    }
    return redigo
}

func (c *Redigo) WithContext(ctx context.Context) RedisClient {
    c.ctx = ctx
    return c
}

func (c *Redigo) SelectDB(index int) error {
    redisPool := c.Get()
    defer redisPool.Close()
    
    if _, err := redisPool.Do("SELECT", index); err != nil {
        return err
    }
    return nil
}

func (c *Redigo) SetKey(key string, value interface{}, expiration time.Duration) error {
    redisPool := c.Get()
    defer redisPool.Close()
    
    if _, err := redisPool.Do("SET", key, value, "EX", int(expiration)); err != nil {
        return err
    }
    return nil
}

// NX:单位秒;PX:单位毫秒
func (c *Redigo) SetNx(key string, value interface{}, expiration time.Duration) (bool, error) {
    redisPool := c.Get()
    defer redisPool.Close()
    
    var str string
    var err error
    if utils.UseMillisecondUnit(expiration) {
        str, err = redis.String(redisPool.Do("SET", key, value, "PX", expiration.Milliseconds(), "NX"))
    } else {
        str, err = redis.String(redisPool.Do("SET", key, value, "EX", expiration.Seconds(), "NX"))
    }
    // 如果 key 已存在则返回 err == redis.ErrNil
    if err == redis.ErrNil {
        return false, nil
    }
    if err != nil {
        return false, err
    }
    if str != redigoStringOk {
        return false, nil
    }
    
    return true, nil
}

func (c *Redigo) GetAndDel(key string, value interface{}) error {
    redisPool := c.Get()
    defer redisPool.Close()
    
    s := redis.NewScript(1, luaRedisGetAndDelScript)
    ok, err := redis.Bool(s.Do(redisPool, key, value))
    
    if !ok || err != nil {
        return fmt.Errorf("[GetAndDel] Fail key:%v;val:%v;ok:%v;err:%v", key, value, ok, err)
    }
    return nil
}

func (c *Redigo) GetString(key string) string {
    redisPool := c.Get()
    defer redisPool.Close()
    
    str, err := redis.String(redisPool.Do("GET", key))
    if err != nil {
        return ""
    }
    return str
}

func (c *Redigo) GetBytes(key string) ([]byte, error) {
    redisPool := c.Get()
    defer redisPool.Close()
    
    str, err := redis.Bytes(redisPool.Do("GET", key))
    return str, err
}

func (c *Redigo) GetScanStruct(key string, value interface{}) error {
    redisPool := c.Get()
    defer redisPool.Close()
    
    valInterface, err := redisPool.Do("GET", key)
    if err != nil {
        return err
    }
    err = redis.ScanStruct([]interface{}{valInterface}, value)
    if err != nil {
        return err
    }
    return nil
}
