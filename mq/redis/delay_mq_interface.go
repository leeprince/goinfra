package redis

import "time"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/5/4 下午2:45
 * @Desc:   redis 实现延迟消息队列：多个发布者，一个消费者
 */

var (
    // 默认延迟的时间：1 秒
    DefaulDelayTime = time.Second * 10
    
    // 默认等待时间：1 秒
    DefaultWaitTime = time.Second * 10
)

// 延迟消息队列消费的回调方法
type DelayMQSubscribeHandle func(data []byte)

type DelayMQ interface {
    // 生产
    //  delayTime：延迟的时间。
    Push(key string, value interface{}, delayTime time.Duration) error
    // 消费
    //  waitTime：等待时间。轮询延迟队列无元素后需等待多少时间后继续轮询。增加等待时间，避免无数据时，浪费 cpu。
    Subscribe(f DelayMQSubscribeHandle, key string, waitTime time.Duration)
}
