package rabbitmq

import (
    "errors"
    "fmt"
    "github.com/leeprince/goinfra/utils/jsonpb"
    "log"
    "strconv"
    "sync"
    "time"
    
    "github.com/streadway/amqp"
)

const Route_ALL string = "#"

var marshaler = &jsonpb.JSONPb{OrigName: true, EmitDefaults: true}

type rabbitConfig struct {
    uri          string
    exchangeName string
    queueType    string // topic direct
    delayQueue   string
    internal     bool
}

type RabbitMqConn struct {
    conn           *amqp.Connection
    conf           *rabbitConfig
    connErrCh      chan *amqp.Error
    chErrCh        chan *amqp.Error
    publishCh      *amqp.Channel
    publishDelayCh *amqp.Channel
    mt             sync.Mutex
}

type MsgHandleFunc func(msg *amqp.Delivery)
type RabbitMqOption func(ms *RabbitMqConn)

func NewRabbitMqConn(rabbitUri string, exchangeName string, queueType string, opts ...RabbitMqOption) (*RabbitMqConn, error) {
    var cli *RabbitMqConn
    fmt.Println("NewRabbitMqConn rabbitUri", rabbitUri)
    conn, err := amqp.Dial(rabbitUri)
    if err != nil {
        log.Fatalf("Failed to connect to RabbitMQ")
        return cli, err
    }
    
    return getRabbitMqClient(conn, rabbitUri, exchangeName, queueType, opts...)
}

func getRabbitMqClient(conn *amqp.Connection, rabbitUri string, exchangeName string, queueType string, opts ...RabbitMqOption) (*RabbitMqConn, error) {
    var cli *RabbitMqConn
    fmt.Println("rabbitUri", rabbitUri)
    
    rabbitConf := &rabbitConfig{
        uri:          rabbitUri,
        exchangeName: exchangeName,
        queueType:    queueType,
    }
    
    publishCh, err := conn.Channel()
    if err != nil {
        panicOnError(err, "fail to alloc channel")
    }
    publishDelayCh, err := conn.Channel()
    if err != nil {
        panicOnError(err, "fail to alloc delay channel")
    }
    
    cli = &RabbitMqConn{
        conn:           conn,
        conf:           rabbitConf,
        connErrCh:      make(chan *amqp.Error),
        publishCh:      publishCh,
        publishDelayCh: publishDelayCh,
        mt:             sync.Mutex{},
    }
    
    for _, opt := range opts {
        opt(cli)
    }
    
    return cli, nil
}

func WithDelayLetter(delayQueue string) RabbitMqOption {
    return func(rc *RabbitMqConn) {
        rc.conf.delayQueue = delayQueue
    }
}

func WithExchangeInternal() RabbitMqOption {
    return func(rc *RabbitMqConn) {
        rc.conf.internal = true
    }
}

// 消费普通队列
func (this *RabbitMqConn) Consume(handle MsgHandleFunc, queueName string, prefech int, RoutingKeys ...string) {
fallback:
    for {
        if this.conn.IsClosed() {
            this.reconnect()
        }
        
        ch, err := this.conn.Channel()
        if err != nil {
            failOnError(err, "Failed to open a channel")
            continue
        }
        
        err = ch.ExchangeDeclare(
            this.conf.exchangeName, // 声明交换机名称。没有会自动创建
            this.conf.queueType,    // type
            true,                   // durable
            false,                  // auto-deleted
            this.conf.internal,     // internal
            false,                  // no-wait
            nil,                    // arguments
        )
        if err != nil {
            failOnError(err, "Failed to declare an exchangeName")
            continue
        }
        
        queue, err := ch.QueueDeclare(
            queueName, // 声明队列名称。没有会自动创建
            true,      // durable
            false,     // auto-deleted
            false,     // exclusive
            false,     // no-wait
            nil,       // arguments
        
        )
        if err != nil {
            failOnError(err, "Failed to declare a queueName")
            continue
        }
        
        // dead letter queueName
        if this.conf.delayQueue != "" {
            _, err = ch.QueueDeclare(
                this.conf.delayQueue, // 声明延迟队列名称。没有会自动创建
                true,                 // durable
                false,                // delete when unused
                false,                // exclusive
                false,                // no-wait
                amqp.Table{
                    // 当消息过期时把消息发送到死信交换机。默认把当前交换机当作死信交换机
                    "x-dead-letter-exchangeName": this.conf.exchangeName,
                    // 当消息过期时把消息发送到死信路由键。默认把 RoutingKeys[0]当作死信路由键
                    "x-dead-letter-routing-key": RoutingKeys[0],
                }, // arguments
            )
            if err != nil {
                failOnError(err, "Failed to declare delay queueName")
                continue
            }
        }
        
        if prefech == 0 {
            prefech = 1
        }
        err = ch.Qos(
            prefech, // prefetch count
            0,       // prefetch size
            false,   // global
        )
        if err != nil {
            failOnError(err, "Failed to set prefetch count")
            continue
        }
        
        for _, route := range RoutingKeys {
            // 将交换机通过路由键绑定指定队列名称
            err = ch.QueueBind(
                queue.Name,             // queueName name
                route,                  // routing key
                this.conf.exchangeName, // exchangeName
                false,
                nil)
            if err != nil {
                failOnError(err, "Failed to bind a queueName")
                goto fallback
            }
        }
        
        this.chErrCh = ch.NotifyClose(make(chan *amqp.Error))
        
        msgs, err := ch.Consume(
            queue.Name, // queueName
            "",         // consumerTag
            false,      // auto ack
            false,      // exclusive
            false,      // no local
            false,      // no wait
            nil,        // args
        )
        if err != nil {
            failOnError(err, "Failed to register a consumer")
            continue
        }
        
        for {
            select {
            case d, ok := <-msgs:
                if !ok {
                    goto fallback
                }
                go handle(&d)
            case <-this.connErrCh:
                goto fallback
            case <-this.chErrCh:
                goto fallback
            }
        }
    }
}

