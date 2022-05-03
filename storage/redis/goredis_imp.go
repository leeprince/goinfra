package redis

import (
    "context"
    "errors"
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
var goredis map[string]*Goredis

// Redigo 实现 RedisClient 接口
var _ RedisClient = (*Goredis)(nil)

// redis 客户端结构体
type Goredis struct {
    ctx context.Context
    cli *redis.Client
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
            conf.DB = consts.RedisClientMinDB
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
            ctx: ctx,
            cli: client,
        }
    }
    
    goredis = clients
    
    return nil
}

func GetGoredis(name string) *Goredis {
    client, ok := goredis[name]
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
    return c.cli.Conn(c.ctx).Select(c.ctx, index).Err()
}

func (c *Goredis) Set(key string, value interface{}, expiration time.Duration) error {
    return c.cli.Set(c.ctx, key, value, expiration).Err()
}

func (c *Goredis) SetNx(key string, value interface{}, expiration time.Duration) (bool, error) {
    boolCmd := c.cli.SetNX(c.ctx, key, value, expiration)
    return boolCmd.Val(), boolCmd.Err()
}

func (c *Goredis) GetAndDel(key string, value interface{}) error {
    ok, err := c.cli.Eval(c.ctx, luaRedisGetAndDelScript, []string{key}, value).Bool()
    if !ok || err != nil {
        return fmt.Errorf("[GetAndDel] Fail key:%v;val:%v;ok:%v;err:%v", key, value, ok, err)
    }
    return nil
}

func (c *Goredis) GetString(key string) string {
    return c.cli.Get(c.ctx, key).String()
}

func (c *Goredis) GetBytes(key string) ([]byte, error) {
    return c.cli.Get(c.ctx, key).Bytes()
}

func (c *Goredis) GetScanStruct(key string, value interface{}) error {
    return c.cli.Get(c.ctx, key).Scan(value)
}

// Push 和 Pop 默认的方向是相反的，符合入队和出队的：先进先出
//  value参数说明
//      value 为切片：当作列表中的多个参数
//      value 为结构体：需结构体实现 `encoding.BinaryMarshaler` 接口(MarshalBinary 方法)。建议直接转成 string 或者 []byte
func (c *Goredis) Push(key string, value interface{}, isRight ...bool) error {
    if len(isRight) > 0 && isRight[0] {
        return c.cli.RPush(c.ctx, key, value).Err()
    }
    return c.cli.LPush(c.ctx, key, value).Err()
}

// Push 和 Pop 默认的方向是相反的，符合入队和出队的：先进先出
func (c *Goredis) Pop(key string, isLeft ...bool) (data []byte, err error) {
    if len(isLeft) > 0 && isLeft[0] {
        return c.cli.LPop(c.ctx, key).Bytes()
    }
    return c.cli.RPop(c.ctx, key).Bytes()
}

// Push 和 BPop 默认的方向是相反的，符合入队和出队的：先进先出
func (c *Goredis) BPop(key string, timeout time.Duration, isLeft ...bool) (data interface{}, err error) {
    // keyValueSlic: 0:key 1:value
    var keyValueSlice []string
    if len(isLeft) > 0 && isLeft[0] {
        err = c.cli.BLPop(c.ctx, timeout, key).ScanSlice(&keyValueSlice)
    } else {
        err = c.cli.BRPop(c.ctx, timeout, key).ScanSlice(&keyValueSlice)
    }
    // fmt.Printf("(c *Goredis) BPop error = %v, keyValueSlice=%v, keyValueSlice Type=%T \n", err, keyValueSlice, keyValueSlice)
    // fmt.Println(cast.ToString(keyValueSlice[0]), cast.ToString(keyValueSlice[1]))
    
    if err != nil {
        return
    }
    if keyValueSlice == nil {
        return
    }
    if len(keyValueSlice) != 2 {
        err = errors.New("Goredis@BPop len(keyValueSlice) != 2")
        return
    }
    data = keyValueSlice[1]
    return
}
