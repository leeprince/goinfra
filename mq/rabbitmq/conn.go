package rabbitmq

import (
    "github.com/leeprince/goinfra/plog"
    "github.com/pkg/errors"
    "github.com/streadway/amqp"
    "sync"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/5/12 上午10:42
 * @Desc:   初始化客户端
 */

type RabbitMQClient struct {
    conf     *rabbitMQConf
    conn     *amqp.Connection
    connChan *amqp.Channel
    queue    amqp.Queue
    mt       sync.Mutex
}

func NewRabbitMQClient(opts ...ConfOption) (cli *RabbitMQClient, err error) {
    cli = new(RabbitMQClient)
    cli.mt = sync.Mutex{}
    
    if err = cli.initConf(opts...); err != nil {
        return
    }
    if err = cli.initConn(); err != nil {
        return
    }
    if err = cli.initChannel(); err != nil {
        return
    }
    if err = cli.initExchangeDeclare(); err != nil {
        return
    }
    if err = cli.initQueueDeclare(); err != nil {
        return
    }
    if err = cli.initQueueBind(); err != nil {
        return
    }
    if err = cli.initQos(); err != nil {
        return
    }
    
    return
}

// 初始化配置
func (cli *RabbitMQClient) initConf(opts ...ConfOption) (err error) {
    conf := &rabbitMQConf{
        url:                   defaultURL,
        vhost:                 defaultVhost,
        cancelexchangeDeclare: false,
        exchangeDeclare:       nil,
        cancelQueueDeclare:    false,
        queueDeclare:          nil,
        routingKey:            "",
        qos:                   qos{},
        errRetryTime:          defaultErrRetryTime,
    }
    for _, opt := range opts {
        err = opt(conf)
        if err != nil {
            err = errors.Wrap(err, "initConf opt(conf) err != nil")
            return
        }
    }
    
    cli.conf = conf
    
    return
}

// 初始化连接
func (cli *RabbitMQClient) initConn() error {
    amqpConfig := amqp.Config{
        // 不为空时，会覆盖 url 解析出来的 vhost
        Vhost: cli.conf.vhost,
    }
    conn, dialErr := amqp.DialConfig(cli.conf.url, amqpConfig)
    if dialErr != nil {
        return dialErr
    }
    
    cli.conn = conn
    
    return nil
}

// 初始化连接通道
func (cli *RabbitMQClient) initChannel() error {
    channel, err := cli.conn.Channel()
    if err != nil {
        return err
    }
    
    cli.connChan = channel
    
    return nil
}

// 初始化声明交换机
func (cli *RabbitMQClient) initExchangeDeclare() error {
    if cli.conf.cancelexchangeDeclare {
        plog.Info("initExchangeDeclare cli.conf.cancelexchangeDeclare")
        return nil
    }
    
    exchange := cli.conf.exchangeDeclare
    if exchange == nil {
        plog.Info("initExchangeDeclare exchange.exchangeName == ''")
        return nil
    }
    
    err := cli.connChan.ExchangeDeclare(
        exchange.exchangeName,
        exchange.exchangeType,
        exchange.durable,
        exchange.autoDelete,
        exchange.internal,
        exchange.noWait,
        exchange.arguments,
    )
    if err != nil {
        return err
    }
    
    return nil
}

// 初始化声明队列
func (cli *RabbitMQClient) initQueueDeclare() error {
    if cli.conf.cancelQueueDeclare {
        plog.Info("initQueueDeclare cli.conf.cancelQueueDeclare")
        return nil
    }
    
    queueDeclare := cli.conf.queueDeclare
    queue, err := cli.connChan.QueueDeclare(
        queueDeclare.queueName,
        queueDeclare.durable,
        queueDeclare.autoDelete,
        queueDeclare.exclusive,
        queueDeclare.noWait,
        queueDeclare.arguments,
    )
    if err != nil {
        return err
    }
    
    cli.queue = queue
    
    return nil
}

// 初始化交换机与队列绑定
func (cli *RabbitMQClient) initQueueBind() error {
    if cli.conf.exchangeDeclare == nil || cli.conf.exchangeDeclare.exchangeName == "" {
        plog.Infof("initQueueBind cli.conf.exchangeDeclare == nil || cli.conf.exchangeDeclare.exchangeName == ''. cli.conf:%+v", cli.conf)
        return nil
    }
    err := cli.connChan.QueueBind(
        cli.queue.Name,
        cli.conf.routingKey,
        cli.conf.exchangeDeclare.exchangeName,
        cli.conf.queueDeclare.noWait,
        cli.conf.exchangeDeclare.arguments,
    )
    if err != nil {
        return err
    }
    
    return nil
}

// 初始化公平调度机制
func (cli *RabbitMQClient) initQos() error {
    return cli.connChan.Qos(
        cli.conf.qos.prefetchCount,
        cli.conf.qos.prefetchSize,
        cli.conf.qos.global,
    )
}

// 尝试重新建立 RabbitMQ 客户端
func (cli *RabbitMQClient) retryNewRabbitMQClient() (err error) {
    // 添加互斥锁
    cli.mt.Lock()
    cli.mt.Unlock()
    
    if cli.conn.IsClosed() {
        if err = cli.initConn(); err != nil {
            err = errors.Wrap(err, "retryNewRabbitMQClient cli.initConn err")
            return
        }
    }
    if err = cli.initChannel(); err != nil {
        err = errors.Wrap(err, "retryNewRabbitMQClient cli.initChannel err")
        return
    }
    
    if err = cli.initExchangeDeclare(); err != nil {
        err = errors.Wrap(err, "retryNewRabbitMQClient cli.initExchangeDeclare err")
        return
    }
    
    if err = cli.initQueueDeclare(); err != nil {
        err = errors.Wrap(err, "retryNewRabbitMQClient cli.initQueueDeclare err")
        return
    }
    
    if err = cli.initQueueBind(); err != nil {
        err = errors.Wrap(err, "retryNewRabbitMQClient cli.initQueueBind err")
        return
    }
    
    if err = cli.initQos(); err != nil {
        err = errors.Wrap(err, "retryNewRabbitMQClient cli.initQos err")
        return
    }
    
    return
}
