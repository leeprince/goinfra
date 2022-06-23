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

func TestRabbitMQClient_ConsumeSimple(t *testing.T) {
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
        WithQueueDeclare(queueNameSimple),
    )
    if err != nil {
        fmt.Println("NewRabbitMQClient err:", err)
        return
    }
    
    err = cli.ConsumeSimple(handle)
    if err != nil {
        t.Errorf("ConsumeSimple() error = %v", err)
    }
}

func TestRabbitMQClient_ConsumeSimple01(t *testing.T) {
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
        WithQueueDeclare(queueNameSimple),
    )
    if err != nil {
        fmt.Println("NewRabbitMQClient err:", err)
        return
    }
    
    err = cli.ConsumeSimple(handle)
    if err != nil {
        t.Errorf("ConsumeSimple() error = %v", err)
    }
}

func TestRabbitMQClient_ConsumeWork(t *testing.T) {
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
            queueNameWork,
            WithQueueDeclareDurable(true),
        ),
        WithQos(2, 0, false),
    )
    if err != nil {
        fmt.Println("NewRabbitMQClient err:", err)
        return
    }
    
    err = cli.ConsumeWork(handle)
    if err != nil {
        t.Errorf("ConsumeWork() error = %v", err)
    }
}

func TestRabbitMQClient_ConsumeWork01(t *testing.T) {
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
            queueNameWork,
            WithQueueDeclareDurable(true),
        ),
        WithQos(1, 0, false),
    )
    if err != nil {
        fmt.Println("NewRabbitMQClient err:", err)
        return
    }
    
    err = cli.ConsumeWork(
        handle,
    )
    if err != nil {
        t.Errorf("ConsumeWork() error = %v", err)
    }
}

func TestRabbitMQClient_ConsumeWork02(t *testing.T) {
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
            queueNameWork,
            WithQueueDeclareDurable(true),
        ),
        WithQos(1, 0, false),
    )
    if err != nil {
        fmt.Println("NewRabbitMQClient err:", err)
        return
    }
    
    err = cli.ConsumeWork(
        handle,
        WithConsumeParamOptExclusive(true),
    )
    if err != nil {
        t.Errorf("ConsumeWork() error = %v", err)
    }
}

func TestRabbitMQClient_ConsumeFanout(t *testing.T) {
    handle := func(msg amqp.Delivery) {
        fmt.Printf("msg:%+v \n", msg)
        fmt.Printf("msg.Headers::%+v \n", msg.Headers)
        fmt.Printf("msg.Headers::%#v \n", msg.Headers)
        fmt.Println("string(msg.Body):", string(msg.Body))
        
        time.Sleep(time.Second * 2)
        
        // 设置自动确认后，不可以手动确认，否则报错
        // --- 手动回复
        // msg.Reject(false)
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
        WithExchangeDeclare(
            exchangeNameFanout,
            ExchangeTypeFanout,
            WithExchangeDeclareDurable(false),
        ),
        WithQueueDeclare(
            "",
            WithQueueDeclareDurable(false),
            WithQueueDeclareExclusive(true),
            WithQueueDeclareAutoDelete(true),
        ),
    )
    if err != nil {
        fmt.Println("NewRabbitMQClient err:", err)
        return
    }
    
    err = cli.ConsumeFanout(handle)
    if err != nil {
        t.Errorf("ConsumeFanout() error = %v", err)
    }
}

func TestRabbitMQClient_ConsumeFanout01(t *testing.T) {
    handle := func(msg amqp.Delivery) {
        fmt.Printf("msg:%+v \n", msg)
        fmt.Printf("msg.Headers::%+v \n", msg.Headers)
        fmt.Printf("msg.Headers::%#v \n", msg.Headers)
        fmt.Println("string(msg.Body):", string(msg.Body))
        
        time.Sleep(time.Second * 2)
        
        // 设置自动确认后，不可以手动确认，否则报错
        // --- 手动回复
        // msg.Reject(false)
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
        WithExchangeDeclare(
            exchangeNameFanout,
            ExchangeTypeFanout,
            WithExchangeDeclareDurable(false),
        ),
        WithQueueDeclare(
            "",
            WithQueueDeclareDurable(false),
            WithQueueDeclareExclusive(true),
            WithQueueDeclareAutoDelete(true),
        ),
    )
    if err != nil {
        fmt.Println("NewRabbitMQClient err:", err)
        return
    }
    
    err = cli.ConsumeFanout(handle)
    if err != nil {
        t.Errorf("ConsumeFanout() error = %v", err)
    }
}

