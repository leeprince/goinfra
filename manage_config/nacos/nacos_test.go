package nacos

import (
    "fmt"
    "github.com/leeprince/goinfra/config"
    "gopkg.in/yaml.v3"
    "testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/6/4 下午5:14
 * @Desc:
 */

func TestNacosClient_ListenConfig(t *testing.T) {
    c, err := NewNacosClient(
        WithNamespaceId("8371bb89-804e-4549-9855-0b581df2fcf6"),
        WithDataId("config:goinfra"),
        // WithGoup("local"),
        WithGoup("dev"),
    )
    if err != nil {
        fmt.Println("NewNacosClient err:", err)
        return
    }
    myConfig := config.DynamicTest{}
    dynamicConfigHandle := func(conf []byte) {
        err := yaml.Unmarshal(conf, &myConfig)
        if err != nil {
            fmt.Println("dynamicConfigHandle json.Unmarshal err:", err)
            return
        }
        fmt.Printf("myConfig:%+v \n", myConfig)
    }
    err = c.ListenConfig(dynamicConfigHandle)
    if err != nil {
        fmt.Println("c.ListenConfig err:", err)
        return
    }
    
    select {}
}
