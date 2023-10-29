package redishelper

/*
 * @Date: 2020-07-11 17:32:50
 * @LastEditors: aiden.deng (Zhenpeng Deng)
 * @LastEditTime: 2020-10-09 11:48:17
 */

import (
	"fmt"

	"github.com/go-redis/redis"
)

type Config struct {
	Host     string
	Port     int
	Password string
	DB       int
}

func NewClient(cfg Config) (*redis.Client, error) {
	redisCli := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	_, err := redisCli.Ping().Result()
	if err != nil {
		return nil, fmt.Errorf("create redis client failed(ping return err: %s)", err)
	}
	return redisCli, nil
}

func MustNewClient(cfg Config) *redis.Client {
	redisCli, err := NewClient(cfg)
	if err != nil {
		panic(err)
	}
	return redisCli
}

func MayNewClient(redisCfg Config) *redis.Client {
	redisCli, _ := NewClient(redisCfg)
	return redisCli
}
