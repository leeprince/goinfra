package rabbitmq

import (
    "errors"
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

func WithQueueDeclare(queueName string, opts ...queueDeclareOption) confOption {
    return func(conf *rabbitMQConf) (err error) {
        if queueName == "" {
            err = errors.New("WithQueueDeclare name is empty")
            return
        }
        
        queueDeclare := &queueDeclare{
            queueName:  queueName,
            durable:    true, // 队列是否持久化
            exclusive:  false,
            autoDelete: false,
            noWait:     false,
        }
    
        for _, opt := range opts {
            opt(queueDeclare)
        }
        
        conf.queueDeclare = queueDeclare
        return
    }
}

// --- confOption -end

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
// --- WithQueueDeclare queueDeclareOption -end
