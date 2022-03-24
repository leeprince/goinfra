package lock

import "time"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/3/23 下午10:17
 * @Desc:
 */

const (
    DefaultTickerTime = time.Millisecond * 200
    DefaultTimeOut    = time.Second * 1 // 相当于总共尝试获取1+5次锁
)

const (
    DefaultLockExpireTime = time.Second * 2 // 默认的锁过期时间
)