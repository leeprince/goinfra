package gostreaming

/*
 * @Date: 2020-07-09 09:47:18
 * @LastEditors: aiden.deng(Zhenpeng Deng)
 * @LastEditTime: 2021-04-15 11:56:04
 */

import (
	"fmt"
	"sync"
	"time"
)

var _ BatchInterface = (*Batch)(nil)

// Batch 是StatusStorage的批量读写接口实现。Thread safe
type Batch struct {
	mu   sync.Mutex
	cmds []*BatchCommand
}

func NewBatch() *Batch {
	return &Batch{
		cmds: make([]*BatchCommand, 0),
	}
}

func (b *Batch) Size() int {
	b.mu.Lock()
	defer b.mu.Unlock()
	return len(b.cmds)
}

func (b *Batch) GetBatchCommands() []*BatchCommand {
	return b.cmds
}

func (b *Batch) String() string {
	b.mu.Lock()
	defer b.mu.Unlock()
	return fmt.Sprintf("<Batch cmds=%+v>", b.cmds)
}

// Read
func (b *Batch) Get(primaryKeys []string, targetName string, descriptions []string) {
	b.mu.Lock()
	defer b.mu.Unlock()
	key := makeKeyName(primaryKeys, targetName, descriptions)
	b.cmds = append(b.cmds, &BatchCommand{Name: "get", Key: key})
}

func (b *Batch) SCard(primaryKeys []string, targetName string, descriptions []string) {
	b.mu.Lock()
	defer b.mu.Unlock()
	key := makeKeyName(primaryKeys, targetName, descriptions)
	b.cmds = append(b.cmds, &BatchCommand{Name: "scard", Key: key})
}

func (b *Batch) SMembers(primaryKeys []string, targetName string, descriptions []string) {
	b.mu.Lock()
	defer b.mu.Unlock()
	key := makeKeyName(primaryKeys, targetName, descriptions)
	b.cmds = append(b.cmds, &BatchCommand{Name: "smembers", Key: key})
}

func (b *Batch) HLen(primaryKeys []string, targetName string, descriptions []string) {
	b.mu.Lock()
	defer b.mu.Unlock()
	key := makeKeyName(primaryKeys, targetName, descriptions)
	b.cmds = append(b.cmds, &BatchCommand{Name: "hlen", Key: key})
}

// Write
func (b *Batch) Incr(primaryKeys []string, targetName string, descriptions []string, expiration time.Duration) {
	b.mu.Lock()
	defer b.mu.Unlock()
	key := makeKeyName(primaryKeys, targetName, descriptions)
	b.cmds = append(b.cmds, &BatchCommand{Name: "incr", Key: key})
	b.cmds = append(b.cmds, &BatchCommand{Name: "expire", Key: key, Expiration: expiration})
}

func (b *Batch) IncrByFloat(primaryKeys []string, targetName string, descriptions []string, value float64, expiration time.Duration) {
	b.mu.Lock()
	defer b.mu.Unlock()
	key := makeKeyName(primaryKeys, targetName, descriptions)
	b.cmds = append(b.cmds, &BatchCommand{Name: "incrbyfloat", Key: key, Value: value})
	b.cmds = append(b.cmds, &BatchCommand{Name: "expire", Key: key, Expiration: expiration})
}

func (b *Batch) Set(primaryKeys []string, targetName string, descriptions []string, value interface{}, expiration time.Duration) {
	b.mu.Lock()
	defer b.mu.Unlock()
	key := makeKeyName(primaryKeys, targetName, descriptions)
	b.cmds = append(b.cmds, &BatchCommand{Name: "set", Key: key, Value: value, Expiration: expiration})
}

func (b *Batch) SAdd(primaryKeys []string, targetName string, descriptions []string, value interface{}, expiration time.Duration) {
	b.mu.Lock()
	defer b.mu.Unlock()
	key := makeKeyName(primaryKeys, targetName, descriptions)
	b.cmds = append(b.cmds, &BatchCommand{Name: "sadd", Key: key, Value: value})
	b.cmds = append(b.cmds, &BatchCommand{Name: "expire", Key: key, Expiration: expiration})
}

func (b *Batch) HIncrBy(primaryKeys []string, targetName string, descriptions []string, field string, incrBy int) {
	b.mu.Lock()
	defer b.mu.Unlock()
	key := makeKeyName(primaryKeys, targetName, descriptions)
	b.cmds = append(b.cmds, &BatchCommand{Name: "hincrby", Key: key, Field: field, IncrBy: incrBy})
}

func (b *Batch) HSet(primaryKeys []string, targetName string, descriptions []string, field string, value interface{}) {
	b.mu.Lock()
	defer b.mu.Unlock()
	key := makeKeyName(primaryKeys, targetName, descriptions)
	b.cmds = append(b.cmds, &BatchCommand{Name: "hset", Key: key, Field: field, Value: value})
}

func (b *Batch) HGet(primaryKeys []string, targetName string, descriptions []string, field string) {
	b.mu.Lock()
	defer b.mu.Unlock()
	key := makeKeyName(primaryKeys, targetName, descriptions)
	b.cmds = append(b.cmds, &BatchCommand{Name: "hget", Key: key, Field: field})
}

func (b *Batch) ZAdd(primaryKeys []string, targetName string, descriptions []string, score float64, member interface{}) {
	b.mu.Lock()
	defer b.mu.Unlock()
	key := makeKeyName(primaryKeys, targetName, descriptions)
	b.cmds = append(b.cmds, &BatchCommand{Name: "zadd", Key: key, Score: score, Member: member})
}

func (b *Batch) ZRangeByScoreWithScores(primaryKeys []string, targetName string, descriptions []string,
	min string, max string, offset int64, count int64) {
	b.mu.Lock()
	defer b.mu.Unlock()
	key := makeKeyName(primaryKeys, targetName, descriptions)
	b.cmds = append(b.cmds, &BatchCommand{Name: "zrangebyscorewithscores", Key: key, Min: min, Max: max, Offset: offset, Count: count})
}

func (b *Batch) ZRevRangeByScoreWithScores(primaryKeys []string, targetName string, descriptions []string,
	max string, min string, offset int64, count int64) {
	b.mu.Lock()
	defer b.mu.Unlock()
	key := makeKeyName(primaryKeys, targetName, descriptions)
	b.cmds = append(b.cmds, &BatchCommand{Name: "zrevrangebyscorewithscores", Key: key, Min: min, Max: max, Offset: offset, Count: count})
}

func (b *Batch) ZCount(primaryKeys []string, targetName string, descriptions []string, min string, max string) {
	b.mu.Lock()
	defer b.mu.Unlock()
	key := makeKeyName(primaryKeys, targetName, descriptions)
	b.cmds = append(b.cmds, &BatchCommand{Name: "zcount", Key: key, Min: min, Max: max})
}

func (b *Batch) ZCard(primaryKeys []string, targetName string, descriptions []string) {
	b.mu.Lock()
	defer b.mu.Unlock()
	key := makeKeyName(primaryKeys, targetName, descriptions)
	b.cmds = append(b.cmds, &BatchCommand{Name: "zcard", Key: key})
}

func (b *Batch) ZRemRangeByScore(primaryKeys []string, targetName string, descriptions []string, min string, max string) {
	b.mu.Lock()
	defer b.mu.Unlock()
	key := makeKeyName(primaryKeys, targetName, descriptions)
	b.cmds = append(b.cmds, &BatchCommand{Name: "zremrangebyscore", Key: key, Min: min, Max: max})
}
