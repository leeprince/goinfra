package rabbitmq

import (
    "errors"
    "fmt"
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
    autoAck     bool
    exclusive   bool
    noWait      bool
    arguments   map[string]interface{}
}

type ConsumeHandle func(data amqp.Delivery)

// 发生错误时，自动重试
func (cli *RabbitMQClient) Consume(handle ConsumeHandle, opts ...consumeParamOpt) (err error) {
    params := &consumeParam{
        queueName:   "",
        consumerTag: "",
        noLocal:     false,
        autoAck:     false,
        exclusive:   false,
        noWait:      false,
        arguments:   nil,
    }
    if cli.conf.queueDeclare.queueName != "" {
        params.queueName = cli.conf.queueDeclare.queueName
    }
    for _, opt := range opts {
        opt(params)
    }
    
    if params.queueName == "" {
        err = errors.New("params.queueName is empty")
        return
    }
    
    for {
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
                
                // 判断是否发现错误
                if data.Acknowledger == nil {
                    err = cli.Consume(handle, opts...)
                    failOnError(err, "<-deliver. cli.Consume(handle, opts...)")
                    return
                }
                go handle(data)
                
            }
        }
        
    }
}
