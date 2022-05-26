package rabbitmq

import (
    "github.com/leeprince/goinfra/consts"
    "github.com/leeprince/goinfra/utils"
    "github.com/streadway/amqp"
    "strconv"
    "time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/5/14 上午10:59
 * @Desc:   发布消息
 */

// --- publishParamOpt
type publishParamOpt func(publishParams *publishParams)

func WithPublishParamMandatory(immediate bool) publishParamOpt {
    return func(publishParams *publishParams) {
        publishParams.immediate = immediate
    }
}
func WithPublishParamImmediate(immediate bool) publishParamOpt {
    return func(publishParams *publishParams) {
        publishParams.immediate = immediate
    }
}

// 设置为0会立即重新投递到死信队列
func WithPublishParamExpiration(t time.Duration) publishParamOpt {
    return func(publishParams *publishParams) {
        if t > 0 {
            publishParams.expiration = strconv.Itoa(int(t / 1e6))
        }
    }
}
func WithPublishParamProperties(body []byte, opts ...propertiesOpt) publishParamOpt {
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

// --- WithPublishParamProperties propertiesOpt
type propertiesOpt func(properties *properties)

// WithPublishParamProperties
func WithPropertiesContentType(contentType string) propertiesOpt {
    return func(properties *properties) {
        properties.contentType = contentType
    }
}

// WithPublishParamProperties
func WithPropertiesHeaders(headers map[string]interface{}) propertiesOpt {
    return func(properties *properties) {
        properties.headers = headers
    }
}

// WithPublishParamProperties
// 消息临时化：amqp.Transient=1;消息持久化:amqp.Persistent=2
func WithPropertiesDeliveryMode(deliveryMode uint8) propertiesOpt {
    return func(properties *properties) {
        if utils.InUint8(deliveryMode, []uint8{PropertiesDeliveryModeTransient, PropertiesDeliveryModePersistent}) {
            deliveryMode = PropertiesDeliveryModePersistent
        }
        properties.deliveryMode = deliveryMode
    }
}

// WithPublishParamProperties
func WithPropertiesPriority(priority uint8) propertiesOpt {
    return func(properties *properties) {
        properties.priority = priority
    }
}

// --- WithPublishParamProperties propertiesOpt -end
