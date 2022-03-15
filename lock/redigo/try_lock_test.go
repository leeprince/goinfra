package redigo

import (
    "fmt"
    "github.com/gomodule/redigo/redis"
    goinfraRedis "goinfra/session/redis"
    "testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/2/26 下午5:53
 * @Desc:
 */
func TestNewTryLockV11(t *testing.T) {
    // ---
    conns, err := goinfraRedis.InitRedisConn(RedisConfs)
    if err != nil {
        fmt.Printf("[TestTryLockV1.InitRedisConn] err:%v \n", err)
        return
    }
    conn := conns[RedisName]
    // ---
    type args struct {
        conn             redis.Conn
        key              string
        value            string
        DefaulExpireTime int
    }
    tests := []struct {
        name     string
        args     args
        wantLock *Lock
        wantOk   bool
        wantErr  bool
    }{
        {
            args: args{
                conn:             conn,
                key:              LockKey,
                value:            LockValue,
                DefaulExpireTime: int(LockExpire.Seconds()),
            },
        },
        {
            args: args{
                conn:             conn,
                key:              LockKey,
                value:            LockValue,
                DefaulExpireTime: int(LockExpire.Seconds()),
            },
        },
    }
    for i, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            gotLock, gotOk, err := NewTryLock(tt.args.conn, tt.args.key, tt.args.value, tt.args.DefaulExpireTime)
            fmt.Printf("#%d[NewTryLock() gotLock:%v",i, gotLock)
            if err != nil {
                t.Errorf("#%dNewTryLockV1() error = %v \n", i, err)
                return
            }
            if !gotOk {
                t.Errorf("#%dNewTryLockV1() !gotOk \n", i)
                return
            }
            
            err = gotLock.Unlock()
            // err = gotLock.Unlock()
            fmt.Printf("#%d[NewTryLock() gotLock.Unlock() err:%v \n",i,  err)
        })
    }
}
