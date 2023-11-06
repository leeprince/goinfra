package main

import (
	"context"
	"fmt"
	"github.com/leeprince/goinfra/config"
	"github.com/sashabaranov/go-openai"
	"net/http"
	"net/url"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/10/29 14:12
 * @Desc:
 */

func init() {
	config.InitConfig("./config/config.yaml")
}

func main() {
	fmt.Println("config.C.AI.OpenAIChatGPT.Token:", config.C.AI.OpenAIChatGPT.Token)
	
	// 无代理
	// client := openai.NewClient(config.C.AI.OpenAIChatGPT.Token)
	
	// 配置代理
	config := openai.DefaultConfig(config.C.AI.OpenAIChatGPT.Token)
	proxyUrl, err := url.Parse("http://localhost:7890")
	if err != nil {
		panic(err)
	}
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
	}
	config.HTTPClient = &http.Client{
		Transport: transport,
	}
	client := openai.NewClientWithConfig(config)
	
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Hello!",
				},
			},
		},
	)
	
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}
	
	fmt.Println(resp.Choices[0].Message.Content)
}
