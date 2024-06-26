package redis

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/cast"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/2/26 下午10:31
 * @Desc:   redis
 *              关于 value 参数为 interface{} 时
 *                  value 为切片：对于lpush(rpush)会当作列表中的多个元素。r.cli.LPush 的 value 支持传入`...interface{}`
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
func InitGoredisList(confs RedisConfs) error {
	clients := make(map[string]*Goredis, len(confs))
	
	ctx := context.Background()
	for name, conf := range confs {
		client, err := InitGoredis(ctx, conf)
		if err != nil {
			return err
		}
		clients[name] = &Goredis{
			ctx: ctx,
			cli: client,
		}
	}
	
	goredis = clients
	
	return nil
}

func InitGoredis(ctx context.Context, conf RedisConf) (*redis.Client, error) {
	if conf.PoolSize <= 0 {
		conf.PoolSize = RedisClientDefautlPoolSize
	}
	if conf.DB < RedisClientMinDB || conf.DB > RedisClientMaxDB {
		conf.DB = RedisClientMinDB
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
		return nil, fmt.Errorf("[InitGoredisList] pong:%s-err:%v", pong, pingErr)
	}
	return client, nil
}

func GetGoredis(name string) *Goredis {
	client, ok := goredis[name]
	if !ok {
		return nil
	}
	return client
}

func (r *Goredis) WithContext(ctx context.Context) RedisClient {
	r.ctx = ctx
	return r
}

func (r *Goredis) SelectDB(index int) error {
	return r.cli.Conn(r.ctx).Select(r.ctx, index).Err()
}

func (r *Goredis) Set(key string, value interface{}, expiration time.Duration) error {
	return r.cli.Set(r.ctx, key, value, expiration).Err()
}

func (r *Goredis) SetNx(key string, value interface{}, expiration time.Duration) (bool, error) {
	boolCmd := r.cli.SetNX(r.ctx, key, value, expiration)
	return boolCmd.Val(), boolCmd.Err()
}

func (r *Goredis) GetAndDel(key string, value interface{}) error {
	ok, err := r.cli.Eval(r.ctx, getAndDelLuaScript, []string{key}, value).Bool()
	if !ok || err != nil {
		return fmt.Errorf("[GetAndDel] Fail key:%v;val:%v;ok:%v;err:%v", key, value, ok, err)
	}
	return nil
}

// 批量删除脚本的散列值：加载Lua脚本到Redis服务器的，并获取该脚本的SHA1散列值以便后续复用。
var delBatchKeySha string

func (r *Goredis) DelBatchKey(keys []string) (deletedCount int64, err error) {
	if delBatchKeySha == "" {
		delBatchKeySha, err = r.cli.ScriptLoad(context.Background(), delBatchKeyLuaScript).Result()
		if err != nil {
			return
		}
	}
	result, err := r.EvalSha(delBatchKeySha, keys)
	if err != nil {
		return
	}
	
	deletedCount, _ = result.(int64)
	return
}

func (r *Goredis) Eval(script string, keys []string, args ...interface{}) (result interface{}, err error) {
	return r.cli.Eval(r.ctx, script, keys, args...).Result()
}

func (r *Goredis) EvalSha(sha string, keys []string, args ...interface{}) (result interface{}, err error) {
	return r.cli.EvalSha(r.ctx, sha, keys, args...).Result()
}

func (r *Goredis) GetString(key string) string {
	return r.cli.Get(r.ctx, key).String()
}

func (r *Goredis) GetInt64(key string) (int64, error) {
	return r.cli.Get(r.ctx, key).Int64()
}

func (r *Goredis) GetBytes(key string) ([]byte, error) {
	return r.cli.Get(r.ctx, key).Bytes()
}

func (r *Goredis) GetScanStruct(key string, value interface{}) error {
	return r.cli.Get(r.ctx, key).Scan(value)
}

// Push 和 Pop 默认的方向是相反的，符合队列的：先进先出
//  value参数说明: 关于 value 说明参考该文件顶部`关于 value 参数为 interface{} 时`
func (r *Goredis) Push(key string, value interface{}, isRight ...bool) error {
	if len(isRight) > 0 && isRight[0] {
		return r.cli.RPush(r.ctx, key, value).Err()
	}
	return r.cli.LPush(r.ctx, key, value).Err()
}

// Push 和 Pop 默认的方向是相反的，符合队列的：先进先出
func (r *Goredis) Pop(key string, isLeft ...bool) (data []byte, err error) {
	if len(isLeft) > 0 && isLeft[0] {
		return r.cli.LPop(r.ctx, key).Bytes()
	}
	return r.cli.RPop(r.ctx, key).Bytes()
}

