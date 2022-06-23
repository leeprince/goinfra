package rabbitmq

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/5/14 下午1:37
 * @Desc:
 */

// --- ConsumeParamOpt
type ConsumeParamOpt func(consumeParam *consumeParam)

func WithConsumeParamOptQueueName(queueName string) ConsumeParamOpt {
    return func(consumeParam *consumeParam) {
        consumeParam.queueName = queueName
    }
}
func WithConsumeParamOptConsumerTag(consumerTag string) ConsumeParamOpt {
    return func(consumeParam *consumeParam) {
        consumeParam.consumerTag = consumerTag
    }
}
func WithConsumeParamOptNoLocal(noLocal bool) ConsumeParamOpt {
    return func(consumeParam *consumeParam) {
        consumeParam.noLocal = noLocal
    }
}
func WithConsumeParamOptAutoAck(autoAck bool) ConsumeParamOpt {
    return func(consumeParam *consumeParam) {
        consumeParam.autoAck = autoAck
    }
}
func WithConsumeParamOptExclusive(exclusive bool) ConsumeParamOpt {
    return func(consumeParam *consumeParam) {
        consumeParam.exclusive = exclusive
    }
}
func WithConsumeParamOptNoWait(noWait bool) ConsumeParamOpt {
    return func(consumeParam *consumeParam) {
        consumeParam.noWait = noWait
    }
}
func WithConsumeParamOptArguments(arguments map[string]interface{}) ConsumeParamOpt {
    return func(consumeParam *consumeParam) {
        consumeParam.arguments = arguments
    }
}

// --- ConsumeParamOpt -end
