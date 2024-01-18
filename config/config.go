package config

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/10/29 14:27
 * @Desc:
 */

type Config struct {
	AI   `yaml:"AI"`
	COS  `yaml:"COS"`
	Host `yaml:"FileAccessHost"`
}

type AI struct {
	OpenAIChatGPT OpenAIChatGPT `yaml:"OpenAIChatGPT"`
}

type OpenAIChatGPT struct {
	SecretKey string `yaml:"SecretKey"`
}

type COS struct {
	SecretID  string `yaml:"SecretID"`
	SecretKey string `yaml:"SecretKey"`
	Bucket    string `yaml:"Bucket"`
	AppID     string `yaml:"AppID"`
	Region    string `yaml:"Region"`
}

type Host struct {
	CosFileAccessHost string `yaml:"CosFileAccessHost"`
}
