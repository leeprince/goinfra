package rabbitmq

import (
    "fmt"
    "github.com/spf13/cast"
    "testing"
    "time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/5/15 下午2:00
 * @Desc:
 */

func TestRabbitMQClient_PublishOne(t *testing.T) {
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
                body: []byte(fmt.Sprintf("PublishOne-%s", cast.ToString(time.Now().UnixNano()/1e3))),
            },
        },
        {
            args: args{
                body: []byte(fmt.Sprintf("PublishOne-%s", cast.ToString(time.Now().UnixNano()/1e3))),
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            cli, err := NewRabbitMQClient(
                WithUrl(myUrl),
                WithQueueDeclare(queueNameOne),
            )
            if err != nil {
                fmt.Println("NewRabbitMQClient err:", err)
            }
            if err := cli.PublishOne(tt.args.body); (err != nil) != tt.wantErr {
                t.Errorf("PublishOne() error = %v, wantErr %v", err, tt.wantErr)
            }
            fmt.Println("test:PublishOne:msg:", string(tt.args.body))
        })
    }
}

func TestRabbitMQClient_PublishTwo(t *testing.T) {
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
                body: []byte(fmt.Sprintf("PublishOne-%s", cast.ToString(time.Now().UnixNano()))),
            },
        },
        {
            args: args{
                body: []byte(fmt.Sprintf("PublishOne-%s", cast.ToString(time.Now().UnixNano()))),
            },
        },
        {
            args: args{
                body: []byte(fmt.Sprintf("PublishOne-%s", cast.ToString(time.Now().UnixNano()))),
            },
        },
        {
            args: args{
                body: []byte(fmt.Sprintf("PublishOne-%s", cast.ToString(time.Now().UnixNano()))),
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            cli, err := NewRabbitMQClient(
                WithUrl(myUrl),
                WithQueueDeclare(
                    queueNameTwo,
                    WithQueueDeclareDurable(true),
                ),
            )
            if err != nil {
                fmt.Println("NewRabbitMQClient err:", err)
            }
            if err := cli.PublishTwo(tt.args.body); (err != nil) != tt.wantErr {
                t.Errorf("PublishOne() error = %v, wantErr %v", err, tt.wantErr)
            }
            fmt.Println("test:PublishOne:msg:", string(tt.args.body))
        })
    }
}
