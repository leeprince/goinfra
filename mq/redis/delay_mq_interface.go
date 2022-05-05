package redis

import "time"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/5/4 下午2:45
 * @Desc:   redis 实现延迟消息队列
 */

// 默认延迟的时间：1 秒
var DefaulDelayTime = time.Second * 10

// 默认等待时间：1 秒
var DefaultWaitTime = time.Second * 10

type DelayMQ interface {
    // 生产
    //  delayTime：延迟的时间。
    Push(key string, value interface{}, delayTime time.Duration) error
    // 消费
    //  waitTime：等待时间。轮询延迟队列无元素后需等待多少时间后继续轮询。增加等待时间，避免无数据时，浪费 cpu。
    Subscribe(key string, waitTime time.Duration) (data []byte, err error)
}
