package redis

import (
    "github.com/leeprince/goinfra/storage/redis"
    "time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/5/1 上午11:57
 * @Desc:   实现 MQ 接口
 */

type PushPop struct {
    cli redis.RedisClient
}

func NewPushPop(cli redis.RedisClient) *PushPop {
    return &PushPop{
        cli: cli,
    }
}

func (mq *PushPop) Publish(key string, value interface{}) error {
    return mq.cli.Push(key, value)
}
func (mq *PushPop) Subscribe(key string, timeout time.Duration) (data interface{}, err error) {
    for {
        data, err = mq.cli.BPop(key, timeout)
        if err != nil {
            return
        }
        if data == nil {
            // fmt.Println("(mq *PushPop) Subscribe data == nil")
            continue
        }
        return
    }
    
}