func (this *RabbitMqConn) Publish(RoutingKey string, msg interface{}) error {
    msgByte, err := marshaler.Marshal(msg)
    if err != nil {
        return err
    }
    
    if this.conn.IsClosed() {
        println("=== publish reconnect")
        this.reconnect()
    }
    
    err = this.publishCh.Publish(
        this.conf.exchangeName, // exchangeName
        RoutingKey,               // routing key
        false,                  // mandatory
        false,                  // immediate
        amqp.Publishing{
            ContentType:  "text/plain",
            DeliveryMode: amqp.Persistent,
            Timestamp:    time.Now(),
            Body:         msgByte,
        })
    
    failOnError(err, "Failed to publish a message")
    println("send", string(msgByte))
    
    return err
}

// 推送到延时队列 使用默认延时队列
func (this *RabbitMqConn) PublishDelay(msg interface{}, seconds int64) error {
    if this.conf.delayQueue == "" {
        return errors.New("this.conf.delayQueue must not empty")
    }
    
    msgByte, err := marshaler.Marshal(msg)
    if err != nil {
        return err
    }
    
    // 关于：Default exchangeName：The default exchangeName is implicitly bound to every queueName, with a routing key equal to the queueName name. It is not possible to explicitly bind to, or unbind from the default exchangeName. It also cannot be deleted.
    // 	- 默认交换机会绑定路由键等于队列名称的队列
    // 	- 默认路由键等于队列名称
    err = this.publishDelayCh.Publish(
        "",                   // 延迟队列投递到默认的交换机（Default exchangeName）。不应投递到延迟队列绑定到的死信交换机中，因为可能把非延迟的队列当死信队列了，导致无法实现延迟效果
        this.conf.delayQueue, // 默认路由键等于队列名称。routing key == this.conf.delayQueue
        false,                // mandatory
        false,                // immediate
        amqp.Publishing{
            ContentType:  "text/plain",
            DeliveryMode: amqp.Persistent,
            Timestamp:    time.Now(),
            Body:         msgByte,
            Expiration:   sec2MillStr(seconds),
        })
    
    failOnError(err, "Failed to publish a message")
    println("send", string(msgByte))
    
    return err
}

func sec2MillStr(seconds int64) string {
    if seconds <= 0 {
        return "0"
    }
    
    return strconv.FormatInt(seconds*1000, 10)
}

func (this *RabbitMqConn) reconnect() {
    this.mt.Lock()
    defer this.mt.Unlock()
    
    if this.conn.IsClosed() {
        this.conn = mustConnRabbit(this.conf.uri)
        this.connErrCh = this.conn.NotifyClose(make(chan *amqp.Error))
        this.publishCh, _ = this.conn.Channel()
        this.publishDelayCh, _ = this.conn.Channel()
    }
}

func mustConnRabbit(uri string) *amqp.Connection {
    for {
        conn, err := amqp.Dial(uri)
        if err == nil {
            return conn
        }
        
        log.Printf("trying to connect rabbitMq err:%v\n", err)
        time.Sleep(time.Second)
    }
}
