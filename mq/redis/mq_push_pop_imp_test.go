package redis

import (
    "encoding/json"
    "fmt"
    "github.com/leeprince/goinfra/config"
    "github.com/leeprince/goinfra/storage/redis"
    "github.com/spf13/cast"
    "testing"
    "time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/5/3 上午12:32
 * @Desc:   测试用例
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
    key = "k"
    // value  = "Value01"
    // value = []string{"v001", "v002"}
    value = ValueStruct{
        Name: "prince",
        Age:  18,
    }
    expire = time.Second * 3
)

type ValueStruct struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func (s ValueStruct) MarshalBinary() ([]byte, error) {
    return json.Marshal(s)
}

func initRedisClient() (redisClient redis.RedisClient) {
    // --- redis 客户端
    var err error
    
    // Goredis 客户端
    err = redis.InitGoredis(RedisConfs)
    if err != nil {
        panic(fmt.Sprintf("[redis.InitGoredis] err:%v \n", err))
        return
    }
    redisClient = redis.GetGoredis(RedisName)
    
    // Redigo 客户端
    /*err = redis.InitRedigo(RedisConfs)
      if err != nil {
          panic(fmt.Sprintf("[redis.InitRedigo] err:%v \n", err))
          return
      }
      redisClient = redis.GetRedigo(RedisName)*/
    
    // --- redis 客户端-end
    
    return
}

func TestListMQ_Push(t *testing.T) {
    
    type fields struct {
        cli redis.RedisClient
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
        {
            args: args{
                key:   key,
                value: value,
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            mq := NewListMQ(initRedisClient())
            if err := mq.Push(tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
                t.Errorf("Push() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}

func TestListMQ_Subscribe(t *testing.T) {
    type fields struct {
        cli redis.RedisClient
    }
    type args struct {
        key     string
        timeout time.Duration
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
                key:     key,
                timeout: time.Second * 2,
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            mq := NewListMQ(initRedisClient())
            
            f := func(data interface{}) {
                fmt.Println("(mq *ListMQ) Subscribe data:", cast.ToString(data))
            }
            
            mq.Subscribe(f, tt.args.key, tt.args.timeout)
        })
    }
}
