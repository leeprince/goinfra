package rabbitmq

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/5/14 上午11:21
 * @Desc:   对应交换机类型的发布消息封装方法
 */

// 简单队列（一次向一个消费者发送消息）
func (cli *RabbitMQClient) PublishSimple(body []byte, opts ...PropertiesOpt) (err error) {
	err = cli.Publish(
		WithPublishParamProperties(
			body,
			opts...,
		),
	)
	return
}

// 工作队列（竞争消费者模式）
func (cli *RabbitMQClient) PublishWork(body []byte, opts ...PropertiesOpt) (err error) {
	opts = append(
		opts,
		WithPropertiesDeliveryMode(PropertiesDeliveryModePersistent),
	)
	err = cli.Publish(
		WithPublishParamProperties(
			body,
			opts...,
		),
	)
	return
}

// 发布和订阅队列（一次向多个消费者发送消息）
func (cli *RabbitMQClient) PublishFanout(body []byte, opts ...PropertiesOpt) (err error) {
	opts = append(
		opts,
		WithPropertiesDeliveryMode(PropertiesDeliveryModeTransient),
	)
	err = cli.Publish(
		WithPublishParamProperties(
			body,
			opts...,
		),
	)
	return
}

// 路由队列（有选择地接收消息）
func (cli *RabbitMQClient) PublishDirect(body []byte, opts ...PropertiesOpt) (err error) {
	err = cli.Publish(
		WithPublishParamProperties(
			body,
			opts...,
		),
	)
	return
}

// 路由队列（有选择地接收消息）
func (cli *RabbitMQClient) PublishTopic(body []byte, opts ...PropertiesOpt) (err error) {
	err = cli.Publish(
		WithPublishParamProperties(
			body,
			opts...,
		),
	)
	return
}
