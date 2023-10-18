package gostreaming

/*
 * @Date: 2020-07-06 12:45:43
 * @LastEditors: aiden.deng (Zhenpeng Deng)
 * @LastEditTime: 2020-07-23 15:57:44
 */

import "github.com/go-redis/redis"

type StatusStorage interface {
	ExecBatch(BatchInterface) ([]interface{}, []error, error)

	AsRedis() *redis.Client
}
