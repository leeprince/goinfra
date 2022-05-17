package rabbitmq

import "github.com/streadway/amqp"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/5/14 上午11:21
 * @Desc:   发布消息的直接应用
 */

func (cli *RabbitMQClient) PublishOne(body []byte) (err error) {
    err = cli.Publish(
        WithPublishProperties(body),
    )
    return
}

func (cli *RabbitMQClient) PublishTwo(body []byte) (err error) {
    err = cli.Publish(
        WithPublishProperties(
            body,
            // 消息持久化
            WithPropertiesDeliveryMode(amqp.Persistent),
        ),
    )
    return
}

