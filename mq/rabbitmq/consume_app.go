package rabbitmq

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/5/15 下午1:56
 * @Desc:   消费消息的直接应用
 */

func (cli *RabbitMQClient) ConsumeOne(handle ConsumeHandle) (err error) {
    err = cli.Consume(
        handle,
    )
    
    return
}

func (cli *RabbitMQClient) ConsumeTwo(handle ConsumeHandle) (err error) {
    err = cli.Consume(
        handle,
    )
    
    return
}
