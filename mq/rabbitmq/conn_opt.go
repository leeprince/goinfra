package rabbitmq

import "github.com/streadway/amqp"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/5/12 上午10:42
 * @Desc:   初始化
 */

type RabbitMQClient struct {
    conf     *rabbitMQConf
    conn     *amqp.Connection
    connChan *amqp.Channel
}

type rabbitMQConf struct {
    url   string
    vhost string
}

type confOption func(conf *rabbitMQConf)

func NewRabbitMQClient(opts ...confOption) (cli *RabbitMQClient, err error) {
    // TODO: 声明跟队列的字段 - prince@todo 2022/5/14 上午12:31
    cli = new(RabbitMQClient)
    
    cli.initConf(opts...)
    
    if err = cli.initConn(); err != nil {
        return
    }
    if err = cli.initChannel(); err != nil {
        return
    }
    
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

// 初始化配置
func (cli *RabbitMQClient) initConf(opts ...confOption) {
    conf := &rabbitMQConf{
        url:   defaultURL,
        vhost: defaultVhost,
    }
    for _, opt := range opts {
        opt(conf)
    }
    cli.conf = conf
}

func WithUrl(url string) confOption {
    return func(conf *rabbitMQConf) {
        conf.url = url
    }
}

func WithVhost(vhost string) confOption {
    return func(conf *rabbitMQConf) {
        conf.vhost = vhost
    }
}
