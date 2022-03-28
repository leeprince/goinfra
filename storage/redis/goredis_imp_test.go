package redis

import (
    "context"
    "fmt"
    "github.com/go-redis/redis/v8"
    "testing"
    "time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/3/25 上午12:20
 * @Desc:
 */

var redisClientGoredis *Goredis

func InitGoredisClient() {
    // Goredis 客户端
    err := InitGoredis(RedisConfs)
    if err != nil {
        fmt.Printf("[goinfraRedis.InitGoredis] err:%v \n", err)
    }
    redisClientGoredis = GetGoredis(RedisName)
}

func TestGoredis_SetNx(t *testing.T) {
    InitGoredisClient()
    
    type fields struct {
        ctx  context.Context
        Client *redis.Client
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
            args: args{
                key:        "k001",
                value:      "v001",
                expiration: time.Second * 10,
            },
        },
        {
            args: args{
                key:        "k002",
                value:      "v002",
                expiration: time.Second * 10,
            },
        },
        {
            args: args{
                key:        "k002",
                value:      "v002",
                expiration: time.Second * 10,
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := redisClientGoredis.SetNx(tt.args.key, tt.args.value, tt.args.expiration)
            fmt.Println("got, err:", got, "--", err)
        })
    }
}

func TestGoredis_GetAndDel(t *testing.T) {
    InitGoredisClient()
    
    type fields struct {
        ctx  context.Context
        Client *redis.Client
    }
    type args struct {
        key   string
        value interface{}
        expiration time.Duration
    }
    tests := []struct {
        name    string
        fields  fields
        args    args
        wantErr bool
    }{
        {
            args: args{
                key:        "k001",
                value:      "v001",
                expiration: time.Second * 10,
                
            },
        },
        {
            args: args{
                key:        "k001",
                value:      "v002",
                expiration: time.Second * 10,
                
            },
        },
        {
            args: args{
                key:        "k001",
                value:      "v002",
                expiration: time.Second * 10,
                
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            GetAndDelErr := redisClientGoredis.GetAndDel(tt.args.key, tt.args.value)
            fmt.Println("GetAndDel>>>>>>>>", GetAndDelErr)
            
            setErr := redisClientGoredis.SetKey(tt.args.key, tt.args.value, tt.args.expiration)
            fmt.Println("SetKey::::::::::::", setErr)
            
        })
    }
}
