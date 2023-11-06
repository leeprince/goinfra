package config

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/10/29 14:27
 * @Desc:
 */

type Config struct {
	AI
}

type AI struct {
	OpenAIChatGPT OpenAIChatGPT
}

type OpenAIChatGPT struct {
	Token string
}
