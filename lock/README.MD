# 锁


## 分布式锁

分布式锁需满足条件：
1. 支持分布式。# redis
2. 支持自己加锁及只能解自己的锁
3. 设置锁过期时间。防止程序异常退出无妨释放锁，导致一直锁等待

### redigo
基于 `github.com/gomodule/redigo/redis` 实现

### go_redis
基于 `github.com/go-redis/redis/v8` 实现
