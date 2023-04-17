package rabbitmq

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/5/15 下午2:15
 * @Desc:
 */

const (
	myUrl = "amqp://guest:guest@127.0.0.1:5672/"

	// 交换机名
	exchangeNameFanout            = "prince.exchangeName.fanout"
	exchangeNameDirect            = "prince.exchangeName.Direct"
	exchangeNameTopic             = "prince.exchangeName.topic"
	exchangeNameDeadLetteredTopic = "prince.dl.exchangeName.topic"

	// 队列名
	queueNameSimple            = "prince.queueName.simple"
	queueNameWork              = "prince.queueName.work"
	queueNameDirect            = "prince.queueName.Direct"
	queueNameDirect01          = "prince.queueName.Direct-01"
	queueNameTopic             = "prince.queueName.topic"
	queueNameDeadLetteredTopic = "prince.dl.queueName.topic"

	// 路由键RoutingKey
	routingKeyDirect  = "prince.RoutingKey.Direct"
	routingKeyTopic   = "prince.RoutingKey.topic"
	routingKeyTopic01 = "prince.RoutingKey.topic01"
	routingKeyTopic02 = "prince.RoutingKey01.topic"
	routingKeyTopic03 = "prince.RoutingKey02.topic02"
)
