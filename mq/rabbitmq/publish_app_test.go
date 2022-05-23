package rabbitmq

import (
    "fmt"
    "github.com/spf13/cast"
    "github.com/streadway/amqp"
    "sync"
    "testing"
    "time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/5/15 下午2:00
 * @Desc:
 */

func TestRabbitMQClient_PublishSimple(t *testing.T) {
    type args struct {
        body []byte
    }
    
    tests := []struct {
        name    string
        args    args
        wantErr bool
    }{
        {
            args: args{
                body: []byte(fmt.Sprintf("PublishSimple-%s", cast.ToString(time.Now().UnixNano()/1e3))),
            },
        },
        {
            args: args{
                body: []byte(fmt.Sprintf("PublishSimple-%s", cast.ToString(time.Now().UnixNano()/1e3))),
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            cli, err := NewRabbitMQClient(
                WithUrl(myUrl),
                WithQueueDeclare(queueNameSimple),
            )
            if err != nil {
                fmt.Println("NewRabbitMQClient err:", err)
                return
            }
            fmt.Println("Boby:", string(tt.args.body))
            if err := cli.PublishSimple(tt.args.body); err != nil {
                t.Errorf("PublishSimple() error = %v", err)
            }
            fmt.Println("test:PublishSimple:msg:", string(tt.args.body))
        })
    }
}

func TestRabbitMQClient_PublishWork(t *testing.T) {
    type args struct {
        body []byte
    }
    
    tests := []struct {
        name    string
        args    args
        wantErr bool
    }{
        {
            args: args{
                body: []byte(fmt.Sprintf("PublishWork-%s", cast.ToString(time.Now().UnixNano()))),
            },
        },
        {
            args: args{
                body: []byte(fmt.Sprintf("PublishWork-%s", cast.ToString(time.Now().UnixNano()))),
            },
        },
        {
            args: args{
                body: []byte(fmt.Sprintf("PublishWork-%s", cast.ToString(time.Now().UnixNano()))),
            },
        },
        {
            args: args{
                body: []byte(fmt.Sprintf("PublishWork-%s", cast.ToString(time.Now().UnixNano()))),
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            cli, err := NewRabbitMQClient(
                WithUrl(myUrl),
                WithQueueDeclare(
                    queueNameWork,
                    WithQueueDeclareDurable(true),
                ),
            )
            if err != nil {
                fmt.Println("NewRabbitMQClient err:", err)
                return
            }
            fmt.Println("Boby:", string(tt.args.body))
            if err := cli.PublishWork(tt.args.body); (err != nil) != tt.wantErr {
                t.Errorf("PublishWork() error = %v, wantErr %v", err, tt.wantErr)
            }
            fmt.Println("test:PublishWork:msg:", string(tt.args.body))
        })
    }
}

func TestRabbitMQClient_PublishFanout(t *testing.T) {
    type fields struct {
        conf     *rabbitMQConf
        conn     *amqp.Connection
        connChan *amqp.Channel
        queue    amqp.Queue
    }
    type args struct {
        body []byte
    }
    tests := []struct {
        name    string
        fields  fields
        args    args
        wantErr bool
    }{
        {
            args: args{
                body: []byte(fmt.Sprintf("TestRabbitMQClient_PublishFanout-%s", cast.ToString(time.Now().UnixNano()))),
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            cli, err := NewRabbitMQClient(
                WithUrl(myUrl),
                WithExchangeDeclare(
                    exchangeNameFanout,
                    exchangeTypeFanout,
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
            fmt.Println("Boby:", string(tt.args.body))
            if err := cli.PublishFanout(tt.args.body); (err != nil) != tt.wantErr {
                t.Errorf("PublishFanout() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}

func TestRabbitMQClient_PublishFanout01(t *testing.T) {
    type fields struct {
        conf     *rabbitMQConf
        conn     *amqp.Connection
        connChan *amqp.Channel
        queue    amqp.Queue
    }
    type args struct {
        body []byte
    }
    tests := []struct {
        name    string
        fields  fields
        args    args
        wantErr bool
    }{
        {
            args: args{
                body: []byte(fmt.Sprintf("TestRabbitMQClient_PublishFanout-%s", cast.ToString(time.Now().UnixNano()))),
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            cli, err := NewRabbitMQClient(
                WithUrl(myUrl),
                WithExchangeDeclare(
                    exchangeNameFanout,
                    exchangeTypeFanout,
                    WithExchangeDeclareDurable(false),
                ),
                WithQueueDeclare(
                    "prince.queueName.Exclusive.tmp", // 声明队列为独占后，断开连接则会自动删除，所以声明队列名是没有意义的
                    WithQueueDeclareDurable(true),
                    WithQueueDeclareExclusive(true),
                    WithQueueDeclareAutoDelete(false),
                ),
            )
            if err != nil {
                fmt.Println("NewRabbitMQClient err:", err)
                return
            }
            fmt.Println("Boby:", string(tt.args.body))
            if err := cli.PublishFanout(tt.args.body); (err != nil) != tt.wantErr {
                t.Errorf("PublishFanout() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}

func TestRabbitMQClient_PublishDirect(t *testing.T) {
    type fields struct {
        conf     *rabbitMQConf
        conn     *amqp.Connection
        connChan *amqp.Channel
        queue    amqp.Queue
        mt       sync.Mutex
    }
    type args struct {
        body []byte
    }
    tests := []struct {
        name    string
        fields  fields
        args    args
        wantErr bool
    }{
        {
            args:args{
                body: []byte(fmt.Sprintf("TestRabbitMQClient_PublishDirect-%s", cast.ToString(time.Now().UnixNano()))),
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            cli, err := NewRabbitMQClient(
                WithUrl(myUrl),
                WithExchangeDeclare(
                    exchangeNameDirect,
                    exchangeTypeDirect,
                    WithExchangeDeclareDurable(true),
                ),
                WithQueueDeclare(
                    queueNameDirect,
                ),
                WithRoutingKey(routingKeyDirect),
            )
            if err != nil {
                fmt.Println("NewRabbitMQClient err:", err)
                return
            }
            fmt.Println("Boby:", string(tt.args.body))
            if err := cli.PublishDirect(tt.args.body); (err != nil) != tt.wantErr {
                t.Errorf("PublishDirect() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}

func TestRabbitMQClient_PublishDirect01(t *testing.T) {
    type fields struct {
        conf     *rabbitMQConf
        conn     *amqp.Connection
        connChan *amqp.Channel
        queue    amqp.Queue
        mt       sync.Mutex
    }
    type args struct {
        body []byte
    }
    tests := []struct {
        name    string
        fields  fields
        args    args
        wantErr bool
    }{
        {
            args:args{
                body: []byte(fmt.Sprintf("TestRabbitMQClient_PublishDirect-%s", cast.ToString(time.Now().UnixNano()))),
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            cli, err := NewRabbitMQClient(
                WithUrl(myUrl),
                WithExchangeDeclare(
                    exchangeNameDirect,
                    exchangeTypeDirect,
                    WithExchangeDeclareDurable(true),
                ),
                WithQueueDeclare(
                    queueNameDirect01,
                    WithQueueDeclareDurable(false),
                ),
                WithRoutingKey(routingKeyDirect),
            )
            if err != nil {
                fmt.Println("NewRabbitMQClient err:", err)
                return
            }
            fmt.Println("Boby:", string(tt.args.body))
            if err := cli.PublishDirect(
                tt.args.body,
                WithPropertiesDeliveryMode(PropertiesDeliveryModePersistent),
            ); (err != nil) != tt.wantErr {
                t.Errorf("PublishDirect() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}

func TestRabbitMQClient_PublishTopic(t *testing.T) {
    type fields struct {
        conf     *rabbitMQConf
        conn     *amqp.Connection
        connChan *amqp.Channel
        queue    amqp.Queue
        mt       sync.Mutex
    }
    type args struct {
        body []byte
        topic string
    }
    tests := []struct {
        name    string
        fields  fields
        args    args
        wantErr bool
    }{
        {
            args:args{
                body: []byte(fmt.Sprintf("TestRabbitMQClient_PublishTopic-%s-%s", routingKeyTopic, cast.ToString(time.Now().UnixNano()))),
                topic: routingKeyTopic,
            },
        },
        {
            args:args{
                body: []byte(fmt.Sprintf("TestRabbitMQClient_PublishTopic-%s-%s", routingKeyTopic01, cast.ToString(time.Now().UnixNano()))),
                topic: routingKeyTopic01,
            },
        },
        {
            args:args{
                body: []byte(fmt.Sprintf("TestRabbitMQClient_PublishTopic-%s-%s", routingKeyTopic02, cast.ToString(time.Now().UnixNano()))),
                topic: routingKeyTopic02,
            },
        },
        {
            args:args{
                body: []byte(fmt.Sprintf("TestRabbitMQClient_PublishTopic-%s-%s", routingKeyTopic03, cast.ToString(time.Now().UnixNano()))),
                topic: routingKeyTopic03,
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            cli, err := NewRabbitMQClient(
                WithUrl(myUrl),
                WithExchangeDeclare(
                    exchangeNameTopic,
                    exchangeTypeTopic,
                    WithExchangeDeclareDurable(true),
                ),
                WithQueueDeclare(
                    queueNameTopic,
                    WithQueueDeclareDurable(false),
                ),
                WithRoutingKey(tt.args.topic),
            )
            if err != nil {
                fmt.Println("NewRabbitMQClient err:", err)
                return
            }
            fmt.Println("Boby:", string(tt.args.body))
            if err := cli.PublishTopic(
                tt.args.body,
            ); (err != nil) != tt.wantErr {
                t.Errorf("PublishDirect() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}
