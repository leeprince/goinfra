package rabbitmq

import (
    "time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/5/12 上午10:42
 * @Desc:   初始化客户端的选项
 */

// --- confOption
type confOption func(conf *rabbitMQConf) (err error)

func WithUrl(url string) confOption {
    return func(conf *rabbitMQConf) (err error) {
        conf.url = url
        return
    }
}

func WithVhost(vhost string) confOption {
    return func(conf *rabbitMQConf) (err error) {
        conf.vhost = vhost
        return
    }
}

func WithErrRetryTime(t time.Duration) confOption {
    return func(conf *rabbitMQConf) (err error) {
        conf.errRetryTime = t
        return
    }
}

func WithCancelQueueDeclare(cancel bool) confOption {
    return func(conf *rabbitMQConf) (err error) {
        conf.cancelQueueDeclare = cancel
        return
    }
}

func WithQos(prefetchCount, prefetchSize int, global bool) confOption {
    return func(conf *rabbitMQConf) (err error) {
        conf.qos = qos{
            prefetchCount: prefetchCount,
            prefetchSize:  prefetchSize,
            global:        global,
        }
        return
    }
}

func WithExchangeDeclare(exchangeName, exchangeType string, opts ...exchangeDeclareOption) confOption {
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

func WithQueueDeclare(queueName string, opts ...queueDeclareOption) confOption {
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

func WithRoutingKey(routingKey string) confOption {
    return func(conf *rabbitMQConf) (err error) {
        conf.routingKey = routingKey
        return
    }
}

// --- confOption -end

// --- WithExchangeDeclare exchangeDeclareOption
type exchangeDeclareOption func(exchangeDeclare *exchangeDeclare)

func WithExchangeDeclarePassive(passive bool) exchangeDeclareOption {
    return func(exchangeDeclare *exchangeDeclare) {
        exchangeDeclare.passive = passive
    }
}
func WithExchangeDeclareDurable(durable bool) exchangeDeclareOption {
    return func(exchangeDeclare *exchangeDeclare) {
        exchangeDeclare.durable = durable
    }
}
func WithExchangeDeclareAutoDelete(autoDelete bool) exchangeDeclareOption {
    return func(exchangeDeclare *exchangeDeclare) {
        exchangeDeclare.autoDelete = autoDelete
    }
}
func WithExchangeDeclareInternal(internal bool) exchangeDeclareOption {
    return func(exchangeDeclare *exchangeDeclare) {
        exchangeDeclare.internal = internal
    }
}
func WithExchangeDeclareNoWait(noWait bool) exchangeDeclareOption {
    return func(exchangeDeclare *exchangeDeclare) {
        exchangeDeclare.noWait = noWait
    }
}
func WithExchangeDeclareArguments(arguments map[string]interface{}) exchangeDeclareOption {
    return func(exchangeDeclare *exchangeDeclare) {
        exchangeDeclare.arguments = arguments
    }
}

// --- WithExchangeDeclare exchangeDeclareOption -end

// --- WithQueueDeclare queueDeclareOption
type queueDeclareOption func(queueDeclare *queueDeclare)
// WithQueueDeclare
func WithQueueDeclareDurable(durable bool) queueDeclareOption {
    return func(queueDeclare *queueDeclare) {
        queueDeclare.durable = durable
        return
    }
}
// WithQueueDeclare
func WithQueueDeclareExclusive(exclusive bool) queueDeclareOption {
    return func(queueDeclare *queueDeclare) {
        queueDeclare.exclusive = exclusive
    }
}
// WithQueueDeclare
func WithQueueDeclareAutoDelete(autoDelete bool) queueDeclareOption {
    return func(queueDeclare *queueDeclare) {
        queueDeclare.autoDelete = autoDelete
        return
    }
}
// WithQueueDeclare
func WithQueueDeclareNoWait(noWait bool) queueDeclareOption {
    return func(queueDeclare *queueDeclare) {
        queueDeclare.noWait = noWait
        return
    }
}
// WithQueueDeclare
func WithQueueDeclareArguments(arguments map[string]interface{}) queueDeclareOption {
    return func(queueDeclare *queueDeclare) {
        queueDeclare.arguments = arguments
        return
    }
}
// --- WithQueueDeclare queueDeclareOption -end
