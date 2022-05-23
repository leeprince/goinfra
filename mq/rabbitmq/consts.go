package rabbitmq

import (
    "time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/5/13 下午3:47
 * @Desc:
 */

// amqp.Delivery.Headers 的 map 的键
const (
    // 关于死信消息
    xDeath = "x-death"
    // 过期时间，单位：毫秒。
    originalExpiration = "original-expiration"
    // 第一次发生死信的原因
    xFirstDeathReason = "x-first-death-reason"
    // 死信的原因:过期
    xFirstDeathReasonOfExpired = "expired"
)

// 连接相关
const (
    // 默认连接的 url
    //  - 该 url 解析出来的 vhost 是空的
    //      - 如果需要通过 url 解析出其他 vhost, 如：`/prince`，则 url="amqp://guest:guest@localhost:5672//prince"
    defaultURL   = "amqp://guest:guest@localhost:5672/"
    defaultVhost = ""
)

// 失败相关
const (
    defaultErrRetryTime = time.Second * 2
)

// 交换机类型相关
const (
    exchangeTypeFanout  = "fanout"  // 发布和订阅
    exchangeTypeDirect  = "direct"  // 路由队列
    exchangeTypeTopic   = "topic"   // 主题队列
    exchangeTypeHeaders = "headers" // RPC队列
)

// 发布消息的投递模式
const (
    PropertiesDeliveryModeTransient  = 1 // 消息临时化：amqp.Transient=1
    PropertiesDeliveryModePersistent = 2 // 消息持久化:amqp.Persistent=2
)
