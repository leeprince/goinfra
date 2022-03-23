package redigo

import (
    "github.com/gomodule/redigo/redis"
)

type Lock struct {
    key        string
    value      string
    conn       redis.Conn
    expireTime int
}

func NewTryLock(conn redis.Conn, key string, value string, DefaulexpireTime int) (lock *Lock, ok bool, err error) {
    return TryLockWithExpireTime(conn, key, value, DefaulexpireTime)
}

func TryLockWithExpireTime(conn redis.Conn, key string, value string, expireTime int) (lock *Lock, ok bool, err error) {
    lock = &Lock{key, value, conn, expireTime}
    ok, err = lock.TryLock()
    
    if !ok || err != nil {
        // fmt.Println("[TryLockWithExpireTime] fail. ok, err:", ok, err)
        lock = nil
        return
    }
    // fmt.Println("[TryLockWithExpireTime] Suucessfuly")
    
    return
}

func (lock *Lock) Unlock() (err error) {
    _, err = lock.conn.Do("del", lock.key)
    return
}

func (lock *Lock) AddExpireTime(exTime int64) (ok bool, err error) {
    ttlTime, err := redis.Int64(lock.conn.Do("TTL", lock.key))
    if err != nil {
        // fmt.Println("[AddExpireTime] lock.key TTL", ttlTime)
        return
    }
    // fmt.Println("[AddExpireTime] lock.key TTL", ttlTime)
    
    if ttlTime > 0 {
        _, err := redis.String(lock.conn.Do("SET", lock.key, lock.value, "EX", int(ttlTime+exTime)))
        if err == redis.ErrNil {
            // fmt.Println("[AddExpireTime] err == redis.ErrNil ")
            return false, nil
        }
        if err != nil {
            // fmt.Println("[AddExpireTime] err != nil")
            return false, err
        }
    }
    return false, nil
}

func (lock *Lock) TryLock() (ok bool, err error) {
    _, err = redis.String(lock.conn.Do("SET", lock.key, lock.value, "EX", int(lock.expireTime), "NX"))
    if err == redis.ErrNil {
        // fmt.Println("[TryLock] err == redis.ErrNil")
        return false, nil
    }
    if err != nil {
        // fmt.Println("[TryLock] err != nil")
        return false, err
    }
    return true, nil
}
