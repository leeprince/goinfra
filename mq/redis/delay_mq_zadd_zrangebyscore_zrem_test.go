package redis

import (
    "fmt"
    "github.com/leeprince/goinfra/storage/redis"
    "testing"
    "time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/5/5 下午9:34
 * @Desc:
 */

func TestSortSetDelayMQ_Push(t *testing.T) {
    type fields struct {
        cli redis.RedisClient
    }
    type args struct {
        key       string
        value     interface{}
        delayTime time.Duration
    }
    tests := []struct {
        name    string
        fields  fields
        args    args
        wantErr bool
    }{
        {
            args:args{
                key:       "k",
                value:     "SortSetDelayMQ-v01",
                delayTime: time.Second * 5,
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            mq := NewSortSetDelayMQ(initRedisClient())
            if err := mq.Push(tt.args.key, tt.args.value, tt.args.delayTime); (err != nil) != tt.wantErr {
                t.Errorf("Push() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}

func TestSortSetDelayMQ_Subscribe(t *testing.T) {
    type fields struct {
        cli redis.RedisClient
    }
    type args struct {
        key      string
        waitTime time.Duration
    }
    tests := []struct {
        name     string
        fields   fields
        args     args
        wantData []byte
        wantErr  bool
    }{
        {
            args:args{
                key:      "k",
                waitTime: time.Second * 1,
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            mq := NewSortSetDelayMQ(initRedisClient())
            gotData, err := mq.Subscribe(tt.args.key, tt.args.waitTime)
            if (err != nil) != tt.wantErr {
                t.Errorf("Subscribe() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            fmt.Println("gotData:", string(gotData))
        })
    }
}