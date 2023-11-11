package config

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/10/29 14:27
 * @Desc:
 */

type Config struct {
	AI `yaml:"AI"`
}

type AI struct {
	OpenAIChatGPT OpenAIChatGPT `yaml:"OpenAIChatGPT"`
}

type OpenAIChatGPT struct {
	SecretKey string `yaml:"SecretKey"`
}
