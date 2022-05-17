package rabbitmq

import (
    "github.com/leeprince/goinfra/consts"
    "github.com/leeprince/goinfra/utils"
    "github.com/streadway/amqp"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/5/14 上午10:59
 * @Desc:   发布消息
 */

// --- publishParamOpt
type publishParamOpt func(publishParams *publishParams)

func WithPublishExchange(exchange string) publishParamOpt {
    return func(publishParams *publishParams) {
        publishParams.exchange = exchange
    }
}
func WithPublishRouteKey(routingKey string) publishParamOpt {
    return func(publishParams *publishParams) {
        publishParams.routingKey = routingKey
    }
}
func WithPublishProperties(body []byte, opts ...propertiesOpt) publishParamOpt {
    return func(publishParams *publishParams) {
        properties := &properties{
            contentType:  consts.ContextTypeTextPlain,
            headers:      nil,
            deliveryMode: amqp.Persistent, // 消息临时化：amqp.Transient;消息持久化:amqp.Persistent
            priority:     0,
            body:         body,
        }
    
        for _, opt := range opts {
            opt(properties)
        }
        
        publishParams.properties = properties
    }
}
// --- publishParamOpt -end

// --- WithPublishProperties propertiesOpt
type propertiesOpt func(properties *properties)

// WithPublishProperties
func WithPropertiesContentType(contentType string) propertiesOpt {
    return func(properties *properties) {
        properties.contentType = contentType
    }
}

// WithPublishProperties
func WithPropertiesHeaders(headers map[string]interface{}) propertiesOpt {
    return func(properties *properties) {
        properties.headers = headers
    }
}

// WithPublishProperties
// 消息临时化：amqp.Transient=1;消息持久化:amqp.Persistent=2
func WithPropertiesDeliveryMode(deliveryMode uint8) propertiesOpt {
    return func(properties *properties) {
        if utils.InUint8(deliveryMode, []uint8{amqp.Transient, amqp.Persistent}) {
            deliveryMode = amqp.Persistent
        }
        properties.deliveryMode = deliveryMode
    }
}

// WithPublishProperties
func WithPropertiesPriority(priority uint8) propertiesOpt {
    return func(properties *properties) {
        properties.priority = priority
    }
}

// --- WithPublishProperties propertiesOpt -end

