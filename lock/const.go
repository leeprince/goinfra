package lock

import "time"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/3/23 下午10:17
 * @Desc:
 */

const (
    DefaultTickerTime = time.Millisecond * 500
    DefaultTimeOut    = time.Second * 2 // 相当于尝试4次+1次获取锁
)

const (
    DefaultLockExpireTime = time.Second * 2 // 默认的锁过期时间
)