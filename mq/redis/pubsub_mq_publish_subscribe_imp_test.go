package redis

import (
	"fmt"
	"github.com/leeprince/goinfra/storage/redis"
	"testing"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/5/10 下午6:07
 * @Desc:
 */

func TestPubishSubscribeMQ_Push(t *testing.T) {
	type fields struct {
		cli redis.RedisClient
	}
	type args struct {
		channel string
		message interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			args: args{
				channel: "k",
				message: "vvvv-001",
			},
		},
		{
			args: args{
				channel: "k",
				message: "vvvv-002",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mq := NewPubishSubscribeMQ(initRedisClient())

			if err := mq.Push(tt.args.channel, tt.args.message); (err != nil) != tt.wantErr {
				t.Errorf("Push() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPubishSubscribeMQ_Subscribe(t *testing.T) {
	type fields struct {
		cli redis.RedisClient
	}
	type args struct {
		f        PubishSubscribeMQSubscribeHandle
		channels []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			args: args{
				channels: []string{"k"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mq := NewPubishSubscribeMQ(initRedisClient())

			callbackFunc := func(data *redis.SubscribeMessage) {
				fmt.Println(">>>>>>>>>>>>>> time:", time.Now().UnixNano()/1e6)

				fmt.Println("data.Channel:", data.Channel)
				fmt.Println("data.Pattern:", data.Pattern)
				fmt.Println("data.Payload:", data.Payload)
				fmt.Println("data.PayloadSlice:", data.PayloadSlice)
			}
			mq.Subscribe(callbackFunc, "k")

		})
	}
}
