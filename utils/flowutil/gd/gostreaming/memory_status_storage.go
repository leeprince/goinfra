package gostreaming

/*
 * @Date: 2020-09-03 17:18:41
 * @LastEditors: aiden.deng(Zhenpeng Deng)
 * @LastEditTime: 2020-09-03 18:01:58
 */

import (
	"errors"
	"fmt"
	"sync"

	"github.com/go-redis/redis"
)

var _ (StatusStorage) = (*MemoryStatusStorage)(nil)

type MemoryStatusStorage struct {
	mu        sync.Mutex
	hashTable map[string]interface{}
	set       map[string]interface{}
}

func NewMemoryStatusStorage() StatusStorage {
	m := &MemoryStatusStorage{
		hashTable: make(map[string]interface{}),
	}
	return m
}

func (m *MemoryStatusStorage) ExecBatch(batch BatchInterface) ([]interface{}, []error, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	cmds := batch.GetBatchCommands()
	results := make([]interface{}, len(cmds))
	errs := make([]error, len(cmds))
	for i, cmd := range cmds {
		switch cmd.Name {
		// set
		case "set":
			m.hashTable[cmd.Key] = cmd.Value
			results[i] = nil
			errs[i] = nil

		// get
		case "get":
			val, ok := m.hashTable[cmd.Key]
			if !ok {
				results[i] = nil
				errs[i] = ErrNil
				continue
			}
			results[i] = val
			errs[i] = nil

		// incr
		case "incr":
			val, ok := m.hashTable[cmd.Key]
			if !ok {
				val = 0
			}
			intVal, ok := val.(int)
			if !ok {
				results[i] = nil
				errs[i] = errors.New("WRONGTYPE Operation against a key holding the wrong kind of value")
				continue
			}
			intVal++
			m.hashTable[cmd.Key] = intVal
			results[i] = nil
			errs[i] = nil

		// sadd
		case "sadd":
			m.set[cmd.Key] = cmd.Value
			results[i] = nil
			errs[i] = nil

		// scard
		case "scard":
			results[i] = len(m.set)
			errs[i] = nil

		// smembers
		case "smembers":
			members, ok := m.set[cmd.Key]
			if !ok {
				results[i] = nil
				errs[i] = ErrNil
				continue
			}
			results[i] = members
			errs[i] = nil

		case "hincrby":
			// pipeline.HIncrBy(cmd.Key, cmd.Field, int64(cmd.IncrBy))
		case "hset":
			// pipeline.HSet(cmd.Key, cmd.Field, cmd.Value)
		case "hget":
			// pipeline.HGet(cmd.Key, cmd.Field)
		case "hlen":
			// pipeline.HLen(cmd.Key)
		case "zadd":
			// pipeline.ZAdd(cmd.Key, redis.Z{Score: cmd.Score, Member: cmd.Member})
		case "zrangebyscorewithscores":
			// pipeline.ZRangeByScoreWithScores(cmd.Key, redis.ZRangeBy{
			// 	Min: cmd.Min, Max: cmd.Max, Offset: cmd.Offset, Count: cmd.Count,
			// })
		case "zrevrangebyscorewithscores":
			// pipeline.ZRevRangeByScoreWithScores(cmd.Key, redis.ZRangeBy{
			// 	Max: cmd.Max, Min: cmd.Min, Offset: cmd.Offset, Count: cmd.Count,
			// })
		case "zcount":
			// pipeline.ZCount(cmd.Key, cmd.Min, cmd.Max)
		default:
			return nil, nil, fmt.Errorf("unknown cmd name: %s", cmd.Name)
		}
	}

	return results, errs, nil
}

func (m *MemoryStatusStorage) AsRedis() *redis.Client {
	return nil
}
