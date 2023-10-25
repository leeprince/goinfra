package redis

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/10/23 10:15
 * @Desc:
 */

var (
	redisName  = "local"
	redisConfs = RedisConfs{
		redisName: RedisConf{
			Network:  "tcp",
			Addr:     "127.0.0.1:6379",
			Username: "",
			Password: "",
			DB:       0,
			PoolSize: 2,
		},
	}
)
