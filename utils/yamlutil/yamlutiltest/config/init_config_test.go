package config

import (
	"flag"
	"github.com/leeprince/goinfra/plog"
	"github.com/leeprince/goinfra/utils/yamlutil"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/5/8 09:21
 * @Desc:	初始化配置
 */

var Config *conf

var (
	configPath string
)

func TestInitConf(t *testing.T) {
	InitConfig()
}

// confPath 优先级大于 flag.StringVar 获取的配置文件路径
func InitConfig(confPath ...string) {
	plog.Info("initConfigLocal")

	if len(confPath) > 0 {
		configPath = confPath[0]
	} else {
		flag.StringVar(&configPath, "conf", "./conf.yaml", "config name")
		flag.Parse()
	}
	plog.WithField("configPath", configPath).Info("initConfigLocal")

	// 解析配置文件
	// 不管是定义：var Config *conf，还是定义 var Config conf。解析时都是传递&Config =>yamlutil.ParseYaml(configPath, &Config)
	yamlutil.ParseYaml(configPath, &Config)

	plog.WithField("configPath", configPath).WithField("Config", Config).Info("initConfigLocal:")
}

type conf struct {
	Env          string `yaml:"Env"`
	AppName      string `yaml:"AppName"`
	Port         int64  `yaml:"Port"`
	LogPath      string `yaml:"LogPath"`
	LogLevel     int32  `yaml:"LogLevel"`
	IsBothStdout bool   `yaml:"IsBothStdout"`

	// 网站信息
	Website website `yaml:"Website"`

	// 监听网络信息
	ListenNetwork listenNetwork `yaml:"ListenNetwork"`
}

type website struct {
	LoginUrl      string `yaml:"LoginUrl"`
	LoginUserName string `yaml:"LoginUserName"`
	LoginPassword string `yaml:"LoginPassword"`
}

type listenNetwork struct {
	Order order `yaml:"Order"`
}
type order struct {
	Url                  string   `yaml:"Url"`
	NewOrderBodyContains string   `yaml:"NewOrderBodyContains"` // 新订单时的响应包含内容
	SaveNewOrderInfo     fileInfo `yaml:"SaveNewOrderInfo"`     // 新订单的保存信息
}
type fileInfo struct {
	FilePath string `yaml:"FilePath"` // 文件路径
	FileName string `yaml:"FileName"` // 文件名
	IsAppend bool   `yaml:"IsAppend"` // 是否在该文件上追加内容
}
