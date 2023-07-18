package idutil

import (
	"fmt"
	"sync"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/5/8 20:43
 * @Desc:	通过雪花算法生成唯一ID: 并发性好
 */

const (
	workerBits     uint8 = 10                        // 工作ID所占位数
	maxWorker      int64 = -1 ^ (-1 << workerBits)   // 支持的最大工作ID数量。最大1023。int64(-1)二进制表示为：64个1;^表示为异或：只有0^1时才等于1
	sequenceBits   uint8 = 12                        // 序列号所占的位数
	sequenceMask   int64 = -1 ^ (-1 << sequenceBits) // 序列号的最大值
	workerShift    uint8 = sequenceBits              // 工作ID左移的位数
	timestampShift uint8 = sequenceBits + workerBits // 时间戳左移的位数
)

type Snowflake struct {
	mutex     sync.Mutex // 互斥锁，确保并发安全
	timestamp int64      // 上一次生成ID的时间
	workerId  int64      // 工作ID
	sequence  int64      // 序列号
}

func NewSnowflake(workerId int64) *Snowflake {
	if workerId < 0 || workerId > maxWorker {
		panic(fmt.Sprintf("worker ID must be between 0 and %d", maxWorker))
	}
	return &Snowflake{
		timestamp: 0,
		workerId:  workerId,
		sequence:  0,
	}
}

func (sf *Snowflake) NextId() int64 {
	sf.mutex.Lock()         // 加锁
	defer sf.mutex.Unlock() // 解锁
	
	divInt := int64(10)
	nowUnixNano := time.Now().UnixNano() / divInt // 获取当前时间，单位看出除于的值
	if sf.timestamp == nowUnixNano {              // 如果当前时间戳与上一次生成ID的时间戳相同
		sf.sequence = (sf.sequence + 1) & sequenceMask // 序列号递增，并与序列号的最大值进行按位与运算
		if sf.sequence == 0 {                          // 如果序列号达到了最大值
			for nowUnixNano <= sf.timestamp { // 等待下一个获取当前时间
				nowUnixNano = time.Now().UnixNano() / divInt
			}
		}
	} else { // 如果当前时间戳与上一次生成ID的时间戳不同
		sf.sequence = 0 // 序列号重置0
	}
	
	sf.timestamp = nowUnixNano                                                             // 更新上一次生成ID的时间戳
	return (sf.timestamp << timestampShift) | (sf.workerId << workerShift) | (sf.sequence) // 组合时间戳、工作ID和序列号，生成唯ID并返回
}
