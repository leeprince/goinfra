package redigo

import (
    "fmt"
    "github.com/gomodule/redigo/redis"
)

type Lock struct {
    key        string
    value      string
    conn       redis.Conn
    expireTime int
}

func NewTryLock(conn redis.Conn, key string, value string, DefaulexpireTime int) (lock *Lock, ok bool, err error) {
    return TryLockV1WithExpireTime(conn, key, value, DefaulexpireTime)
}

func TryLockV1WithExpireTime(conn redis.Conn, key string, value string, expireTime int) (lock *Lock, ok bool, err error) {
    lock = &Lock{key, value, conn, expireTime}
    ok, err = lock.TryLockV1()
    
    if !ok || err != nil {
        fmt.Println("[TryLockV1WithExpireTime] fail. ok, err:", ok, err)
        lock = nil
        return
    }
    fmt.Println("[TryLockV1WithExpireTime] Suucessfuly")
    
    return
}

func (lock *Lock) Unlock() (err error) {
    _, err = lock.conn.Do("del", lock.key)
    return
}

func (lock *Lock) AddExpireTime(exTime int64) (ok bool, err error) {
    ttlTime, err := redis.Int64(lock.conn.Do("TTL", lock.key))
    if err != nil {
        fmt.Println("[AddExpireTime] lock.key TTL", ttlTime)
        return
    }
    fmt.Println("[AddExpireTime] lock.key TTL", ttlTime)
    
    if ttlTime > 0 {
        _, err := redis.String(lock.conn.Do("SET", lock.key, lock.value, "EX", int(ttlTime+exTime)))
        if err == redis.ErrNil {
            fmt.Println("[AddExpireTime] err == redis.ErrNil ")
            return false, nil
        }
        if err != nil {
            fmt.Println("[AddExpireTime] err != nil")
            return false, err
        }
    }
    return false, nil
}

func (lock *Lock) TryLockV1() (ok bool, err error) {
    _, err = redis.String(lock.conn.Do("SET", lock.key, lock.value, "EX", int(lock.expireTime), "NX"))
    if err == redis.ErrNil {
        fmt.Println("[TryLockV1] err == redis.ErrNil")
        return false, nil
    }
    if err != nil {
        fmt.Println("[TryLockV1] err != nil")
        return false, err
    }
    return true, nil
}
