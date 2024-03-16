package main

import (
	"getwebpage-tomarkdown/internel/Infrastructure/app"
	"getwebpage-tomarkdown/internel/api"
)

func main() {
	// 初始化配置
	app.InitConfig()
	
	// 初始化接口层：暂不外放；后面结合 gin框架去实现路由
	api.NewAPI()
	
}
