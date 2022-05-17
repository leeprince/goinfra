package rabbit

import (
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/streadway/amqp"
	"gitlab.yewifi.com/golden-cloud/common/gcjsonpb"
)

const Route_ALL string = "#"

var marshaler = &gcjsonpb.JSONPb{OrigName: true, EmitDefaults: true}

func failOnError(err error, msg string) {
	if err != nil {
		println(fmt.Sprintf("msg:%s, err:%s", msg, err.Error()))
		log.Println(fmt.Sprintf("%s: %v", msg, err))
	}
}

func panicOnError(err error, msg string) {
	if err != nil {
		println(fmt.Sprintf("msg:%s, err:%s", msg, err.Error()))
		log.Fatalf("%s: %v", msg, err)
	}
}

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

func NewRabbitMqConnByVhost(rabbitUri string, vhost string, exchangeName string, queueType string, opts ...RabbitMqOption) (*RabbitMqConn, error) {
	var cli *RabbitMqConn
	fmt.Println("NewRabbitMqConnByConfig rabbitUri", rabbitUri)
	config := amqp.Config{
		Vhost: vhost,
	}
	conn, err := amqp.DialConfig(rabbitUri, config)
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

//消费普通队列
func (this *RabbitMqConn) Consume(handle MsgHandleFunc, queueName string, prefech int, routeKeys ...string) {
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
			this.conf.exchangeName, // name
			this.conf.queueType,    // type
			true,                   // durable
			false,                  // auto-deleted
			this.conf.internal,     // internal
			false,                  // no-wait
			nil,                    // arguments
		)
		if err != nil {
			failOnError(err, "Failed to declare an exchange")
			continue
		}

		queue, err := ch.QueueDeclare(
			queueName, // name
			true,      // durable
			false,     // auto-deleted
			false,     // exclusive
			false,     // no-wait
			nil,       // arguments

		)
		if err != nil {
			failOnError(err, "Failed to declare a queue")
			continue
		}

		// dead letter queue
		if this.conf.delayQueue != "" {
			_, err = ch.QueueDeclare(
				this.conf.delayQueue, // name
				true,                 // durable
				false,                // delete when unused
				false,                // exclusive
				false,                // no-wait
				amqp.Table{
					// 当消息过期时把消息发送到exchange
					"x-dead-letter-exchange": this.conf.exchangeName,
				}, // arguments
			)
			if err != nil {
				failOnError(err, "Failed to declare delay queue")
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

		for _, route := range routeKeys {
			err = ch.QueueBind(
				queue.Name,             // queue name
				route,                  // routing key
				this.conf.exchangeName, // exchange
				false,
				nil)
			if err != nil {
				failOnError(err, "Failed to bind a queue")
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

func (this *RabbitMqConn) Publish(routeKey string, msg interface{}) error {
	msgByte, err := marshaler.Marshal(msg)
	if err != nil {
		return err
	}

	// TODO(corvinFn): seems useless if publish fail
	if this.conn.IsClosed() {
		println("=== publish reconnect")
		this.reconnect()
	}
	// ch, err := this.conn.Channel()
	// failOnError(err, "Failed to open a channel")

	err = this.publishCh.Publish(
		this.conf.exchangeName, // exchange
		routeKey,               // routing key
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

//推送到延时队列 使用默认延时队列
func (this *RabbitMqConn) PublishDelay(msg interface{}, seconds int64) error {
	msgByte, err := marshaler.Marshal(msg)
	if err != nil {
		return err
	}

	err = this.publishDelayCh.Publish(
		"",
		this.conf.delayQueue, // routing key
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

//消费延时队列
//注：需要先创建好延时队列（调用DeclareDelayQueue）
func (this *RabbitMqConn) ConsumeDelay(handle MsgHandleFunc, queueName string, prefech int, routeKeys ...string) {
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
			this.conf.exchangeName, // name
			this.conf.queueType,    // type
			true,                   // durable
			false,                  // auto-deleted
			this.conf.internal,     // internal
			false,                  // no-wait
			nil,                    // arguments
		)
		if err != nil {
			failOnError(err, "Failed to declare an exchange")
			continue
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

		for _, route := range routeKeys {
			err = ch.QueueBind(
				queueName,              // queue name
				route,                  // routing key
				this.conf.exchangeName, // exchange
				false,
				nil)
			if err != nil {
				failOnError(err, "Failed to bind a queue")
				goto fallback
			}
		}

		this.chErrCh = ch.NotifyClose(make(chan *amqp.Error))

		msgs, err := ch.Consume(
			queueName, // queueName
			"",        // consumerTag
			false,     // auto ack
			false,     // exclusive
			false,     // no local
			false,     // no wait
			nil,       // args
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

//创建延时队列
//delayQueueName：延时队列名称
//deadRouteKey：延时队列中消息超时后，投递到死信交换机时使用的路由键（跟Publish方法的routeKey类似）
func (this *RabbitMqConn) DeclareDelayQueue(delayQueueName string, deadRouteKey string) error {
	if this.conn.IsClosed() {
		this.reconnect()
	}

	ch, err := this.conn.Channel()
	if err != nil {
		failOnError(err, "Failed to open a channel")
		return err
	}

	err = ch.ExchangeDeclare(
		this.conf.exchangeName, // name
		this.conf.queueType,    // type
		true,                   // durable
		false,                  // auto-deleted
		this.conf.internal,     // internal
		false,                  // no-wait
		nil,                    // arguments
	)
	if err != nil {
		failOnError(err, "Failed to declare an exchange")
		return err
	}

	// dead letter queue
	_, err = ch.QueueDeclare(
		delayQueueName, // name
		true,           // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		amqp.Table{
			"x-dead-letter-exchange":    this.conf.exchangeName, //expired message wil send to this exchange
			"x-dead-letter-routing-key": deadRouteKey,           //dead message publish routeKey
		}, // arguments
	)
	if err != nil {
		failOnError(err, "Failed to declare delay queue")
		return err
	}

	return nil
}

//推送到延时队列 并指定队列名称
func (this *RabbitMqConn) PublishDelayWithRouteKey(delayQueueName string, msg interface{}, seconds int64) error {
	msgByte, err := marshaler.Marshal(msg)
	if err != nil {
		return err
	}

	err = this.publishDelayCh.Publish(
		"",             // 如果为空，则进入rabbitmq默认的直流交换机
		delayQueueName, // routing key，实际为延时队列名称
		false,          // mandatory
		false,          // immediate
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

//============================x-delayed-message==============================//
//消费延时队列 (满足不定长延时时间)
func (this *RabbitMqConn) ConsumeXDelayMsg(handle MsgHandleFunc, xExchange, queueName, deadRouteKey string, prefech int, routeKeys ...string) {
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

		//declare x-delayed exchange
		err = ch.ExchangeDeclare(
			xExchange,           // name
			"x-delayed-message", // a new exchange type
			true,                // durable
			false,               // auto-deleted
			false,               // internal
			false,               // no-wait
			amqp.Table{
				"x-delayed-type": this.conf.queueType,
			},
		)
		if err != nil {
			failOnError(err, fmt.Sprintf("Failed to declare x-delayed-exchange %s", xExchange))
			continue
		}

		// delay letter queue
		_, err = ch.QueueDeclare(
			queueName, // name
			true,      // durable
			false,     // delete when unused
			false,     // exclusive
			false,     // no-wait
			amqp.Table{
				"x-dead-letter-exchange":    this.conf.exchangeName, //expired message wil send to this exchange
				"x-dead-letter-routing-key": deadRouteKey,           //dead message publish routeKey
			}, // arguments
		)
		if err != nil {
			failOnError(err, "Failed to declare delay queue")
			continue
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

		for _, route := range routeKeys {
			err = ch.QueueBind(
				queueName, // queue name
				route,     // routing key
				xExchange, // exchange
				false,
				nil)
			if err != nil {
				failOnError(err, "Failed to bind a queue")
				goto fallback
			}
		}

		this.chErrCh = ch.NotifyClose(make(chan *amqp.Error))

		msgs, err := ch.Consume(
			queueName, // queueName
			"",        // consumerTag
			false,     // auto ack
			false,     // exclusive
			false,     // no local
			false,     // no wait
			nil,       // args
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

//推送到延时交换机
func (this *RabbitMqConn) PublishXDelayMsg(exchange, routingKey string, msg interface{}, delay int64) error {
	msgByte, err := marshaler.Marshal(msg)
	if err != nil {
		return err
	}

	headers := make(map[string]interface{})
	if delay != 0 {
		headers["x-delay"] = delay
	}

	err = this.publishDelayCh.Publish(
		exchange,   // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType:  "text/plain",
			DeliveryMode: amqp.Persistent,
			Timestamp:    time.Now(),
			Body:         msgByte,
			Headers:      headers,
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
