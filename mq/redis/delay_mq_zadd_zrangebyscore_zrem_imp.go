package redis

import (
    "github.com/leeprince/goinfra/plog"
    "github.com/leeprince/goinfra/storage/redis"
    "github.com/spf13/cast"
    "time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/5/4 下午2:54
 * @Desc:   通过 redis `zadd`、`zrangescore`、`zrem` 有序集合命令实现 DelayMQ 接口
 */

type SortSetDelayMQ struct {
    cli redis.RedisClient
}

// SortSetDelayMQ 实现 DelayMQ 接口
var _ DelayMQ = (*SortSetDelayMQ)(nil)

func NewSortSetDelayMQ(cli redis.RedisClient) *SortSetDelayMQ {
    return &SortSetDelayMQ{
        cli: cli,
    }
}

// 支持最小延迟时间为秒
func (mq *SortSetDelayMQ) Push(key string, value interface{}, delayTime time.Duration) error {
    if delayTime == 0 {
        delayTime = DefaulDelayTime
    }
    
    z := &redis.Z{
        Score:  cast.ToFloat64(time.Now().Add(delayTime).Second()),
        Member: value,
    }
    
    return mq.cli.ZAdd(key, z)
}

func (mq *SortSetDelayMQ) Subscribe(f DelayMQSubscribeHandle, key string, waitTime time.Duration) {
    if waitTime == 0 {
        waitTime = DefaultWaitTime
    }
    
    for {
        var stringSliceData []string
        var err error
        
        stringSliceData, err = mq.cli.ZRangeByScore(key, &redis.ZRangeBy{
            Min:    "0",
            Max:    cast.ToString(time.Now().Unix()),
            Offset: 0,
            Count:  1, // 每次只取出一个成员
        })
        if err != nil {
            plog.Error("(mq *SortSetDelayMQ) Subscribe mq.cli.ZRangeByScore err:", err)
            // 防止 redis 连接断开(重启、网络抖动)后无限循环，浪费 cpu
            time.Sleep(waitTime)
            continue
        }
        if len(stringSliceData) == 0 {
            time.Sleep(waitTime)
            // fmt.Println("time.Sleep(waitTime):", time.Now().UnixNano() / 1e6)
            continue
        }
        
        // fmt.Println("stringSliceData:", stringSliceData)
        data := []byte(stringSliceData[0])
        err = mq.cli.ZRem(key, data)
        if err != nil {
            plog.Error("(mq *SortSetDelayMQ) Subscribe mq.cli.ZRem err:", err)
            continue
        }
        
        go f(data)
    }
}
