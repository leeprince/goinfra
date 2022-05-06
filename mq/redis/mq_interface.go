package redis

import (
    "time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/5/1 上午11:13
 * @Desc:   redis 实现消息队列
 */

// 消息队列回调方法
type listMQSubscribeFunc func(data interface{})

type MQ interface {
    // 生产
    Push(key string, value interface{}) error
    // 消费
    //  timeout：超时时间。超时继续轮询
    Subscribe(f listMQSubscribeFunc, key string, timeout time.Duration)
}