package rabbitmq

import "time"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/5/14 上午10:59
 * @Desc:   发布消息
 */

type publishParams struct {
    exchange   string
    routingKey string
    mandatory  bool
    immediate  bool
    properties properties
    body       []byte
}

type publishParamsOpt func(publishParams *publishParams)

type properties struct {
    ContentType     string                 // MIME content type
    ContentEncoding string                 // MIME content encoding
    Headers         map[string]interface{} // Application or header exchange table。// amqp.Table - prince@comm 2022/5/14 上午11:09
    DeliveryMode    uint8                  // queue implementation use - Transient (1) or Persistent (2)
    Priority        uint8                  // queue implementation use - 0 to 9
    CorrelationId   string                 // application use - correlation identifier
    ReplyTo         string                 // application use - address to to reply to (ex: RPC)
    Expiration      string                 // implementation use - message expiration spec
    MessageId       string                 // application use - message identifier
    Timestamp       time.Time              // application use - message timestamp
    Type            string                 // application use - message type name
    UserId          string                 // application use - creating user id
    AppId           string                 // application use - creating application
    reserved1       string                 // was cluster-id - process for buffer consumption
}

func (cli *RabbitMQClient) Publish(opts ...publishParamsOpt) (err error) {
    cli.connChan.Publish()
}

func (cli *RabbitMQClient) PublishOne(queueName string) (err error) {
    cli.connChan.Publish()
}
