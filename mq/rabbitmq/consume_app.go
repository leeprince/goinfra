package rabbitmq

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/5/15 下午1:56
 * @Desc:   消费消息的直接应用
 */

// 简单队列（一次向一个消费者发送消息）
func (cli *RabbitMQClient) ConsumeSimple(handle ConsumeHandle, opts ...ConsumeParamOpt) (err error) {
    err = cli.Consume(
        handle,
        opts...,
    )
    
    return
}

// 工作队列（竞争消费者模式）
func (cli *RabbitMQClient) ConsumeWork(handle ConsumeHandle, opts ...ConsumeParamOpt) (err error) {
    err = cli.Consume(
        handle,
        opts...,
    )
    
    return
}

// 发布和订阅队列（一次向多个消费者发送消息）
func (cli *RabbitMQClient) ConsumeFanout(handle ConsumeHandle, opts ...ConsumeParamOpt) (err error) {
    opts = append(
        opts,
        WithConsumeParamOptAutoAck(true),
    )
    err = cli.Consume(
        handle,
        opts...,
    )
    
    return
}

// 路由队列（有选择地接收消息）
func (cli *RabbitMQClient) ConsumeDirect(handle ConsumeHandle, opts ...ConsumeParamOpt) (err error) {
    opts = append(
        opts,
        WithConsumeParamOptAutoAck(true),
    )
    err = cli.Consume(
        handle,
        opts...,
    )
    
    return
}

// 主题队列（基于模式（主题）接收消息 ）
func (cli *RabbitMQClient) ConsumeTopic(handle ConsumeHandle, opts ...ConsumeParamOpt) (err error) {
    opts = append(
        opts,
        WithConsumeParamOptAutoAck(true),
    )
    err = cli.Consume(
        handle,
        opts...,
    )
    
    return
}
