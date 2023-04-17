package redis

import (
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/5/1 上午11:13
 * @Desc:   redis 实现普通消息队列：多个发布者，一个消费者
 */

// 消息队列消费的回调方法
type ListMQSubscribeHandle func(data interface{})

type MQ interface {
	// 生产
	Push(key string, value interface{}) error
	// 消费
	//  timeout：超时时间。超时继续轮询
	Subscribe(f ListMQSubscribeHandle, key string, timeout time.Duration)
}
