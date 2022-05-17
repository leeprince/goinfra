package rabbitmq

import (
    "fmt"
    "github.com/streadway/amqp"
    "testing"
    "time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/5/15 下午2:15
 * @Desc:
 */

const (
    myUrl        = "amqp://guest:guest@127.0.0.1:5672/"
    queueNameOne = "prince.queueName.one"
    queueNameTwo = "prince.queueName.work"
)

func TestRabbitMQClient_ConsumeOne(t *testing.T) {
    handle := func(msg amqp.Delivery) {
        fmt.Printf("msg:%+v \n", msg)
        fmt.Printf("msg.Headers::%+v \n", msg.Headers)
        fmt.Printf("msg.Headers::%#v \n", msg.Headers)
        fmt.Println("string(msg.Body):", string(msg.Body))
        
        // --- 手动回复
        msg.Reject(false)
        // msg.Reject(true)
        // ---0
        // msg.Ack(false)
        // msg.Ack(true)
        // ---1
        // msg.Nack(false, false)
        // msg.Nack(false, true)
        // msg.Nack(true, true)
        // msg.Nack(true, false)
        // --- 手动回复
        
        fmt.Println("msg --- end")
    }
    
    cli, err := NewRabbitMQClient(
        WithUrl(myUrl),
        WithQueueDeclare(queueNameOne),
    )
    if err != nil {
        fmt.Println("NewRabbitMQClient err:", err)
    }
    
    err = cli.ConsumeOne(handle)
    if err != nil {
        t.Errorf("ConsumeOne() error = %v", err)
    }
}

func TestRabbitMQClient_ConsumeTwo(t *testing.T) {
    handle := func(msg amqp.Delivery) {
        fmt.Printf("msg:%+v \n", msg)
        fmt.Printf("msg.Headers::%+v \n", msg.Headers)
        fmt.Printf("msg.Headers::%#v \n", msg.Headers)
        fmt.Println("string(msg.Body):", string(msg.Body))
        
        // --- 手动回复
        msg.Reject(false)
        // msg.Reject(true)
        // ---0
        // msg.Ack(false)
        // msg.Ack(true)
        // ---1
        // msg.Nack(false, false)
        // msg.Nack(false, true)
        // msg.Nack(true, true)
        // msg.Nack(true, false)
        // --- 手动回复
        
        fmt.Println("msg --- end")
    }
    
    cli, err := NewRabbitMQClient(
        WithUrl(myUrl),
        WithQueueDeclare(
            queueNameTwo,
            // 定义队列持久化
            WithQueueDeclareDurable(true),
        ),
        WithQos(2, 0, false),
    )
    if err != nil {
        fmt.Println("NewRabbitMQClient err:", err)
    }
    
    err = cli.ConsumeTwo(handle)
    if err != nil {
        t.Errorf("ConsumeOne() error = %v", err)
    }
}

func TestRabbitMQClient_ConsumeTwo01(t *testing.T) {
    handle := func(msg amqp.Delivery) {
        fmt.Printf("msg:%+v \n", msg)
        fmt.Printf("msg.Headers::%+v \n", msg.Headers)
        fmt.Printf("msg.Headers::%#v \n", msg.Headers)
        fmt.Println("string(msg.Body):", string(msg.Body))
        
        time.Sleep(time.Second * 2)

        // --- 手动回复
        msg.Reject(false)
        // msg.Reject(true)
        // ---0
        // msg.Ack(false)
        // msg.Ack(true)
        // ---1
        // msg.Nack(false, false)
        // msg.Nack(false, true)
        // msg.Nack(true, true)
        // msg.Nack(true, false)
        // --- 手动回复
        
        fmt.Println("msg --- end")
    }
    
    cli, err := NewRabbitMQClient(
        WithUrl(myUrl),
        WithQueueDeclare(
            queueNameTwo,
            // 定义队列持久化
            WithQueueDeclareDurable(true),
        ),
        WithQos(1, 0, false),
    )
    if err != nil {
        fmt.Println("NewRabbitMQClient err:", err)
    }
    
    err = cli.ConsumeTwo(handle)
    if err != nil {
        t.Errorf("ConsumeOne() error = %v", err)
    }
}
