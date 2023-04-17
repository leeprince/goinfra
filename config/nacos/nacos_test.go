package nacos

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"testing"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/6/4 下午5:14
 * @Desc:
 */

type DynamicConfigTest struct {
	IsMock       bool   `yaml:"is_mock"`
	AppName      string `yaml:"appName"`
	Env          string `yaml:"env"`
	Version      string `yaml:"version"`
	SignType     string `yaml:"signType"`
	RandomNumber int    `yaml:"randomNumber"`
}

var myConfig *DynamicConfigTest

// 测试初始化操作：初始化配置
func TestMain(m *testing.M) {
	myConfig = &DynamicConfigTest{}

	os.Exit(m.Run())
}

func TestNacosClient_ListenConfig(t *testing.T) {
	c, err := NewNacosClient("8371bb89-804e-4549-9855-0b581df2fcf6",
		"dev",
		"config:goinfra",
		WithLogDir("./logs"),
		WithIpAddr("https://apigw-local-test.goldentec.com"),
		WithPort(443),
	)

	if err != nil {
		fmt.Println("NewNacosClient err:", err)
		return
	}
	dynamicConfigHandle := func(conf []byte) {
		err := yaml.Unmarshal(conf, myConfig)
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

// 当远程配置变更时通知接口层。项目中配置变更处于基础服务层，所以采用注册的方式把配置变更时需要通知的函数注册进来
func TestNacosClient_ListenConfigAndNoticeInterfaceEvent(t *testing.T) {
	c, err := NewNacosClient("8371bb89-804e-4549-9855-0b581df2fcf6",
		"dev",
		"config:goinfra",
		WithLogDir("./logs"),
		WithIpAddr("https://apigw-local-test.goldentec.com"),
		WithPort(443),
	)

	if err != nil {
		fmt.Println("NewNacosClient err:", err)
		return
	}

	// 配置变更时，需要通知的函数数组
	var events []RemoteConfigUpdateEvent
	// 注册事件到通知函数数组
	events = append(events, ListentConfigChange)

	dynamicConfigHandle := func(conf []byte) {
		oldConfig := *myConfig
		err := yaml.Unmarshal(conf, myConfig)
		if err != nil {
			fmt.Println("dynamicConfigHandle json.Unmarshal err:", err)
			return
		}
		fmt.Printf("myConfig:%+v \n", myConfig)

		// 配置变更时触发的事件函数。
		// 注意：dynamicConfigHandle该函数`先立即同步获取配置`时会触发一次
		for _, event := range events {
			event(&oldConfig)
		}
	}
	err = c.ListenConfig(dynamicConfigHandle)
	if err != nil {
		fmt.Println("c.ListenConfig err:", err)
		return
	}

	// 调试，定时打印
	go func() {
		for {
			time.Sleep(time.Second * 5)
			fmt.Println(">>>>>>>>>>", myConfig)
		}
	}()

	select {}
}

type RemoteConfigUpdateEvent func(oldConfig *DynamicConfigTest)

func ListentConfigChange(oldConfig *DynamicConfigTest) {
	fmt.Println("ListentConfigChange old------", oldConfig)
	fmt.Println("ListentConfigChange new>>>>>>", myConfig)
	return
}
