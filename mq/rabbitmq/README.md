# RabbitMQ 消息队列

> 参考：https://www.rabbitmq.com/getstarted.html

## 消息队列类型 // TODO:  - prince@todo 2022/5/18 上午1:23
### 简单队列

阅读文档：
https://www.rabbitmq.com/tutorials/tutorial-one-go.html

测试文件及方法：
publish_app_test.go@TestRabbitMQClient_PublishOne,
consume_app_test.go@TestRabbitMQClient_ConsumeOne


### 工作队列（竞争消费者模式）

阅读文档：
https://www.rabbitmq.com/tutorials/tutorial-two-go.html

测试文件及方法：
publish_app_test.go@TestRabbitMQClient_PublishTwo,
consume_app_test.go@TestRabbitMQClient_ConsumeTwo @TestRabbitMQClient_ConsumeTwo01

### 发布和订阅（一次向多个消费者发送消息）
// TODO:  - prince@todo 2022/5/18 上午1:24

### 路由队列（有选择地接收消息 ）
// TODO:  - prince@todo 2022/5/18 上午1:24

### 主题队列（基于模式（主题）接收消息 ）
// TODO:  - prince@todo 2022/5/18 上午1:24

### RPC 队列（请求/回复模式）
// TODO:  - prince@todo 2022/5/18 上午1:24

## 关于生产者与消费者
- 非自动确认（autoAck=false）的消息
    - 消费者确认消息告诉 RabbitMQ 已正确接收 RabbitMQ 会安全的把消息从队列上删除。
    - 拒绝接收:可以使用 reject 命令拒绝接收消息，参数为 true，会发送给下一个消费者，false 时，会把消息从队列中移除不会分发给下一个消费者
        - reject(true):重新回到队列，会发送给下一个消费者
        - reject(false):会把消息从队列中移除（或者进入死信队列 // TODO:  - prince@todo 2022/5/16 下午11:13）不会分发给下一个消费者

- 确保即使消费者死亡， 任务没有丢失
    - 如果消费者收到一条消息，在确认之前，如果消费者死亡（其通道关闭，连接关闭，或 TCP 连接丢失）不发送 ack，RabbitMQ 将 了解消息未完全处理并将重新排队。 如果同时有其他消费者在线，它会快速重新发送 给另一个消费者。 这样您就可以确保不会丢失任何消息， 即使工人偶尔死亡。
    - 对消费者交付确认强制执行超时（默认为 30 分钟）。 这有助于检测从不确认交付的错误（卡住）消费者。 您可以按照中所述增加此交货确认超时时间。
- 虽确保即使消费者死亡， 任务没有丢失。但是 RabbitMQ 服务器停止（奔溃、重启、意外退出），任务仍然不会丢失。如何确保 RabbitMQ 服务器停止时，确保消息不会丢失，需要做两件事
    - 持久化持久化：声明队列持久化（durable=true）。确保RabbitMQ 服务器停止，队列不会丢失。
    - 消息持久化：设置 amqp.Publishing 的 DeliveryMode=amqp.Persistent。确保RabbitMQ 服务器停止，消息不会丢失。
    > 注意：RabbitMQ 不允许您重新定义现有队列具有不同的参数，并将向任何程序返回错误，阻止重新定义现有队列具有不同的参数。

- Oos 公平调度机制
    - channel.Qos 控制服务器在接收传递确认之前，将尝试在网络上为消费者保留多少消息或多少字节（换句话说：在消费者未确认前为消费者最多保留多少条消息或多少字节）。Qos的目的是确保服务器和客户端之间的网络缓冲区保持完整。

## 使用说明
- 发布者或者手动声明的队列统一以消费者去定义队列声明为准！
> 注意：RabbitMQ 不允许您重新定义现有队列具有不同的参数，并将向任何程序返回错误，阻止重新定义现有队列具有不同的参数。

## 部署
https://hub.docker.com/_/rabbitmq

## 端口说明

- 4369：epmd，RabbitMQ节点和CLI工具使用的对等发现服务
- 5672、5671：由不带TLS和带TLS的AMQP 0-9-1和1.0客户端使用
- 25672：用于节点间和CLI工具通信（Erlang分发服务器端口），并从动态范围分配（默认情况下限制为单个端口，计算为AMQP端口+ 20000）。除非确实需要这些端口上的外部连接（例如，群集使用联合身份验证或在子网外部的计算机上使用CLI工具），否则这些端口不应公开。有关详细信息，请参见网络指南。
- 35672-35682：由CLI工具（Erlang分发客户端端口）用于与节点进行通信，并从动态范围（计算为服务器分发端口+ 10000通过服务器分发端口+ 10010）分配。有关详细信息，请参见网络指南。
- 15672：HTTP API客户端，管理UI和Rabbitmqadmin （仅在启用了管理插件的情况下，默认启动）
- 61613、61614：不带TLS和带TLS的STOMP客户端（仅在启用STOMP插件的情况下）
- 1883、8883 ：（不带和带有TLS的MQTT客户端，如果启用了MQTT插件
- 15674：STOMP-over-WebSockets客户端（仅在启用了Web STOMP插件的情况下）
- 15675：MQTT-over-WebSockets客户端（仅当启用了Web MQTT插件时）
- 15692：Prometheus指标（仅在启用Prometheus插件的情况下）


## 基于消息队列类型实现延迟消息队列
// TODO:  - prince@todo 2022/5/18 上午1:23