package rabbitmq

import (
    "errors"
    "fmt"
    "github.com/leeprince/goinfra/plog"
    "github.com/streadway/amqp"
    "time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/5/14 下午1:33
 * @Desc:   消费消息
 */

type consumeParam struct {
    queueName   string
    consumerTag string
    noLocal     bool
    // 是否自动回复
    autoAck   bool
    // 该消费者是否独占该队列
    // exclusive=true时，服务器将确保这是此队列中的唯一使用者。
    // exclusive=false时，服务器将在多个消费者之间公平地分发传递。
    exclusive bool
    noWait    bool
    arguments map[string]interface{}
}

type ConsumeHandle func(data amqp.Delivery)

// 发生错误时，自动重试
func (cli *RabbitMQClient) Consume(handle ConsumeHandle, opts ...ConsumeParamOpt) (err error) {
    params := &consumeParam{
        queueName:   "",
        consumerTag: "",
        noLocal:     false,
        autoAck:     false,
        exclusive:   false,
        noWait:      false,
        arguments:   nil,
    }

    for {
        // 考虑重试后对于RabbitMQ会自动创建一个随机命名的队列名的情况，所以将判断在循环中
        if cli.queue.Name != "" {
            params.queueName = cli.queue.Name
        }
        for _, opt := range opts {
            opt(params)
        }
        
        if params.queueName == "" {
            err = errors.New("params.queueName is empty")
            return
        }
        
        // 只读通道（channel）
        var delivery <-chan amqp.Delivery
        delivery, err = cli.connChan.Consume(
            params.queueName,
            params.consumerTag,
            params.autoAck,
            params.exclusive,
            params.noLocal,
            params.noWait,
            params.arguments,
        )
        if err != nil {
            failOnError(err, fmt.Sprintf("cli.connChan.Consume err. to sleep:%s", cli.conf.errRetryTime))
            
            // 发生错误时，重试的等待时间。避免重试失败无限重试
            time.Sleep(cli.conf.errRetryTime)
            
            // 尝试重新建立 RabbitMQ 客户端
            //  - 解决队列名在监听过程中被删除时，会自动创建队列名并恢复监听
            //  - 解决RabbitMQ服务器在监听过程中重启，重启后会自动创建队列名并恢复监听
            //  - 注意：
            //      - cli.connChan.Consume()在返回错误的同时也会往写空数据到通道中，所以需要判断<-delivery是否不可用需要重试的情况
            //      - 对于RabbitMQ会自动创建一个随机命名的队列名，需要使用声明队列后的amqp.Queue.Name 当作新的队列名
            err = cli.retryNewRabbitMQClient()
            failOnError(err, "cli.connChan.Consume err > cli.retryNewRabbitMQClient err")
            
            continue
        }
        
        // fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>cli.connChan.Consume", time.Now().UnixNano()/1e6)
        // 监听 <-delivery 通道
        //  - 监听过程中，如果发现错误（队列被删除），则进行重试
        for {
            select {
            case data := <-delivery:
                // fmt.Printf("------------------time:%d, data := <-delivery:%+v \n", time.Now().UnixNano()/1e6, data)
                
                // 判断 <-delivery 的数据是否可用
                if data.Acknowledger == nil && string(data.Body) == "" {
                    plog.Info("<-delivery: data.Acknowledger == nil && string(data.Body) == ''")
                    err = cli.Consume(handle, opts...)
                    failOnError(err, "<-delivery: data.Acknowledger == nil && string(data.Body) == '' > cli.Consume(handle, opts...) err")
                    return
                }
                go handle(data)
            }
        }
        
    }
}
