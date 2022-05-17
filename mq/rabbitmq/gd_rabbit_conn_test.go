package rabbitmq

import (
    "fmt"
    "github.com/leeprince/goinfra/plog"
    "github.com/spf13/cast"
    "github.com/streadway/amqp"
    "testing"
    "time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   1022/5/12 下午10:41
 * @Desc:
 */


const (
    // url            = "amqp://rabbit:aTjHMj7opZ3d5Kw6@10.21.32.14:5672/"
    url            = "amqp://guest:guest@127.0.0.1:5672/"
    exchangeName   = "prince.ex"
    queueName      = "prince.queueName"
    delayQueueName = "prince.queueName.delay"
    routeKey       = "prince.key"
    queueType      = "direct"
    delayInterval  = 300
    prefech        = 10
)

// const (
//     url            = "amqp://rabbit:aTjHMj7opZ3d5Kw6@10.21.32.14:5672/"
//     exchangeName   = "open_bank.ex.query_merchant_register"
//     queueName      = "open_bank.queueName.query_merchant_register"
//     delayQueueName = "open_bank.delay_queue.query_merchant_register"
//     routeKey       = "open_bank.key.query_merchant_register"
//     queueType      = "direct"
//     delayInterval  = 300
//     prefech        = 10
// )


func TestRabbitMqConn_Publish(t *testing.T) {
    this, err := NewRabbitMqConn(
        url,
        exchangeName,
        queueType,
        WithDelayLetter(delayQueueName))
    if err != nil {
        plog.Panic("NewRabbitMqConn err", err)
    }
    
    ctime := cast.ToString(time.Now().UnixNano() / 1e6)
    msg := "prince:msg:TestRabbitMqConn_Publish:" + ctime
    err = this.Publish(routeKey, msg)
    if err != nil {
        plog.Error("PublishDelay err", err)
        return
    }
    plog.Info("Publish succusss!")
}

func TestRabbitMqConn_PublishDelay(t *testing.T) {
    this, err := NewRabbitMqConn(
        url,
        exchangeName,
        queueType,
        WithDelayLetter(delayQueueName))
    if err != nil {
        plog.Panic("NewRabbitMqConn err", err)
    }
    
    ctime := cast.ToString(time.Now().UnixNano() / 1e6)
    msg := "prince:msg:TestRabbitMqConn_PublishDelay:" + ctime
    // seconds := int64(0) // 投递到延迟队列。直接过期进入指定（或默认）死信交换机，并由指定（或默认）死信路由键进行转发到响应的消息队列
    seconds := int64(10) // 投递到延迟队列。等待过期进入指定（或默认）死信交换机，并由指定（或默认）死信路由键进行转发到响应的消息队列
    err = this.PublishDelay(msg, seconds)
    if err != nil {
        plog.Error("PublishDelay err", err)
        return
    }
    plog.Info("PublishDelay succusss!")
}

func TestRabbitMqConn_Consume(t *testing.T) {
    this, err := NewRabbitMqConn(
      url,
      exchangeName,
      queueType)
    if err != nil {
        plog.Panic("NewRabbitMqConn err", err)
    }
    
    handle := func(msg *amqp.Delivery) {
        fmt.Printf("msg:%+v \n", msg)
        fmt.Printf("msg.Headers::%+v \n", msg.Headers)
        fmt.Printf("msg.Headers::%#v \n", msg.Headers)
        fmt.Println("string(msg.Body):", string(msg.Body))
        
        // 长时间堵塞等待再回复情况
        //  - rabbitmq 消息一只是未回复状态
        //  - 等待过程中，服务器宕机或者重启，消息会重新进入 ready 状态（等待消费）
        // time.Sleep(time.Minute * 2)
        
        // --- 手动回复
        // msg.Reject(false)
        // msg.Reject(true)
        // ---0
        // msg.Ack(false)
        // msg.Ack(true)
        // ---1
        msg.Nack(false, false)
        // msg.Nack(false, true)
        // msg.Nack(true, true)
        // msg.Nack(true, false)
        // --- 手动回复
        
        fmt.Println("msg --- end")
    }
    this.Consume(handle, queueName, prefech, routeKey)
}

func TestRabbitMqConn_Consume_Delay(t *testing.T) {
    this, err := NewRabbitMqConn(
      url,
      exchangeName,
      queueType,
      WithDelayLetter(delayQueueName))
    if err != nil {
        plog.Panic("NewRabbitMqConn err", err)
    }
    
    handle := func(msg *amqp.Delivery) {
        fmt.Printf("msg:%+v \n", msg)
        fmt.Printf("msg.Headers::%+v \n", msg.Headers)
        fmt.Printf("msg.Headers::%#v \n", msg.Headers)
        fmt.Println("string(msg.Body):", string(msg.Body))
        
        // 长时间堵塞等待再回复情况
        //  - 堵塞过程 rabbitmq 消息一只是未回复状态
        //  - 等待过程中，服务器宕机或者重启，消息会重新进入 ready 状态（等待消费）
        //      - 消息会重新进入 ready 状态（等待消费）:内部机制？跟连接有关？// TODO: 多个连接测试 - prince@todo 2022/5/13 下午6:18
        // time.Sleep(time.Minute * 2)
        
        // 达到定时效果：方案一检查是否延迟队列设置的过期时间，堵塞等待这个过期时间过期后重新回到队列实现定时任务效果。
        //  - 检查流程复杂，建议读取应用程序中配置的过期时间
        //  - 堵塞过程 rabbitmq 消息一只是未回复状态
        //  - 等待过程中，服务器宕机或者重启，消息会重新进入 ready 状态（等待消费）
        //  - 风险：
        //      - 消息会重新进入 ready 状态（等待消费）:内部机制？跟连接有关？// TODO: 多个连接测试 - prince@todo 2022/5/13 下午6:18
        //      - 等待过程中，服务器宕机或者重启，消息会重新进入 ready 状态（等待消费），会导致定时时间不可控
        //  - 建议：使用方案二
        /*if xDeathValue, ok := msg.Headers[xDeath]; ok {
            switch xDeathValue.(type) {
            case []interface{}:
                fmt.Println("xDeathValue.(type) == []interface{}")
                xDeathValueSliceInterface := xDeathValue.([]interface{})
                for _, xDeathValueInterfaceValue := range xDeathValueSliceInterface {
                    switch xDeathValueInterfaceValue.(type) {
                    case amqp.Table:
                        fmt.Println("xDeathValueInterfaceValue.(type) == amqp.Table")
                        amqpTable := xDeathValueInterfaceValue.(amqp.Table)
                        if t, ok := amqpTable[originalExpiration]; ok {
                            fmt.Println("amqpTable[originalExpiration]; ok")
        
                            mTimeMillisecondInt64 := cast.ToInt64(t)
                            fmt.Println("mTimeMillisecondInt64:", mTimeMillisecondInt64)
                            mTimeDuration := cast.ToDuration(mTimeMillisecondInt64 * 1e6)
                            fmt.Println("mTimeDuration:", mTimeDuration)
        
                            fmt.Println("time.Sleep")
                            time.Sleep(mTimeDuration)
                        }
                    }
                }
            default:
                fmt.Println("msg.Headers[xDeath] default:")
            }
        }
        time.Sleep(time.Second * 10)*/
        
        // 达到定时效果：方案二重新投递到延迟队列
        go func() {
            seconds := int64(20)
            msg := string(msg.Body)
            pushDelayErr := this.PublishDelay(msg, seconds)
            if pushDelayErr != nil {
                plog.Error("PublishDelay pushDelayErr != nil", pushDelayErr)
                return
            }
        }()
        time.Sleep(time.Second * 5)
        
        // --- 手动回复
        // msg.Reject(false)
        // msg.Reject(true)
        // ---0
        // msg.Ack(false)
        // msg.Ack(true)
        // ---1
        msg.Nack(false, false)
        // msg.Nack(false, true)
        // msg.Nack(true, true)
        // msg.Nack(true, false)
        // --- 手动回复
        
        fmt.Println("msg --- end")
    }
    this.Consume(handle, queueName, prefech, routeKey)
}