// Push 和 BPop 默认的方向是相反的，符合队列的：先进先出
func (r *Goredis) BPop(key string, timeout time.Duration, isLeft ...bool) (data interface{}, err error) {
	// keyValueSlic: 0:key 1:value
	var keyValueSlice []string
	if len(isLeft) > 0 && isLeft[0] {
		err = r.cli.BLPop(r.ctx, timeout, key).ScanSlice(&keyValueSlice)
	} else {
		err = r.cli.BRPop(r.ctx, timeout, key).ScanSlice(&keyValueSlice)
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
func (r *Goredis) ZAdd(key string, members ...*Z) error {
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
	return r.cli.ZAdd(r.ctx, key, redisMembers...).Err()
}

func (r *Goredis) ZRangeByScore(key string, opt *ZRangeBy) (data []string, err error) {
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
		zSlice, err = r.cli.ZRangeByScoreWithScores(r.ctx, key, ZRangeBy).Result()
		// fmt.Printf("zSlice type:%T zSlice:%v \n", zSlice, zSlice)
		if err != nil {
			return
		}
		
		for _, i2 := range zSlice {
			data = append(data, cast.ToString(i2.Member), cast.ToString(i2.Score))
		}
		return
	}
	data, err = r.cli.ZRangeByScore(r.ctx, key, ZRangeBy).Result()
	// fmt.Printf("data type:%T data:%v \n", data, data)
	
	return
}

// 移除有序集合中的一个或多个成员
func (r *Goredis) ZRem(key string, members ...interface{}) error {
	if len(members) == 0 {
		return errors.New("(c *Goredis) ZAdd len(members) == 0")
	}
	
	return r.cli.ZRem(r.ctx, key, members...).Err()
}

// 将 key 中储存的数字值增一。
func (r *Goredis) Incr(key string) (int64, error) {
	cmd := r.cli.Incr(r.ctx, key)
	return cmd.Val(), cmd.Err()
}

func (r *Goredis) Decr(key string) (int64, error) {
	cmd := r.cli.Decr(r.ctx, key)
	return cmd.Val(), cmd.Err()
}

// 将 key 所储存的值加上给定的增量值（increment） 。
func (r *Goredis) IncrBy(key string, value int64) (int64, error) {
	cmd := r.cli.IncrBy(r.ctx, key, value)
	return cmd.Val(), cmd.Err()
}

func (r *Goredis) DecrBy(key string, value int64) (int64, error) {
	cmd := r.cli.DecrBy(r.ctx, key, value)
	return cmd.Val(), cmd.Err()
}

func (r *Goredis) Publish(channel string, message interface{}) error {
	return r.cli.Publish(r.ctx, channel, message).Err()
}

func (r *Goredis) Subscribe(channels ...string) *SubscribeMessage {
	subscribeChannel := r.cli.Subscribe(r.ctx, channels...).Channel()
	
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

// 不存在则创建值等于1；存在则+1并且返回+1后的结果:
func (r *Goredis) GetSetIncrLua(key string, expiration time.Duration) (int64, error) {
	// 纳秒转为秒
	expirationSecond := int64(expiration / 1e9)
	
	// lua 索引从1开始
	// 第一次设置值时就开始给该键名设置过期时间
	value, err := r.cli.Eval(r.ctx, `
		local i = redis.call("INCR", KEYS[1])
		if i == 1 then
			redis.call("EXPIRE", KEYS[1], ARGV[1])
		end
		return i
	`,
		[]string{key}, expirationSecond).Int64()
	if err != nil {
		return 0, err
	}
	
	return value, nil
}

// 增量值增加
func (r *Goredis) GetSetIncrByLua(key string, value int64, expiration time.Duration) (int64, error) {
	// 纳秒转为秒
	expirationSecond := int64(expiration / 1e9)
	
	// lua 索引从1开始
	// 第一次设置值时就开始给该键名设置过期时间
	cmd := r.cli.Eval(r.ctx, `
	local i = redis.call("INCRBY", KEYS[1], ARGV[1])
	-- 注意必须将：ARGV[1] 转为数字后再比较
	if i == tonumber(ARGV[1]) then
		redis.call("EXPIRE", KEYS[1], ARGV[2])
	end
	return i
	-- return {ARGV[1], type(ARGV[1]), tonumber(ARGV[1]), i, ARGV[2]}
`,
		[]string{key}, value, expirationSecond)
	if cmd.Err() != nil {
		return 0, cmd.Err()
	}
	
	return cmd.Int64()
}

// 不存在则创建值等于1；存在则+1并且返回+1后的结果。
func (r *Goredis) GetSetIncrTxPipeline(key string, expiration time.Duration) (int64, error) {
	// 支持原子性的pipeline
	pipeliner := r.cli.TxPipeline()
	
	// pipeliner.Exec()执行完之后可以通过incrCmd获取到返回的值
	incrCmd := pipeliner.Incr(r.ctx, key)
	
	// 支持纳秒
	pipeliner.Expire(r.ctx, key, expiration)
	
	// cmders, err := pipeliner.Exec()
	// fmt.Println(cmders)
	// // 按pipeliner执行的顺序放入数组中，即：Incr =》 Expire
	// for _, cmder := range cmders {
	//	if cmder.Err() != nil {
	//		return 0, err
	//	}
	//	fmt.Println(cmder.Name())
	//	fmt.Println(cmder.Args())
	// }
	
	_, err := pipeliner.Exec(r.ctx)
	if err != nil {
		return 0, err
	}
	
	return incrCmd.Val(), nil
}
