package redis

import (
    "context"
    "errors"
    "fmt"
    "github.com/go-redis/redis/v8"
    "github.com/leeprince/goinfra/config"
    "github.com/leeprince/goinfra/consts"
    "github.com/spf13/cast"
    "time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/2/26 下午10:31
 * @Desc:   redis
 *              关于 value 参数为 interface{} 时
 *                  value 为切片：对于lpush(rpush)会当作列表中的多个元素。c.cli.LPush 的 value 支持传入`...interface{}`
 *                  value 为结构体或者部分命令传入切片（ZAdd 方法的 Z.Member 为切片）时：需实现 `encoding.BinaryMarshaler` 接口(MarshalBinary 方法), 否则报错`redis: can't marshal []string (implement encoding.BinaryMarshaler)`。建议直接转成 json string 或者 []byte
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

// Push 和 Pop 默认的方向是相反的，符合队列的：先进先出
//  value参数说明: 关于 value 说明参考该文件顶部`关于 value 参数为 interface{} 时`
func (c *Goredis) Push(key string, value interface{}, isRight ...bool) error {
    if len(isRight) > 0 && isRight[0] {
        return c.cli.RPush(c.ctx, key, value).Err()
    }
    return c.cli.LPush(c.ctx, key, value).Err()
}

// Push 和 Pop 默认的方向是相反的，符合队列的：先进先出
func (c *Goredis) Pop(key string, isLeft ...bool) (data []byte, err error) {
    if len(isLeft) > 0 && isLeft[0] {
        return c.cli.LPop(c.ctx, key).Bytes()
    }
    return c.cli.RPop(c.ctx, key).Bytes()
}

// Push 和 BPop 默认的方向是相反的，符合队列的：先进先出
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

// value参数说明: 关于 value 说明参考该文件顶部`关于 value 参数为 interface{} 时`
func (c *Goredis) ZAdd(key string, members ...*Z) error {
    if len(members) == 0 {
        return errors.New("(c *Goredis) ZAdd len(members) == 0")
    }
    
    var redisMembers []*redis.Z
    for _, i2 := range members {
        redisMembers = append(redisMembers, &redis.Z{
            Score:  i2.Score,
            Member: i2.Member,
        })
    }
    return c.cli.ZAdd(c.ctx, key, redisMembers...).Err()
}

func (c *Goredis) ZRangeByScore(key string, opt *ZRangeBy) (data []string, err error) {
    if opt.Max == "" {
        err = errors.New("opt.Max can not empty")
        return
    }
    if opt.Count == 0 || opt.Count > ZRangeByMaxCount {
        opt.Count = ZRangeByMaxCount
    }
    ZRangeBy := &redis.ZRangeBy{
        Min:    opt.Min,
        Max:    opt.Max,
        Offset: opt.Offset,
        Count:  opt.Count,
    }
    
    // 返回分数的格式为：[]string{成员1 分数1 成员2 分数2}。
    //  []string 返回的顺序与 redigo 兼容。
    if opt.isReturnScore {
        var zSlice []redis.Z
        zSlice, err = c.cli.ZRangeByScoreWithScores(c.ctx, key, ZRangeBy).Result()
        // fmt.Printf("zSlice type:%T zSlice:%v \n", zSlice, zSlice)
        if err != nil {
            return
        }
        
        for _, i2 := range zSlice {
            data = append(data, cast.ToString(i2.Member), cast.ToString(i2.Score))
        }
        return
    }
    data, err = c.cli.ZRangeByScore(c.ctx, key, ZRangeBy).Result()
    // fmt.Printf("data type:%T data:%v \n", data, data)
    
    return
}

func (c *Goredis) ZRem(key string, members ...interface{}) error {
    if len(members) == 0 {
        return errors.New("(c *Goredis) ZAdd len(members) == 0")
    }
    
    return c.cli.ZRem(c.ctx, key, members...).Err()
}

func (c *Goredis) Publish(channel string, message interface{}) error {
    return c.cli.Publish(c.ctx, channel, message).Err()
}

func (c *Goredis) Subscribe(channels ...string) *SubscribeMessage {
    subscribeChannel := c.cli.Subscribe(c.ctx, channels...).Channel()
    
    // for 与 select...case... 都能接收通道（channel）的数据
    // for 通道（channel）只有一个参数, 并且需要返回时外面也需要 return
    /*for channel := range subscribeChannel {
        return &SubscribeMessage{
            Channel:      channel.Channel,
            Pattern:      channel.Pattern,
            Payload:      channel.Payload,
            PayloadSlice: channel.PayloadSlice,
        }
    }
    return nil*/
    
    select {
    case msg := <-subscribeChannel:
        return &SubscribeMessage{
            Channel:      msg.Channel,
            Pattern:      msg.Pattern,
            Payload:      msg.Payload,
            PayloadSlice: msg.PayloadSlice,
        }
    }
}
