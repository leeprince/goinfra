package nacos

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"testing"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/6/4 下午5:14
 * @Desc:
 */

type DynamicTest struct {
	IsMock       bool   `yaml:"is_mock"`
	AppName      string `yaml:"appName"`
	ENV          string `yaml:"env"`
	Version      string `yaml:"version"`
	SignType     string `yaml:"signType"`
	RandomNumber int    `yaml:"randomNumber"`
}

var myConfig DynamicTest

func TestNacosClient_ListenConfig(t *testing.T) {
	c, err := NewNacosClient("8371bb89-804e-4549-9855-0b581df2fcf6",
		"dev",
		"config:goinfra",
		WithLogDir("./logs"),
	)

	if err != nil {
		fmt.Println("NewNacosClient err:", err)
		return
	}
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

	// 调试，定时打印
	go func() {
		for {
			time.Sleep(time.Second * 10)
			fmt.Println(">>>>>>>>>>", myConfig)
		}
	}()

	select {}
}
