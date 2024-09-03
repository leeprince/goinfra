package redis

import (
	"context"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/3/23 上午2:15
 * @Desc:   redis 接口：具体实现依赖：github.com/go-redis/redis/v8、github.com/gomodule/redigo/redis
 */

// redis 配置
type RedisConfs map[string]RedisConf
type RedisConf struct {
	Network  string // 网络协议：tcp、unix
	Addr     string // 地址：127.0.0.1:6379
	Username string // 用户名
	Password string // 密码
	DB       int    // 库:0~15
	PoolSize int    // 连接池数量
}

// 有序集合的成员
type Z struct {
	Score  float64
	Member interface{}
}

// 有序集合的区间条件
type ZRangeBy struct {
	Min, Max      string
	Offset, Count int64
	isReturnScore bool // 是否返回成员的分数
}

const (
	// 有序集合最大的返回数量，超过需程序兼容
	ZRangeByMaxCount = 10000
)

// 订阅频道获取到的数据
type SubscribeMessage struct {
	Channel      string
	Pattern      string
	Payload      string
	PayloadSlice []string
}

type RedisClient interface {
	// 上下文：设置上下文
	WithContext(ctx context.Context) RedisClient
	// SelectDB DB库：选择 redis DB 库
	SelectDB(index int) error
	// Set 字符串：设置指定 key 的值
	Set(key string, value interface{}, expiration time.Duration) error
	// SetNx 字符串：将值 value 关联到 key ，并将 key 的过期时间设为 seconds (以秒/毫秒为单位)。
	SetNx(key string, value interface{}, expiration time.Duration) (bool, error)
	// GetAndDel 字符串：获取指定键名，并将键值与 value 对比，一致则删除，是通过 lua 脚本执行，做到原子性
	GetAndDel(key string, value interface{}) error
	// GetString 字符串：获取指定键名的字符串
	GetString(key string) string
	// GetBytes 字符串：获取指定键名的字节切片
	GetBytes(key string) ([]byte, error)
	// GetScanStruct 字符串：获取指定键名的键值，并转化为指定数据结构
	GetScanStruct(key string, value interface{}) error
	// Push 列表：将一个或多个值插入到列表头部（最左，默认）或者尾部（最右）
	Push(key string, value interface{}, isRight ...bool) error
	// Pop 列表：移除列表的最左或者最右（默认）一个元素，返回值为移除的元素。
	Pop(key string, isLeft ...bool) (data []byte, err error)
	// BPop 列表：移出并获取列表的第一个元素， 如果列表没有元素会阻塞（设置堵塞时间，不建议设置为0，防止堵塞时间过长，redis 服务端主动断开连接，导致数据丢失）列表直到等待超时或发现可弹出元素为止。
	BPop(key string, timeout time.Duration, isLeft ...bool) (data interface{}, err error)
	// ZAdd 有序集合：向有序集合添加一个或多个成员，或者更新已存在成员的分数。完整的命令：ZADD key [NX|XX] [CH] [INCR] score member [score member ...] 可选项都是默认的
	ZAdd(key string, members ...*Z) error
	// ZRangeByScore 有序集合：通过分数返回有序集合指定区间内的成员(默认从低到高排序，从高到低排序需使用：ZREVRANGEBYSCORE)
	ZRangeByScore(key string, opt *ZRangeBy) (data []string, err error)
	// ZRem 有序集合：移除有序集合中的一个或多个成员
	ZRem(key string, members ...interface{}) error
	// Publish 将信息发送到指定的频道
	Publish(channel string, message interface{}) error
	// Subscribe 订阅给定的一个或多个频道的信息
	Subscribe(channels ...string) *SubscribeMessage
}
