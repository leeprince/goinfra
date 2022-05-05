package redis

import (
    "context"
    "errors"
    "fmt"
    "github.com/gomodule/redigo/redis"
    "github.com/leeprince/goinfra/config"
    "github.com/leeprince/goinfra/utils"
    "github.com/spf13/cast"
    "time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/2/27 上午12:22
 * @Desc:   redigo
 *              关于有序结合 member 参数 *Z
 *                  不支持 Z.Member 为切片、结构体
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

func (c *Redigo) Set(key string, value interface{}, expiration time.Duration) error {
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

// Push 和 Pop 默认的方向是相反的，符合入队和出队的：先进先出
// 如果 value 为切片，则只是插入一个数组
// 解决：通过 `redis.Args{}.Add(key).AddFlat(value)...` 的方式插入多个元素
//  如：value = []string{"v001", "v002"}
//      redisPool.Do("LPUSH", key, value): 插入后为：`[v001 v002]` 的一个元素
//      redisPool.Do("LPUSH", redis.Args{}.Add(key).AddFlat(value)...)：插入后为：`v001`,`v002` 的两个元素
func (c *Redigo) Push(key string, value interface{}, isRight ...bool) error {
    redisPool := c.Get()
    defer redisPool.Close()
    
    if len(isRight) > 0 && isRight[0] {
        _, err := redisPool.Do("RPUSH", key, value)
        return err
    }
    
    // _, err := redisPool.Do("LPUSH", key, value)
    _, err := redisPool.Do("LPUSH", redis.Args{}.Add(key).AddFlat(value)...)
    return err
}

// Push 和 Pop 默认的方向是相反的，符合入队和出队的：先进先出
func (c *Redigo) Pop(key string, isLeft ...bool) (data []byte, err error) {
    redisPool := c.Get()
    defer redisPool.Close()
    
    if len(isLeft) > 0 && isLeft[0] {
        return redis.Bytes(redisPool.Do("LPOP", key))
    }
    return redis.Bytes(redisPool.Do("RPOP", key))
}

// Push 和 BPop 默认的方向是相反的，符合入队和出队的：先进先出
//  timeout 应小于等于 redis.Pool 设置的超时时间，否则会报 `i/o timeout`
func (c *Redigo) BPop(key string, timeout time.Duration, isLeft ...bool) (data interface{}, err error) {
    redisPool := c.Get()
    defer redisPool.Close()
    
    // keyValueSlic: 0:key 1:value
    var keyValueSlice []interface{}
    if len(isLeft) > 0 && isLeft[0] {
        keyValueSlice, err = redis.Values(redisPool.Do("BLPOP", key, timeout.Seconds()))
    } else {
        keyValueSlice, err = redis.Values(redisPool.Do("BRPOP", key, timeout.Seconds()))
    }
    // fmt.Printf("(c *Redigo) BPop error = %v, keyValueSlice=%v, keyValueSlice Type=%T \n", err, keyValueSlice, keyValueSlice)
    // fmt.Println(cast.ToString(keyValueSlice[0]), cast.ToString(keyValueSlice[1]))
    
    if err == redis.ErrNil {
        err = nil
        return
    }
    if err != nil {
        return
    }
    if len(keyValueSlice) != 2 {
        err = errors.New("Redigo@BPop keyValueSlice len not equal to 2")
        return
    }
    data = keyValueSlice[1]
    return
}

// 不支持 Z.Member 为切片、结构体
func (c *Redigo) ZAdd(key string, members ...*Z) error {
    if len(members) == 0 {
        return errors.New("(c *Redigo) ZAdd len(members) == 0")
    }
    
    redisPool := c.Get()
    defer redisPool.Close()
    
    args := redis.Args{}.Add(key)
    for _, i2 := range members {
        args = args.AddFlat(i2.Score)
        args = args.AddFlat(i2.Member)
    }
    _, err := redisPool.Do("ZADD", args...)
    
    return err
}

func (c *Redigo) ZRangeByScore(key string, opt *ZRangeBy) (data []string, err error) {
    if opt.Max == "" {
        err = errors.New("opt.Max can not empty")
        return
    }
    if opt.Count == 0 || opt.Count > ZRangeByMaxCount  {
        opt.Count = ZRangeByMaxCount
    }
    
    redisPool := c.Get()
    defer redisPool.Close()
    
    var sliceInterface []interface{}
    
    // 返回分数的格式为：[]string{成员1 分数1 成员2 分数2}
    if opt.isReturnScore {
        sliceInterface, err = redis.Values(redisPool.Do("ZRANGEBYSCORE", key, opt.Min, opt.Max, "WITHSCORES", "LIMIT", opt.Offset, opt.Count))
    } else {
        sliceInterface, err = redis.Values(redisPool.Do("ZRANGEBYSCORE", key, opt.Min, opt.Max, "LIMIT", opt.Offset, opt.Count))
    }
    // fmt.Printf("sliceInterface data type:%T data:%v \n", sliceInterface, sliceInterface)
    
    if err == redis.ErrNil {
        err = nil
        return
    }
    
    for _, i2 := range sliceInterface {
        data = append(data, cast.ToString(i2))
    }
    
    // fmt.Printf("data type:%T data:%v \n", data, data)
    return
}

func (c *Redigo) ZRem(key string, members ...interface{}) error {
    if len(members) == 0 {
        return errors.New("(c *Redigo) ZAdd len(members) == 0")
    }
    
    redisPool := c.Get()
    defer redisPool.Close()
    
    _, err := redisPool.Do("ZREM", redis.Args{}.Add(key).AddFlat(members)...)
    
    return err
}
