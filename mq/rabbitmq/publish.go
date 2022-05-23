package rabbitmq

import (
    "errors"
    "github.com/streadway/amqp"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/5/14 上午10:59
 * @Desc:   发布消息
 */

type publishParams struct {
    mandatory  bool
    immediate  bool
    properties *properties
}

type properties struct {
    // MIME
    contentType  string                 // MIME content type
    // 自定义头部信息
    headers      map[string]interface{} // Application or header exchangeName table。// amqp.Table - prince@comm 2022/5/14 上午11:09
    // 消息投递模式。消息临时化：amqp.Transient - 消息持久化:amqp.Persistent
    deliveryMode uint8                  // queueName implementation use - Transient (1) or Persistent (2)
    // 优先级
    priority     uint8                  // queueName implementation use - 0 to 9
    // 消息体
    body []byte
}

func (cli *RabbitMQClient) Publish(opts ...publishParamOpt) (err error) {
    params := &publishParams{
        mandatory:  false,
        immediate:  false,
        properties: &properties{},
    }

    for _, opt := range opts {
        opt(params)
    }
    
    if string(params.properties.body) == "" {
        err = errors.New("string(params.properties.body) is empty")
        return
    }
    
    var exchangeName string
    if cli.conf.exchangeDeclare != nil && cli.conf.exchangeDeclare.exchangeName != "" {
        exchangeName = cli.conf.exchangeDeclare.exchangeName
    }
    
    routingKey := cli.queue.Name
    if cli.conf.routingKey != "" {
        routingKey = cli.conf.routingKey
    }
    
    err = cli.connChan.Publish(
        exchangeName,
        routingKey,
        params.mandatory,
        params.immediate,
        amqp.Publishing{
            Headers:      params.properties.headers,
            ContentType:  params.properties.contentType,
            DeliveryMode: params.properties.deliveryMode,
            Priority:     params.properties.priority,
            Body:         params.properties.body,
        },
    )
    return
}
