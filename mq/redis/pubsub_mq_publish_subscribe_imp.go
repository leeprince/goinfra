package redis

import (
    "github.com/leeprince/goinfra/storage/redis"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/5/10 下午5:48
 * @Desc:   通过 redis `publish`、`subscribe` 发布订阅命令实现 PubSubMQ 接口
 */

type PubishSubscribeMQ struct {
    cli redis.RedisClient
}

// PubishSubscribeMQ 实现 PubSubMQ 接口
var _ PubSubMQ = (*PubishSubscribeMQ)(nil)

func NewPubishSubscribeMQ(cli redis.RedisClient) *PubishSubscribeMQ {
    return &PubishSubscribeMQ{
        cli: cli,
    }
}

func (mq *PubishSubscribeMQ) Push(channel string, message interface{}) error {
    return mq.cli.Publish(channel, message)
}

func (mq *PubishSubscribeMQ) Subscribe(f PubishSubscribeMQSubscribeHandle, channels ...string) {
    // 启动一个守护进程去订阅消息
    for {
        msgChan := mq.cli.Subscribe(channels...)
    
        // 非堵塞进程处理消息
        go f(msgChan)
    }
}
