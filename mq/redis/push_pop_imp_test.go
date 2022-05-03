package redis

import (
    "encoding/json"
    "fmt"
    "github.com/leeprince/goinfra/config"
    goinfraRedis "github.com/leeprince/goinfra/storage/redis"
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
    key = "princeKey01"
    // value  = "princeValue02"
    // value  = []string{"v001", "v002"}
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

func TestPushPop_Publish(t *testing.T) {
    // --- redis 客户端
    
    // Goredis 客户端
    err := goinfraRedis.InitGoredis(RedisConfs)
    if err != nil {
        fmt.Printf("[goinfraRedis.InitGoredis] err:%v \n", err)
        return
    }
    redisClient := goinfraRedis.GetGoredis(RedisName)
    
    // Redigo 客户端
    /*err := goinfraRedis.InitRedigo(RedisConfs)
      if err != nil {
          fmt.Printf("[goinfraRedis.InitGoredis] err:%v \n", err)
          return
      }
      redisClient := goinfraRedis.GetRedigo(RedisName)*/
    
    // --- redis 客户端-end
    
    type fields struct {
        cli goinfraRedis.RedisClient
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
            fields: fields{cli: redisClient},
            args: args{
                key:   key,
                value: value,
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            mq := &PushPop{
                cli: tt.fields.cli,
            }
            if err := mq.Publish(tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
                t.Errorf("Publish() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}

func TestPushPop_Subscribe(t *testing.T) {
    // --- redis 客户端
    
    // Goredis 客户端
    err := goinfraRedis.InitGoredis(RedisConfs)
    if err != nil {
        fmt.Printf("[goinfraRedis.InitGoredis] err:%v \n", err)
        return
    }
    redisClient := goinfraRedis.GetGoredis(RedisName)
    
    // Redigo 客户端
    /*err := goinfraRedis.InitRedigo(RedisConfs)
      if err != nil {
          fmt.Printf("[goinfraRedis.InitGoredis] err:%v \n", err)
          return
      }
      redisClient := goinfraRedis.GetRedigo(RedisName)*/
    
    // --- redis 客户端-end
    
    type fields struct {
        cli goinfraRedis.RedisClient
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
            fields: fields{cli: redisClient},
            args: args{
                key:     key,
                timeout: time.Second * 2,
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            mq := &PushPop{
                cli: tt.fields.cli,
            }
            gotData, err := mq.Subscribe(tt.args.key, tt.args.timeout)
            if (err != nil) != tt.wantErr {
                t.Errorf("Subscribe() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            // data := string(gotData.([]byte))
            data := cast.ToString(gotData)
            fmt.Println("data string:", data)
        })
    }
}
