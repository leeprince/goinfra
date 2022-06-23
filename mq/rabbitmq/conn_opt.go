package rabbitmq

import (
    "time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/5/12 上午10:42
 * @Desc:   初始化客户端的选项
 */

type rabbitMQConf struct {
    // RabbitMQ 连接 url. 格式：amqp协议(固定amqp)//用户名(默认guest):密码(默认guest)@主机(默认localhost):端口(默认5672)。完整示例：amqp://guest:guest@127.0.0.1:5672/
    url string
    // 虚拟主机。默认/
    vhost string
    
    // 强制取消声明交换机
    cancelexchangeDeclare bool
    // 声明交换机
    exchangeDeclare *exchangeDeclare
    
    // 强制取消声明队列
    cancelQueueDeclare bool
    // 声明队列
    queueDeclare *queueDeclare
    
    // 绑定交换机与队列的键
    routingKey string
    
    // 公平调度机制
    qos qos
    // 发生错误时，重试的等待时间
    errRetryTime time.Duration
}

// 声明交换机
type exchangeDeclare struct {
    exchangeName string
    exchangeType string
    passive      bool
    durable      bool
    autoDelete   bool
    internal     bool
    noWait       bool
    arguments    map[string]interface{}
}

// 声明队列
type queueDeclare struct {
    // 队列名
    queueName string
    
    // 队列是否持久化(RabbitMQ管理后台feature=D)
    durable bool
    
    // 是否当前连接通道独占该队列(RabbitMQ管理后台feature=Excl)
    // 如果想拥有私人队列只为一个消费者服务，可以设置 exclusive 参数为 true
    // 如果需要临时队列, 则结合exclusive和autoDelete都设置为true(RabbitMQ管理后台feature=AD,Excl)
    exclusive bool
    
    // 是否自动删除(RabbitMQ管理后台feature=AD).autoDelete=true在消费者取消订阅时，会自动删除。
    autoDelete bool
    
    // 是否不等待
    noWait bool
    
    // map 参数
    arguments map[string]interface{}
}

type qos struct {
    prefetchCount int
    prefetchSize  int
    global        bool
}

// --- ConfOption
type ConfOption func(conf *rabbitMQConf) (err error)

func WithUrl(url string) ConfOption {
    return func(conf *rabbitMQConf) (err error) {
        conf.url = url
        return
    }
}

func WithVhost(vhost string) ConfOption {
    return func(conf *rabbitMQConf) (err error) {
        conf.vhost = vhost
        return
    }
}

func WithErrRetryTime(t time.Duration) ConfOption {
    return func(conf *rabbitMQConf) (err error) {
        conf.errRetryTime = t
        return
    }
}

func WithCancelQueueDeclare(cancel bool) ConfOption {
    return func(conf *rabbitMQConf) (err error) {
        conf.cancelQueueDeclare = cancel
        return
    }
}

func WithQos(prefetchCount, prefetchSize int, global bool) ConfOption {
    return func(conf *rabbitMQConf) (err error) {
        conf.qos = qos{
            prefetchCount: prefetchCount,
            prefetchSize:  prefetchSize,
            global:        global,
        }
        return
    }
}

func WithExchangeDeclare(exchangeName, exchangeType string, opts ...ExchangeDeclareOption) ConfOption {
    return func(conf *rabbitMQConf) (err error) {
        exchangeDeclare := &exchangeDeclare{
            exchangeName: exchangeName,
            exchangeType: exchangeType,
            passive:      false,
            durable:      true,
            autoDelete:   false,
            internal:     false,
            noWait:       false,
            arguments:    nil,
        }
        
        for _, opt := range opts {
            opt(exchangeDeclare)
        }
        
        conf.exchangeDeclare = exchangeDeclare
        return
    }
}

func WithQueueDeclare(queueName string, opts ...QueueDeclareOption) ConfOption {
    return func(conf *rabbitMQConf) (err error) {
        queueDeclare := &queueDeclare{
            queueName:  queueName,
            durable:    true, // 队列是否持久化.默认持久化
            exclusive:  false,
            autoDelete: false,
            noWait:     false,
            arguments:  nil,
        }
        
        for _, opt := range opts {
            opt(queueDeclare)
        }
        
        conf.queueDeclare = queueDeclare
        
        return
    }
}

func WithRoutingKey(routingKey string) ConfOption {
    return func(conf *rabbitMQConf) (err error) {
        conf.routingKey = routingKey
        return
    }
}

// --- ConfOption -end

// --- WithExchangeDeclare ExchangeDeclareOption
type ExchangeDeclareOption func(exchangeDeclare *exchangeDeclare)

func WithExchangeDeclarePassive(passive bool) ExchangeDeclareOption {
    return func(exchangeDeclare *exchangeDeclare) {
        exchangeDeclare.passive = passive
    }
}
func WithExchangeDeclareDurable(durable bool) ExchangeDeclareOption {
    return func(exchangeDeclare *exchangeDeclare) {
        exchangeDeclare.durable = durable
    }
}
func WithExchangeDeclareAutoDelete(autoDelete bool) ExchangeDeclareOption {
    return func(exchangeDeclare *exchangeDeclare) {
        exchangeDeclare.autoDelete = autoDelete
    }
}
func WithExchangeDeclareInternal(internal bool) ExchangeDeclareOption {
    return func(exchangeDeclare *exchangeDeclare) {
        exchangeDeclare.internal = internal
    }
}
func WithExchangeDeclareNoWait(noWait bool) ExchangeDeclareOption {
    return func(exchangeDeclare *exchangeDeclare) {
        exchangeDeclare.noWait = noWait
    }
}
func WithExchangeDeclareArguments(arguments map[string]interface{}) ExchangeDeclareOption {
    return func(exchangeDeclare *exchangeDeclare) {
        exchangeDeclare.arguments = arguments
    }
}

// --- WithExchangeDeclare ExchangeDeclareOption -end

// --- WithQueueDeclare QueueDeclareOption
type QueueDeclareOption func(queueDeclare *queueDeclare)
// WithQueueDeclare
func WithQueueDeclareDurable(durable bool) QueueDeclareOption {
    return func(queueDeclare *queueDeclare) {
        queueDeclare.durable = durable
        return
    }
}
// WithQueueDeclare
func WithQueueDeclareExclusive(exclusive bool) QueueDeclareOption {
    return func(queueDeclare *queueDeclare) {
        queueDeclare.exclusive = exclusive
    }
}
// WithQueueDeclare
func WithQueueDeclareAutoDelete(autoDelete bool) QueueDeclareOption {
    return func(queueDeclare *queueDeclare) {
        queueDeclare.autoDelete = autoDelete
        return
    }
}
// WithQueueDeclare
func WithQueueDeclareNoWait(noWait bool) QueueDeclareOption {
    return func(queueDeclare *queueDeclare) {
        queueDeclare.noWait = noWait
        return
    }
}
// WithQueueDeclare
func WithQueueDeclareArguments(opts ...QueueDeclareArgumentsOpt) QueueDeclareOption {
    return func(queueDeclare *queueDeclare) {
        arguments := make(map[string]interface{})
        for _, opt := range opts {
            opt(arguments)
        }
        queueDeclare.arguments = arguments
        return
    }
}

// WithQueueDeclare WithQueueDeclareArguments
// map/slice 为引用类型，无需指针
type QueueDeclareArguments map[string]interface{}

type QueueDeclareArgumentsOpt func(args QueueDeclareArguments)

func WithQueueDeclareArgumentsXDeadLetterExchange(exchangeName string) QueueDeclareArgumentsOpt {
    return func(args QueueDeclareArguments) {
        args[XDeadLetterExchange] = exchangeName
    }
}

func WithQueueDeclareArgumentsXDeadLetterRoutingKey(routingKey string) QueueDeclareArgumentsOpt {
    return func(args QueueDeclareArguments) {
        args[XDeadLetterRoutingKey] = routingKey
    }
}

func WithQueueDeclareArgumentsXMessageTTL(t time.Duration) QueueDeclareArgumentsOpt {
    return func(args QueueDeclareArguments) {
        if t > 0 {
            args[XMessageTTL] = int(t / 1e6)
        }
    }
}

// --- WithQueueDeclare QueueDeclareOption -end
