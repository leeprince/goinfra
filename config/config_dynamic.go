package config

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/6/4 下午3:49
 * @Desc:   分布式配置中心的动态配置
 */

type DynamicTest struct {
    AppName      string `yaml:"appName"`
    ENV          string `yaml:"env"`
    Version      string `yaml:"version"`
    SignType     string `yaml:"signType"`
    RandomNumber int    `yaml:"randomNumber"`
}
