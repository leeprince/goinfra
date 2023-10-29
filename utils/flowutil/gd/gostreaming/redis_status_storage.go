package gostreaming

/*
 * @Date: 2020-07-06 13:49:48
 * @LastEditors: aiden.deng(Zhenpeng Deng)
 * @LastEditTime: 2021-04-15 11:56:26
 */

import (
	"fmt"

	"github.com/go-redis/redis"
)

var _ (StatusStorage) = (*RedisStatusStorage)(nil)

type RedisStatusStorage struct {
	db *redis.Client
}

func NewRedisStatusStorage(redisCli *redis.Client) StatusStorage {
	r := &RedisStatusStorage{
		db: redisCli,
	}
	return r
}

func (s *RedisStatusStorage) ExecBatch(batch BatchInterface) ([]interface{}, []error, error) {
	// 开启事务+批量操作
	pipeline := s.db.TxPipeline()
	// pipeline := s.db.Pipeline()

	for _, cmd := range batch.GetBatchCommands() {
		switch cmd.Name {
		case "expire":
			pipeline.Expire(cmd.Key, cmd.Expiration)
		case "set":
			pipeline.Set(cmd.Key, cmd.Value, cmd.Expiration)
		case "get":
			pipeline.Get(cmd.Key)
		case "incr":
			pipeline.Incr(cmd.Key)
		case "incrbyfloat":
			pipeline.IncrByFloat(cmd.Key, cmd.Value.(float64))
		case "sadd":
			pipeline.SAdd(cmd.Key, cmd.Value)
		case "scard":
			pipeline.SCard(cmd.Key)
		case "smembers":
			pipeline.SMembers(cmd.Key)
		case "hincrby":
			pipeline.HIncrBy(cmd.Key, cmd.Field, int64(cmd.IncrBy))
		case "hset":
			pipeline.HSet(cmd.Key, cmd.Field, cmd.Value)
		case "hget":
			pipeline.HGet(cmd.Key, cmd.Field)
		case "hlen":
			pipeline.HLen(cmd.Key)
		case "zadd":
			pipeline.ZAdd(cmd.Key, redis.Z{Score: cmd.Score, Member: cmd.Member})
		case "zrangebyscorewithscores":
			pipeline.ZRangeByScoreWithScores(cmd.Key, redis.ZRangeBy{
				Min: cmd.Min, Max: cmd.Max, Offset: cmd.Offset, Count: cmd.Count,
			})
		case "zrevrangebyscorewithscores":
			pipeline.ZRevRangeByScoreWithScores(cmd.Key, redis.ZRangeBy{
				Max: cmd.Max, Min: cmd.Min, Offset: cmd.Offset, Count: cmd.Count,
			})
		case "zcount":
			pipeline.ZCount(cmd.Key, cmd.Min, cmd.Max)
		case "zcard":
			pipeline.ZCard(cmd.Key)
		case "zremrangebyscore":
			pipeline.ZRemRangeByScore(cmd.Key, cmd.Min, cmd.Max)
		default:
			return nil, nil, fmt.Errorf("unknown cmd name: %s", cmd.Name)
		}
	}
	cmders, fatalErr := pipeline.Exec()
	if fatalErr == redis.Nil {
		fatalErr = nil
	}

	results := make([]interface{}, 0)
	errors := make([]error, 0)

	for _, cmder := range cmders {
		err := cmder.Err()
		if err != nil {
			results = append(results, nil)
			errors = append(errors, err)
			continue
		}
		errors = append(errors, nil)

		switch typedCmder := cmder.(type) {
		case *redis.StatusCmd:
			results = append(results, typedCmder.Val())
		case *redis.IntCmd:
			results = append(results, int(typedCmder.Val()))
		case *redis.FloatCmd:
			results = append(results, float64(typedCmder.Val()))
		case *redis.StringCmd:
			results = append(results, typedCmder.Val())
		case *redis.BoolCmd:
			results = append(results, typedCmder.Val())
		case *redis.StringSliceCmd:
			results = append(results, typedCmder.Val())
		case *redis.ZSliceCmd:
			results = append(results, typedCmder.Val())
		default:
			return nil, nil, fmt.Errorf("gostreaming: unimplemented typeCmder: %+v", typedCmder)
		}
	}
	return results, errors, fatalErr
}

func (s *RedisStatusStorage) AsRedis() *redis.Client {
	return s.db
}
