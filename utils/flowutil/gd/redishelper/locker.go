package redishelper

/*
 * @Date: 2020-08-06 09:41:04
 * @LastEditors: aiden.deng (Zhenpeng Deng)
 * @LastEditTime: 2020-08-06 09:50:50
 */
import (
	"time"

	"github.com/go-redis/redis"
)

func Lock(db *redis.Client, key, secret string, expiration time.Duration) (bool, error) {
	ok, err := db.SetNX(key, secret, expiration).Result()
	if err != nil {
		return false, err
	}
	return ok, nil
}

func Unlock(db *redis.Client, key, secret string) error {
	scripts := "if redis.call('get', KEYS[1]) == ARGV[1] then return redis.call('del', KEYS[1]) else return 0 end"
	_, err := db.Eval(scripts, []string{key}, secret).Result()
	if err != nil {
		return err
	}
	return nil
}
