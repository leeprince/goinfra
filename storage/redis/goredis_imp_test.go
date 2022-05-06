package redis

import (
    "context"
    "encoding/json"
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

func initGoredisClient() *Goredis {
    // Goredis 客户端
    err := InitGoredis(RedisConfs)
    if err != nil {
        panic(fmt.Sprintf("[goinfraRedis.InitGoredis] err:%v \n", err))
    }
    return GetGoredis(RedisName)
}

type ValueStruct struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func (s ValueStruct) MarshalBinary() ([]byte, error) {
    return json.Marshal(s)
}

func TestGoredis_SetNx(t *testing.T) {
    
    type fields struct {
        ctx    context.Context
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
            got, err := initGoredisClient().SetNx(tt.args.key, tt.args.value, tt.args.expiration)
            fmt.Println("got, err:", got, "--", err)
        })
    }
}

func TestGoredis_GetAndDel(t *testing.T) {
    
    type fields struct {
        ctx    context.Context
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
            GetAndDelErr := initGoredisClient().GetAndDel(tt.args.key, tt.args.value)
            fmt.Println("GetAndDel>>>>>>>>", GetAndDelErr)
            
            setErr := initGoredisClient().Set(tt.args.key, tt.args.value, tt.args.expiration)
            fmt.Println("SetKey::::::::::::", setErr)
            
        })
    }
}

func TestGoredis_Push(t *testing.T) {
    
    type fields struct {
        ctx context.Context
        cli *redis.Client
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
                value:   "v-001",
                isRight: nil,
            },
        },
        {
            args: args{
                key:     "k",
                value:   []string{"v-v001", "v-v002"},
                isRight: nil,
            },
        },
        {
            args: args{
                key: "k",
                value: ValueStruct{
                    Name: "prince",
                    Age:  18,
                },
                isRight: nil,
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if err := initGoredisClient().Push(tt.args.key, tt.args.value, tt.args.isRight...); (err != nil) != tt.wantErr {
                t.Errorf("Push() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}

func TestGoredis_ZAdd(t *testing.T) {
    
    type fields struct {
        ctx context.Context
        cli *redis.Client
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
                        Member: "m000",
                    },
                    &Z{
                        Score:  1,
                        Member: "m001",
                    },
                    // &Z{
                    //     Score:  2,
                    //     Member: []string{ // value 为结构体或者部分命令传入切片（ZAdd 方法的 Z.Member 为切片）时：需实现 `encoding.BinaryMarshaler` 接口(MarshalBinary 方法), 否则报错`redis: can't marshal []string (implement encoding.BinaryMarshaler)`。建议直接转成 json string 或者 []byte
                    //         "m00201",
                    //         "m00202",
                    //     },
                    // },
                    // &Z{
                    //     Score:  3,
                    //     Member: struct { // value 为结构体或者部分命令传入切片（ZAdd 方法的 Z.Member 为切片）时：需实现 `encoding.BinaryMarshaler` 接口(MarshalBinary 方法), 否则报错`redis: can't marshal []string (implement encoding.BinaryMarshaler)`。建议直接转成 json string 或者 []byte
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
            if err := initGoredisClient().ZAdd(tt.args.key, tt.args.members...); (err != nil) != tt.wantErr {
                t.Errorf("ZAdd() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}

func TestGoredis_ZRangeByScore(t *testing.T) {
    
    type fields struct {
        ctx context.Context
        cli *redis.Client
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
            gotData, err := initGoredisClient().ZRangeByScore(tt.args.key, tt.args.opt)
            if (err != nil) != tt.wantErr {
                t.Errorf("ZRangeByScore() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            fmt.Printf("gotData type:%T data:%v \n", gotData, gotData)
        })
    }
}

func TestGoredis_ZRem(t *testing.T) {
    
    type fields struct {
        ctx context.Context
        cli *redis.Client
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
                key: "k",
                members: []interface{}{
                    "v000",
                    "v001",
                },
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if err := initGoredisClient().ZRem(tt.args.key, tt.args.members...); (err != nil) != tt.wantErr {
                t.Errorf("ZRem() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}
