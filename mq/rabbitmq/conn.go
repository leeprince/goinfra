package rabbitmq

import (
    "github.com/leeprince/goinfra/plog"
    "github.com/pkg/errors"
    "github.com/streadway/amqp"
    "time"
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
}

type rabbitMQConf struct {
    url                string
    vhost              string
    cancelQueueDeclare bool // 强制取消队列声明
    queueDeclare       *queueDeclare
    qos                qos
    // 发生错误时，重试的等待时间
    errRetryTime time.Duration
}

type queueDeclare struct {
    queueName string
    
    // 队列是否持久化
    durable bool
    
    // 如果想拥有私人队列只为一个消费者服务，可以设置 exclusive 参数为 true
    // 如果需要临时队列和结合 exclusive 和 autoDelete，autoDelete 在消费者取消订阅时，会自动删除，都设置为 true
    exclusive bool
    
    // 是否自动删除
    autoDelete bool
    
    // 是否不等待
    noWait bool
}

type qos struct {
    prefetchCount int
    prefetchSize  int
    global        bool
}

func NewRabbitMQClient(opts ...confOption) (cli *RabbitMQClient, err error) {
    cli = new(RabbitMQClient)
    
    if err = cli.initConf(opts...); err != nil {
        return
    }
    if err = cli.initConn(); err != nil {
        return
    }
    if err = cli.initChannel(); err != nil {
        return
    }
    if err = cli.initQueueDeclare(); err != nil {
        return
    }
    if err = cli.initQos(); err != nil {
        return
    }
    
    return
}

// 初始化配置
func (cli *RabbitMQClient) initConf(opts ...confOption) (err error) {
    conf := &rabbitMQConf{
        url:                defaultURL,
        vhost:              defaultVhost,
        cancelQueueDeclare: false,
        queueDeclare:       &queueDeclare{},
        errRetryTime:       defaultErrRetryTime,
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

// 初始化队列声明
func (cli *RabbitMQClient) initQueueDeclare() error {
    if cli.conf.cancelQueueDeclare {
        plog.Info("initQueueDeclare cli.conf.cancelQueueDeclare")
        return nil
    }
    if cli.conf.queueDeclare.queueName == "" {
        err := errors.New("initQueueDeclare cli.conf.queueDeclare.queueName is empty")
        plog.Error(err)
        return err
    }
    queueDeclare := cli.conf.queueDeclare
    queue, err := cli.connChan.QueueDeclare(
        queueDeclare.queueName,
        queueDeclare.durable,
        queueDeclare.autoDelete,
        queueDeclare.exclusive,
        queueDeclare.noWait,
        nil,
    )
    if err != nil {
        return err
    }
    
    cli.queue = queue
    
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
    
    if err = cli.initQueueDeclare(); err != nil {
        err = errors.Wrap(err, "retryNewRabbitMQClient cli.initQueueDeclare err")
        return
    }
    
    return
}
