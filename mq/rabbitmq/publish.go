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
    exchange   string
    routingKey string
    mandatory  bool
    immediate  bool
    properties *properties
}

type properties struct {
    contentType  string                 // MIME content type
    headers      map[string]interface{} // Application or header exchange table。// amqp.Table - prince@comm 2022/5/14 上午11:09
    deliveryMode uint8                  // queueName implementation use - Transient (1) or Persistent (2)
    priority     uint8                  // queueName implementation use - 0 to 9
    
    body []byte
}

func (cli *RabbitMQClient) Publish(opts ...publishParamOpt) (err error) {
    params := &publishParams{
        exchange:   "",
        routingKey: "",
        mandatory:  false,
        immediate:  false,
        properties: &properties{},
    }
    if cli.conf.queueDeclare.queueName != "" {
        params.routingKey = cli.conf.queueDeclare.queueName
    }
    for _, opt := range opts {
        opt(params)
    }
    
    if params.routingKey == "" {
        err = errors.New("params.routingKey is empty")
        return
    }
    
    if string(params.properties.body) == "" {
        err = errors.New("string(params.properties.body) is empty")
        return
    }
    
    err = cli.connChan.Publish(
        params.exchange,
        params.routingKey,
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
