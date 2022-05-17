package rabbitmq

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/5/14 下午1:37
 * @Desc:
 */

// --- consumeParamOpt
type consumeParamOpt func(consumeParam *consumeParam)

func WithConsumeParamOptQueueName(queueName string) consumeParamOpt {
    return func(consumeParam *consumeParam) {
        consumeParam.queueName = queueName
    }
}
func WithConsumeParamOptConsumerTag(consumerTag string) consumeParamOpt {
    return func(consumeParam *consumeParam) {
        consumeParam.consumerTag = consumerTag
    }
}
func WithConsumeParamOptNoLocal(noLocal bool) consumeParamOpt {
    return func(consumeParam *consumeParam) {
        consumeParam.noLocal = noLocal
    }
}
func WithConsumeParamOptAutoAck(autoAck bool) consumeParamOpt {
    return func(consumeParam *consumeParam) {
        consumeParam.autoAck = autoAck
    }
}
func WithConsumeParamOptExclusive(exclusive bool) consumeParamOpt {
    return func(consumeParam *consumeParam) {
        consumeParam.exclusive = exclusive
    }
}
func WithConsumeParamOptNoWait(noWait bool) consumeParamOpt {
    return func(consumeParam *consumeParam) {
        consumeParam.noWait = noWait
    }
}
func WithConsumeParamOptArguments(arguments map[string]interface{}) consumeParamOpt {
    return func(consumeParam *consumeParam) {
        consumeParam.arguments = arguments
    }
}

// --- consumeParamOpt -end
