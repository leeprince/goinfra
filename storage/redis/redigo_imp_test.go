package redis

import (
    "context"
    "fmt"
    "github.com/gomodule/redigo/redis"
    "github.com/leeprince/goinfra/config"
    "testing"
    "time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/3/25 上午12:20
 * @Desc:
 */

var (
    RedisName = "local"
    RedisConfs = config.RedisConfs{
        RedisName: config.RedisConf{
            Network:  "tcp",
            Addr:     "127.0.0.1:6379",
            Username: "",
            Password: "",
            DB:       0,
            PoolSize: 2,
        },
    }
)

var redisClientRedigo *Redigo
func InitRedigoClient()  {
    // Redigo 客户端
    err := InitRedigo(RedisConfs)
    if err != nil {
        fmt.Printf("[goinfraRedis.InitGoredis] err:%v \n", err)
    }
    redisClientRedigo = GetRedigo(RedisName)
}

func TestRedigo_SetNx(t *testing.T) {
    InitRedigoClient()
    
    type fields struct {
        ctx  context.Context
        Pool redis.Pool
    }
    type args struct {
        key        string
        value      interface{}
        expiration time.Duration
    }
    tests := []struct {
        name    string
        fields  fields
        args    args
        want    bool
        wantErr bool
    }{
        {
            args:args{
                key:        "k001",
                value:      "v001",
                expiration: time.Second * 10,
            },
        },
        {
            args:args{
                key:        "k002",
                value:      "v002",
                expiration: time.Second * 10,
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := redisClientRedigo.SetNx(tt.args.key, tt.args.value, tt.args.expiration)
            fmt.Println("got, err:", got, err)
        })
    }
}

// TODO: 测试。GetAndDel 返回 bool, error - prince@todo 2022/3/26 下午2:31
func TestRedigo_GetAndDel(t *testing.T) {
    type fields struct {
        ctx  context.Context
        Pool redis.Pool
    }
    type args struct {
        key   string
        value interface{}
    }
    tests := []struct {
        name    string
        fields  fields
        args    args
        wantErr bool
    }{
        // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            c := &Redigo{
                ctx:  tt.fields.ctx,
                Pool: tt.fields.Pool,
            }
            if err := c.GetAndDel(tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
                t.Errorf("GetAndDel() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}