func TestRabbitMQClient_ConsumeDirect(t *testing.T) {
    handle := func(msg amqp.Delivery) {
        fmt.Printf("msg:%+v \n", msg)
        fmt.Printf("msg.Headers::%+v \n", msg.Headers)
        fmt.Printf("msg.Headers::%#v \n", msg.Headers)
        fmt.Println("string(msg.Body):", string(msg.Body))
        
        time.Sleep(time.Second * 2)
        
        // 设置自动确认后，不可以手动确认，否则报错
        // --- 手动回复
        // msg.Reject(false)
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
        WithExchangeDeclare(
            exchangeNameDirect,
            ExchangeTypeDirect,
            WithExchangeDeclareDurable(true),
        ),
        WithQueueDeclare(
            "",
            WithQueueDeclareDurable(false),
            WithQueueDeclareExclusive(true),
            WithQueueDeclareAutoDelete(false),
        ),
        WithRoutingKey(routingKeyDirect),
    )
    if err != nil {
        fmt.Println("NewRabbitMQClient err:", err)
        return
    }
    
    err = cli.ConsumeDirect(handle)
    if err != nil {
        t.Errorf("ConsumeDirect() error = %v", err)
    }
}

func TestRabbitMQClient_ConsumeDirect01(t *testing.T) {
    handle := func(msg amqp.Delivery) {
        fmt.Printf("msg:%+v \n", msg)
        fmt.Printf("msg.Headers::%+v \n", msg.Headers)
        fmt.Printf("msg.Headers::%#v \n", msg.Headers)
        fmt.Println("string(msg.Body):", string(msg.Body))
        
        time.Sleep(time.Second * 2)
        
        // 设置自动确认后，不可以手动确认，否则报错
        // --- 手动回复
        // msg.Reject(false)
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
        WithExchangeDeclare(
            exchangeNameDirect,
            ExchangeTypeDirect,
            WithExchangeDeclareDurable(true),
        ),
        WithQueueDeclare(
            "",
            WithQueueDeclareDurable(false),
            WithQueueDeclareExclusive(true),
            WithQueueDeclareAutoDelete(false),
        ),
        WithRoutingKey(routingKeyDirect),
    )
    if err != nil {
        fmt.Println("NewRabbitMQClient err:", err)
        return
    }
    
    err = cli.ConsumeDirect(handle)
    if err != nil {
        t.Errorf("ConsumeDirect() error = %v", err)
    }
}

