package config

import (
	"flag"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var Config Conf

type Conf struct {
	Env      string `yaml:"env"`
	LogLevel uint32 `yaml:"log_level"`
	Mode     string `yaml:"mode"`
	
	GinHost struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"gin_host"`
	
	Heartbeat      int    `yaml:"heartbeat"`
	JaegerAgentURL string `yaml:"jaegeragenturl"`
}

func init() {
	config := flag.String("config", "./configs/conf.yaml", "Set the config (default conf.yml)")
	ParseYaml(*config, &Config)
	flag.Parse()
}

func ParseYaml(file string, configRaw interface{}) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		panic("加载配置文件错误" + file + "错误原因" + err.Error())
	}
	
	err = yaml.Unmarshal(content, configRaw)
	if err != nil {
		panic("解析配置文件错误" + file + "错误原因" + err.Error())
	}
}
