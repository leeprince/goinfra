package redis

import (
    "fmt"
    "github.com/leeprince/goinfra/storage/redis"
    "testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/5/10 下午6:07
 * @Desc:
 */

func TestPubishSubscribeMQ_Subscribe(t *testing.T) {
    type fields struct {
        cli redis.RedisClient
    }
    type args struct {
        f        pubishSubscribeMQSubscribeFunc
        channels []string
    }
    tests := []struct {
        name   string
        fields fields
        args   args
    }{
        // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            mq := NewPubishSubscribeMQ(initRedisClient())
            
            callbackFunc := func(data <- chan redis.SubscribeChannelMessage) {
                // for 与 select...case... 同样能接收通道（channel）的数据
                select {
                case msg := <-data:
                    fmt.Println("msg.Channel:", msg.Channel)
                    fmt.Println("msg.Pattern:", msg.Pattern)
                    fmt.Println("msg.Payload:", msg.Payload)
                    fmt.Println("msg.PayloadSlice:", msg.PayloadSlice)
                }
            }
            mq.Subscribe(callbackFunc, "k")
        })
    }
}