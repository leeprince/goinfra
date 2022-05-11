package redis

import "github.com/leeprince/goinfra/storage/redis"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/5/10 下午5:38
 * @Desc:   redis 实现发布订阅者模型：多个发布者，多个消费者订阅
 *              当然如果我们会发布大量的消息， 同时会有多个消费者去消费，也可以将通道分成多个，
 *              每个通道有自己的订阅者订阅，然后发布者在发布消息的时候根据节点ID或随机分配的方式分配到每个通道上来实现。
 *              》注意，因为 redis 发布订阅没有缓存，一定要先使订阅者订阅到频道后，再有发布操作，否则发布后的消息没被订阅会丢失。
 */



// 发布订阅者模型订阅的回调方法
type pubishSubscribeMQSubscribeFunc func(data *redis.SubscribeMessage)

type PubSubMQ interface {
    // 发布
    Push(channel string, message interface{}) error
    // 订阅
    Subscribe(f pubishSubscribeMQSubscribeFunc, channels ...string)
}