func TestRabbitMQClient_ConsumeTopic(t *testing.T) {
    handle := func(msg amqp.Delivery) {
        fmt.Printf("msg:%+v \n", msg)
        fmt.Printf("msg.Headers::%+v \n", msg.Headers)
        fmt.Printf("msg.Headers::%#v \n", msg.Headers)
        fmt.Println("string(msg.Body):", string(msg.Body))
        
        time.Sleep(time.Second * 2)
        
        // 设置自动确认后，不可以手动确认，否则报错
        // --- 手动回复
        // msg.Reject(false)
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
    
    topic := "prince.RoutingKey.*"
    cli, err := NewRabbitMQClient(
        WithUrl(myUrl),
        WithExchangeDeclare(
            exchangeNameTopic,
            ExchangeTypeTopic,
            WithExchangeDeclareDurable(true),
        ),
        WithQueueDeclare(
            "",
            WithQueueDeclareDurable(false),
            WithQueueDeclareExclusive(true),
            WithQueueDeclareAutoDelete(false),
        ),
        WithRoutingKey(topic),
    )
    if err != nil {
        fmt.Println("NewRabbitMQClient err:", err)
        return
    }
    
    fmt.Println("topic::::::::::::::", topic)
    err = cli.ConsumeTopic(
        handle,
        WithConsumeParamOptExclusive(true),
    )
    if err != nil {
        t.Errorf("ConsumeTopic() error = %v", err)
    }
}

func TestRabbitMQClient_ConsumeTopic00(t *testing.T) {
    handle := func(msg amqp.Delivery) {
        fmt.Printf("msg:%+v \n", msg)
        fmt.Printf("msg.Headers::%+v \n", msg.Headers)
        fmt.Printf("msg.Headers::%#v \n", msg.Headers)
        fmt.Println("string(msg.Body):", string(msg.Body))
        
        time.Sleep(time.Second * 2)
        
        // 设置自动确认后，不可以手动确认，否则报错
        // --- 手动回复
        // msg.Reject(false)
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
    
    topic := "prince.RoutingKey.*"
    cli, err := NewRabbitMQClient(
        WithUrl(myUrl),
        WithExchangeDeclare(
            exchangeNameTopic,
            ExchangeTypeTopic,
            WithExchangeDeclareDurable(true),
        ),
        WithQueueDeclare(
            "",
            WithQueueDeclareDurable(false),
            WithQueueDeclareExclusive(true),
            WithQueueDeclareAutoDelete(false),
        ),
        WithRoutingKey(topic),
    )
    if err != nil {
        fmt.Println("NewRabbitMQClient err:", err)
        return
    }
    
    fmt.Println("topic::::::::::::::", topic)
    err = cli.ConsumeTopic(
        handle,
        WithConsumeParamOptExclusive(true),
    )
    if err != nil {
        t.Errorf("ConsumeTopic() error = %v", err)
    }
}

func TestRabbitMQClient_ConsumeTopic01(t *testing.T) {
    handle := func(msg amqp.Delivery) {
        fmt.Printf("msg:%+v \n", msg)
        fmt.Printf("msg.Headers::%+v \n", msg.Headers)
        fmt.Printf("msg.Headers::%#v \n", msg.Headers)
        fmt.Println("string(msg.Body):", string(msg.Body))
        
        time.Sleep(time.Second * 2)
        
        // 设置自动确认后，不可以手动确认，否则报错
        // --- 手动回复
        // msg.Reject(false)
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
    
    topic := "prince.*.topic"
    cli, err := NewRabbitMQClient(
        WithUrl(myUrl),
        WithExchangeDeclare(
            exchangeNameTopic,
            ExchangeTypeTopic,
            WithExchangeDeclareDurable(true),
        ),
        WithQueueDeclare(
            "",
            WithQueueDeclareDurable(false),
            WithQueueDeclareExclusive(true),
            WithQueueDeclareAutoDelete(false),
        ),
        WithRoutingKey(topic),
    )
    if err != nil {
        fmt.Println("NewRabbitMQClient err:", err)
        return
    }
    
    fmt.Println("topic::::::::::::::", topic)
    err = cli.ConsumeTopic(handle)
    if err != nil {
        t.Errorf("ConsumeTopic() error = %v", err)
    }
}

func TestRabbitMQClient_ConsumeTopic02(t *testing.T) {
    handle := func(msg amqp.Delivery) {
        fmt.Printf("msg:%+v \n", msg)
        fmt.Printf("msg.Headers::%+v \n", msg.Headers)
        fmt.Printf("msg.Headers::%#v \n", msg.Headers)
        fmt.Println("string(msg.Body):", string(msg.Body))
        
        time.Sleep(time.Second * 2)
        
        // 设置自动确认后，不可以手动确认，否则报错
        // --- 手动回复
        // msg.Reject(false)
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
    
    topic := "prince.#"
    cli, err := NewRabbitMQClient(
        WithUrl(myUrl),
        WithExchangeDeclare(
            exchangeNameTopic,
            ExchangeTypeTopic,
            WithExchangeDeclareDurable(true),
        ),
        WithQueueDeclare(
            "",
            WithQueueDeclareDurable(false),
            WithQueueDeclareExclusive(true),
            WithQueueDeclareAutoDelete(false),
        ),
        WithRoutingKey(topic),
    )
    if err != nil {
        fmt.Println("NewRabbitMQClient err:", err)
        return
    }
    
    fmt.Println("topic::::::::::::::", topic)
    err = cli.ConsumeTopic(handle)
    if err != nil {
        t.Errorf("ConsumeTopic() error = %v", err)
    }
}

func TestRabbitMQClient_ConsumeDeadLettered(t *testing.T) {
    handle := func(msg amqp.Delivery) {
        fmt.Printf("msg:%+v \n", msg)
        fmt.Printf("msg.Headers::%+v \n", msg.Headers)
        fmt.Printf("msg.Headers::%#v \n", msg.Headers)
        fmt.Println("string(msg.Body):", string(msg.Body))
        
        time.Sleep(time.Second * 2)
        
        // 设置自动确认后，不可以手动确认，否则报错
        // --- 手动回复
        // msg.Reject(false)
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
    
    topic := "prince.#"
    cli, err := NewRabbitMQClient(
        WithUrl(myUrl),
        WithExchangeDeclare(
            exchangeNameTopic,
            ExchangeTypeTopic,
            WithExchangeDeclareDurable(true),
        ),
        WithQueueDeclare(
            "",
            WithQueueDeclareDurable(false),
            WithQueueDeclareExclusive(true),
            WithQueueDeclareAutoDelete(false),
        ),
        WithRoutingKey(topic),
    )
    if err != nil {
        fmt.Println("NewRabbitMQClient err:", err)
        return
    }
    
    fmt.Println("topic::::::::::::::", topic)
    err = cli.ConsumeTopic(handle)
    if err != nil {
        t.Errorf("ConsumeTopic() error = %v", err)
    }
}
