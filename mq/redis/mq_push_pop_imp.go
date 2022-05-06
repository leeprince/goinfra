package redis

import (
    "github.com/leeprince/goinfra/plog"
    "github.com/leeprince/goinfra/storage/redis"
    "time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/5/1 上午11:57
 * @Desc:   结合 redis `lpush 和 rpop`(默认) 或者 `rpush 和 lpop` 列表命令实现 MQ 接口
 *              Push：lpush、rpush
 *              Subscribe：rpop、lpop
 */

type ListMQ struct {
    cli redis.RedisClient
}

// ListMQ 实现 MQ 接口
var _ MQ = (*ListMQ)(nil)

func NewListMQ(cli redis.RedisClient) *ListMQ {
    return &ListMQ{
        cli: cli,
    }
}

func (mq *ListMQ) Push(key string, value interface{}) error {
    return mq.cli.Push(key, value)
}
func (mq *ListMQ) Subscribe(f listMQSubscribeFunc, key string, timeout time.Duration) {
    for {
        var data interface{}
        var err error
        
        // 当列表为空时，消费者就会不断的轮训来获取数据，但是每次都获取不到数据，就会陷入一个取不到数据的死循环里，这不仅拉高了客户端的CPU，还拉高了Redis的QPS，并且这些访问都是无效的
        // 解决：通过堵塞从列表中获取
        data, err = mq.cli.BPop(key, timeout)
        if err != nil {
            plog.Error("(mq *ListMQ) Subscribe mq.cli.BPop err:", err)
            // 防止 redis 连接断开(重启、网络抖动)后无限循环，浪费 cpu
            time.Sleep(timeout)
            continue
        }
        if data == nil {
            // fmt.Println("(mq *ListMQ) Subscribe data == nil")
            continue
        }
        
        go f(data)
    }
}