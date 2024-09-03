package idutil

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/5/8 20:43
 * @Desc:	通过雪花算法生成唯一ID: 并发性好
 */

const (
	// 时间戳偏移量
	epoch = int64(1288834974657) // 1262304000000：2010-01-01 08:00:00；1288834974657（兼容 mybatis plus 内置的雪花算法）：2010-11-04 09:42:54
	// 数据中心ID的最大值
	maxDatacenterId = 31 // 5 bits
	// 工作节点ID的最大值
	maxWorkerId = 31 // 5 bits
	// 序列号的最大值
	maxSequence = 4095 // 12 bits
)

var (
	mu           sync.Mutex
	datacenterId int64
	workerId     int64
	sequence     = int64(0)
	// 保存的是上一次生成ID的时间戳（通常是毫秒精度）
	/*
	   如果当前时间戳大于 lastTimestamp：
	       这意味着时间向前推进了，可以正常生成新的ID。
	       lastTimestamp 的值会被更新为当前的时间戳。
	   如果当前时间戳等于 lastTimestamp：
	       这意味着还在同一个毫秒内，此时可以通过增加序列号 sequence 来生成不同的ID。
	       如果序列号已经达到了最大值 (maxSequence)，则需要等待直到下一个毫秒到来，再次尝试生成ID。
	   如果当前时间戳小于 lastTimestamp：
	       这意味着系统时钟回退了，这是不允许的情况，因为这会导致生成的ID不具有单调递增性。
	       通常在这种情况下程序会抛出异常或错误。
	*/
	lastTimestamp = int64(1)
)

func init() {
	randNew := rand.New(rand.NewSource(time.Now().UnixNano()))
	datacenterId = randNew.Int63n(maxDatacenterId + 1)
	workerId = randNew.Int63n(maxWorkerId + 1)
}

type SnowflakeGenerator struct{}

// Generate datacenterId=dataAndWorkId[0];workerId=dataAndWorkId[1]
func (g *SnowflakeGenerator) Generate(dataIdAndWorkId ...int64) int64 {
	mu.Lock()
	defer mu.Unlock()
	if len(dataIdAndWorkId) == 2 {
		datacenterId = dataIdAndWorkId[0]
		workerId = dataIdAndWorkId[1]
	}
	if len(dataIdAndWorkId) == 1 {
		datacenterId = dataIdAndWorkId[0]
	}
	timestamp := time.Now().UnixMilli()
	if timestamp < lastTimestamp {
		panic(fmt.Sprintf("Clock moved backwards. Refusing to generate id for %d milliseconds", lastTimestamp-timestamp))
	}
	if timestamp == lastTimestamp {
		sequence++
		if sequence > maxSequence {
			for timestamp <= lastTimestamp {
				timestamp = time.Now().UnixMilli()
			}
			sequence = 0
		}
	} else {
		sequence = 0
	}
	lastTimestamp = timestamp
	// fmt.Println(">>>>timestamp",
	// 	timestamp,
	// 	datacenterId,
	// 	workerId,
	// 	sequence)
	return (((timestamp-epoch)<<10)|(datacenterId<<5)|workerId)<<12 | sequence
}
