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
    RedisName  = "local"
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

func initRedigoClient() {
    // Redigo 客户端
    err := InitRedigo(RedisConfs)
    if err != nil {
        fmt.Printf("[goinfraRedis.InitGoredis] err:%v \n", err)
    }
    redisClientRedigo = GetRedigo(RedisName)
}

func TestRedigo_SetNx(t *testing.T) {
    initRedigoClient()
    
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
            got, err := redisClientRedigo.SetNx(tt.args.key, tt.args.value, tt.args.expiration)
            fmt.Println("got, err:", got, "--", err)
        })
    }
}

func TestRedigo_GetAndDel(t *testing.T) {
    initRedigoClient()
    
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
            GetAndDelErr := redisClientRedigo.GetAndDel(tt.args.key, tt.args.value)
            fmt.Println("GetAndDel>>>>>>>>", GetAndDelErr)
            
            setErr := redisClientRedigo.Set(tt.args.key, tt.args.value, tt.args.expiration)
            fmt.Println("SetKey::::::::::::", setErr)
            
        })
    }
}

func TestRedigo_Push(t *testing.T) {
    initRedigoClient()
    
    type fields struct {
        ctx  context.Context
        Pool redis.Pool
    }
    type args struct {
        key     string
        value   interface{}
        isRight []bool
    }
    tests := []struct {
        name    string
        fields  fields
        args    args
        wantErr bool
    }{
        {
            args: args{
                key:     "k",
                value:   "v001",
                isRight: nil,
            },
        },
        {
            args: args{
                key:     "k",
                value:   []string{"vv001", "vv002"},
                isRight: nil,
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if err := redisClientRedigo.Push(tt.args.key, tt.args.value, tt.args.isRight...); (err != nil) != tt.wantErr {
                t.Errorf("Push() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}

func TestRedigo_ZAdd(t *testing.T) {
    initRedigoClient()
    
    type fields struct {
        ctx  context.Context
        Pool redis.Pool
    }
    type args struct {
        key     string
        members []*Z
    }
    tests := []struct {
        name    string
        fields  fields
        args    args
        wantErr bool
    }{
        {
            args: args{
                key: "k",
                members: []*Z{
                    &Z{
                        Score:  0,
                        Member: "mm000",
                    },
                    &Z{
                        Score:  1,
                        Member: "mm001",
                    },
                    // &Z{
                    //     Score: 2,
                    //     Member: []string{ // 不支持 Z.Member 为切片、结构体
                    //         "mm00201",
                    //         "mm00202",
                    //     },
                    // },
                    // &Z{
                    //     Score:  3,
                    //     Member: struct { // 不支持 Z.Member 为切片、结构体
                    //         Name string
                    //         Age int
                    //     }{
                    //         Name: "prince01",
                    //         Age:  18,
                    //     },
                    // },
                },
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if err := redisClientRedigo.ZAdd(tt.args.key, tt.args.members...); (err != nil) != tt.wantErr {
                t.Errorf("ZAdd() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}

func TestRedigo_ZRangeByScore(t *testing.T) {
    initRedigoClient()
    
    type fields struct {
        ctx  context.Context
        Pool redis.Pool
    }
    type args struct {
        key string
        opt *ZRangeBy
    }
    tests := []struct {
        name     string
        fields   fields
        args     args
        wantData interface{}
        wantErr  bool
    }{
        {
            args: args{
                key: "k",
                opt: &ZRangeBy{
                    Min:           "0",
                    Max:           "10",
                    Offset:        0,
                    Count:         1,
                    isReturnScore: false,
                },
            },
        },
        {
            args: args{
                key: "k",
                opt: &ZRangeBy{
                    Min:           "0",
                    Max:           "10",
                    Offset:        0,
                    Count:         1,
                    isReturnScore: true,
                },
            },
        },
        {
            args: args{
                key: "k",
                opt: &ZRangeBy{
                    Min:           "0",
                    Max:           "10",
                    Offset:        0,
                    Count:         2,
                    isReturnScore: false,
                },
            },
        },
        {
            args: args{
                key: "k",
                opt: &ZRangeBy{
                    Min:           "0",
                    Max:           "10",
                    Offset:        0,
                    Count:         2,
                    isReturnScore: true,
                },
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            gotData, err := redisClientRedigo.ZRangeByScore(tt.args.key, tt.args.opt)
            if (err != nil) != tt.wantErr {
                t.Errorf("ZRangeByScore() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            fmt.Printf("gotData type:%T data:%v \n", gotData, gotData)
        })
    }
}

func TestRedigo_ZRem(t *testing.T) {
    initRedigoClient()
    
    type fields struct {
        ctx  context.Context
        Pool redis.Pool
    }
    type args struct {
        key     string
        members []interface{}
    }
    tests := []struct {
        name    string
        fields  fields
        args    args
        wantErr bool
    }{
        {
            args: args{
                key: "ZADD-key01",
                members: []interface{}{
                    "vv000",
                    "vv001",
                },
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if err := redisClientRedigo.ZRem(tt.args.key, tt.args.members...); (err != nil) != tt.wantErr {
                t.Errorf("ZRem() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}
